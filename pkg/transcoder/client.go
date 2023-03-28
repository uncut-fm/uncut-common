package transcoder

import (
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"time"
)

const (
	requestTimeout = 1 * time.Minute
	gifFirstFrame  = "%s/gif/first-frame"
)

type API struct {
	log         logger.Logger
	url         string
	adminToken  string
	restyClient *resty.Client
}

func New(log logger.Logger, url, token string) *API {
	return &API{
		log:         log,
		url:         url,
		adminToken:  token,
		restyClient: createRestyClient(),
	}
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout).
		SetRetryCount(2)

	return client
}
