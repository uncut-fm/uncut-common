package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
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

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(RequestTimeout)
	client.SetRetryCount(1)

	return client
}

func NewWebappClient(log logger.Logger, webappURL string) *WebappClient {
	return &WebappClient{
		log:                    log,
		restyClient:            createRestyClient(),
		markdownToHtmlEndpoint: fmt.Sprintf("%s/api/markdown-to-html", webappURL),
	}
}

func (w WebappClient) GetHTMLFromMarkdown(markdown string) (string, error) {
	html, err := w.makeMarkdownToHtmlRequest(markdown)

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