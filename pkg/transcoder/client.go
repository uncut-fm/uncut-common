package transcoder

import (
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/tracing"
	"go.opentelemetry.io/otel/trace"
	"time"
)

const (
	requestTimeout        = 1 * time.Minute
	gifFirstFrame         = "%s/gif/first-frame"
	videoFirstFrame       = "%s/video/first-frame"
	videoMetadata         = "%s/video/metadata"
	imageMetadata         = "%s/image/metadata"
	audioMetadataEndpoint = "%s/audio/metadata"
)

type API struct {
	log         logger.Logger
	url         string
	adminToken  string
	restyClient *resty.Client
}

func New(log logger.Logger, tp trace.TracerProvider, url, token string) *API {
	return &API{
		log:         log,
		url:         url,
		adminToken:  token,
		restyClient: createRestyClient(tp),
	}
}

func createRestyClient(tp trace.TracerProvider) *resty.Client {
	client := resty.New().
		SetTransport(tracing.NewTransport(tp)).
		SetTimeout(requestTimeout).
		SetRetryCount(2)

	return client
}
