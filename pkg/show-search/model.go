package show_search

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

type SearchShow struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	FeedURL       string `json:"feed_url"`
	ArtworkURL600 string `json:"artwork_url_600"`
	AppStoreURL   string `json:"app_store_url"`
}
