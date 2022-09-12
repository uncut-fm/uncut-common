package rss

import (
	"context"
	"errors"
	"github.com/mmcdole/gofeed"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	audioURLErr      = errors.New("no audioURL detected")
	itunesExtErr     = errors.New("itunesExt is missing")
	enclosuresExtErr = errors.New("enclosures are missing")
)

type Client struct {
	log logger.Logger
}

func NewClient(log logger.Logger) *Client {
	return &Client{log: log}
}

func (c Client) GetFeedByFeedURL(feedURL string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := fp.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		return nil, err
	}

	c.sortFeedItemsByPublishedField(feed.Items)

	return feed, nil
}

func (c Client) getHost(urlString string) (*string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	hostname := u.Hostname()

	return &hostname, nil
}

func (c Client) sortFeedItemsByPublishedField(items []*gofeed.Item) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].PublishedParsed == nil {
			log.Println("PUBLISHED_PARSED IS MISSING")
			return true
		}
		return items[i].PublishedParsed.Before(*items[j].PublishedParsed)
	})
}

type NewEpisode struct {
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	GUID          *string    `json:"guid"`
	CoverImageURL *string    `json:"coverImageUrl"`
	AudioURL      *string    `json:"audioUrl"`
	Duration      *int       `json:"duration"`
	ReleaseDate   *time.Time `json:"releaseDate"`
	EpisodeNumber *int       `json:"episodeNumber"`
}

func (c Client) ParseEpisode(item *gofeed.Item) (*NewEpisode, error) {
	return c.getNewEpisodeFromFeedItem(item)
}

func (c Client) getNewEpisodeFromFeedItem(item *gofeed.Item) (*NewEpisode, error) {
	if item.ITunesExt == nil {
		return nil, itunesExtErr
	}

	if item.Enclosures == nil {
		return nil, enclosuresExtErr
	}

	ep := &NewEpisode{
		Title:       item.Title,
		Description: item.ITunesExt.Subtitle,
	}

	if ep.Description == "" {
		ep.Description = item.ITunesExt.Summary
		if ep.Description == "" {
			ep.Description = item.Description
		}
	}

	if !strings.Contains(ep.Description, "<p>") {
		ep.Description = parseStringDescription(ep.Description)
	}

	duration, err := parseDuration(item.ITunesExt.Duration)
	if err != nil {
		return nil, err
	}

	image := item.ITunesExt.Image

	ep.GUID = &item.GUID
	ep.Duration = &duration
	ep.ReleaseDate = item.PublishedParsed
	ep.CoverImageURL = &image

	var audioURL string

	for _, vv := range item.Enclosures {
		if strings.Contains(vv.Type, "video") {
			continue
		}
		audioURL = vv.URL
	}

	if audioURL == "" {
		return ep, audioURLErr
	}

	ep.AudioURL = &audioURL

	return ep, nil
}
