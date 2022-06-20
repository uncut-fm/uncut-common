package show_search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/errors"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"strconv"
	"time"
)

const (
	appleSearchURL = "https://itunes.apple.com/search"
	requestTimeout = 5 * time.Second
)

type ItunesClient struct {
	log         logger.Logger
	restyClient *resty.Client
}

func NewItunesClient(log logger.Logger) *ItunesClient {
	return &ItunesClient{
		log:         log,
		restyClient: createRestyClient(),
	}
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout)

	return client
}

func (c ItunesClient) GetShowByName(showName string) (*SearchShow, error) {
	searchResult, err := c.searchShowByName(showName)
	if c.log.CheckError(err, c.GetShowByName) != nil {
		return nil, err
	}

	if len(searchResult) == 0 {
		return nil, errors.NoSearchShowsFoundErr
	}

	show, err := c.getShowFromSearchShowsByName(searchResult, showName), nil
	if c.log.CheckError(err, c.GetShowByName) != nil {
		return nil, err
	}

	if show == nil {
		showNames := c.getShowNamesFromSearchShows(searchResult)
		return getNewSearchShowFromAppleShow(searchResult[0]), c.log.CheckError(errors.NoSearchShowsExactFoundErr(showNames), c.GetShowByName)
	}

	return show, nil
}

func (c ItunesClient) searchShowByName(showName string) ([]AppleStoreShow, error) {
	resp, err := c.restyClient.R().EnableTrace().
		SetQueryParams(map[string]string{
			"term":    showName,
			"media":   "podcast",
			"country": "US",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		Get(appleSearchURL)
	if err != nil {
		return []AppleStoreShow{}, nil
	}

	return c.parseAppleStoreResult(resp)
}

func (c ItunesClient) getShowFromSearchShowsByName(searchShows []AppleStoreShow, showName string) *SearchShow {
	for _, show := range searchShows {
		if show.CollectionName == showName || len(searchShows) == 1 {
			return getNewSearchShowFromAppleShow(show)
		}
	}

	return nil
}

func getNewSearchShowFromAppleShow(appleShow AppleStoreShow) *SearchShow {
	return &SearchShow{
		ID:            strconv.FormatInt(appleShow.CollectionId, 10),
		Name:          appleShow.CollectionName,
		FeedURL:       appleShow.FeedUrl,
		ArtworkURL600: appleShow.ArtworkUrl600,
		AppStoreURL:   appleShow.CollectionViewUrl,
	}
}

func (c ItunesClient) parseAppleStoreResult(resp *resty.Response) ([]AppleStoreShow, error) {
	result := AppleStoreResult{}
	err := json.Unmarshal(resp.Body(), &result)
	if c.log.CheckError(err, c.parseAppleStoreResult) != nil {
		return []AppleStoreShow{}, nil
	}

	return result.Results, nil
}

func (c ItunesClient) getShowNamesFromSearchShows(searchShows []AppleStoreShow) (names []string) {
	for _, show := range searchShows {
		names = append(names, show.CollectionName)
	}

	return names
}
