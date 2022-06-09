package show_search

import (
	"encoding/json"
	"github.com/go-resty/resty"
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

func (c ItunesClient) SearchShowByName(showName string) ([]AppleStoreShow, error) {
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

type AppleStoreResult struct {
	ResultCount int              `json:"resultCount"`
	Results     []AppleStoreShow `json:"results"`
}

type AppleStoreShow struct {
	WrapperType            string `json:"wrapperType"`
	CollectionId           int64  `json:"collectionId"`
	TrackId                int64  `json:"trackId"`
	ArtistId               int64  `json:"artistId"`
	ArtistName             string `json:"artistName"`
	CollectionName         string `json:"collectionName"`
	TrackName              string `json:"trackName"`
	CollectionCensoredName string `json:"collectionCensoredName"`
	TrackCensoredName      string `json:"trackCensoredName"`
	CollectionViewUrl      string `json:"collectionViewUrl"`
	FeedUrl                string `json:"feedUrl"`
	TrackViewUrl           string `json:"trackViewUrl"`
	ArtworkUrl30           string `json:"artworkUrl30"`
	ArtworkUrl60           string `json:"artworkUrl60"`
	ArtworkUrl100          string `json:"artworkUrl100"`
	ArtworkUrl600          string `json:"artworkUrl600"`
	ReleaseDate            string `json:"releaseDate"`
	CollectionExplicitness string `json:"collectionExplicitness"`
	ContentAdvisoryRating  string `json:"contentAdvisoryRating"`
	PrimaryGenreName       string `json:"primaryGenreName"`
}

func (c ItunesClient) parseAppleStoreResult(resp *resty.Response) ([]AppleStoreShow, error) {
	result := AppleStoreResult{}
	err := json.Unmarshal(resp.Body(), &result)
	if c.log.CheckError(err, c.parseAppleStoreResult) != nil {
		return []AppleStoreShow{}, nil
	}

	return result.Results, nil
}
