package twitter

import (
	"context"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClient(ctx context.Context, twitterConfigs config.TwitterConfigs) *twitter.Client {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     twitterConfigs.ConsumerKey,
		ClientSecret: twitterConfigs.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Users will automatically authorize Requests
	httpClient := config.Client(ctx)

	// Twitter client
	return twitter.NewClient(httpClient)
}
