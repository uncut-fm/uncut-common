package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/tracing"
	"go.opentelemetry.io/otel/trace"
	"time"
)

const (
	RequestTimeout = 5 * time.Second
)

type WebappClient struct {
	log                    logger.Logger
	restyClient            *resty.Client
	markdownToHtmlEndpoint string
}

func createRestyClient(tp trace.TracerProvider) *resty.Client {
	client := resty.New().
		SetTransport(tracing.NewTransport(tp)).
		SetTimeout(RequestTimeout).
		SetRetryCount(1)

	return client
}

func NewWebappClient(log logger.Logger, tp trace.TracerProvider, webappURL string) *WebappClient {
	return &WebappClient{
		log:                    log,
		restyClient:            createRestyClient(tp),
		markdownToHtmlEndpoint: fmt.Sprintf("%s/api/markdown-to-html", webappURL),
	}
}

func (w WebappClient) GetHTMLFromMarkdown(markdown string) (string, error) {
	var (
		html string
		err  error
	)

	operation := func() error {
		html, err = w.makeMarkdownToHtmlRequest(markdown)
		if err != nil {
			w.log.Warn(err)
			return err
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 30 * time.Second

	err = backoff.Retry(operation, b)

	return html, w.log.CheckError(err, w.GetHTMLFromMarkdown)
}

func (w WebappClient) makeMarkdownToHtmlRequest(markdown string) (string, error) {
	resp, err := w.restyClient.R().EnableTrace().
		SetBody(map[string]string{"markdown": markdown}).
		Post(w.markdownToHtmlEndpoint)

	if w.log.CheckError(err, w.makeMarkdownToHtmlRequest) != nil {
		return "", err
	}

	responseMap := make(map[string]string)
	err = json.Unmarshal(resp.Body(), &responseMap)

	return responseMap["html"], w.log.CheckError(err, w.makeMarkdownToHtmlRequest)
}
