package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/tracing"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

const (
	batchObjectsReadBaseAPIUrl   = "https://api.hubapi.com/crm/v3/objects/%s/batch/read"
	batchObjectsCreateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/batch/create"
	batchObjectsUpdateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/batch/update"

	objectsSearchBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/search" // /crm/v3/objects/:object_type/search

	objectCreateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s"    // /crm/v3/objects/:object_type
	objectDeleteBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/%s" // /crm/v3/objects/:object_type/:id
	objectUpdateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/%s" // /crm/v3/objects/:object_type/:id

	propertiesBaseAPIUrl = "https://api.hubapi.com/crm/v3/properties"

	batchAssociationsReadBaseAPIUrl    = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/read"    // /crm/v3/associations/:from-object-type/:to-object-type/batch/read
	batchAssociationsCreateBaseAPIUrl  = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/create"  // /crm/v3/associations/:from-object-type/:to-object-type/batch/create
	batchAssociationsArchiveBaseAPIUrl = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/archive" // /crm/v3/associations/:from-object-type/:to-object-type/batch/archive
)

type Client struct {
	log    logger.Logger
	apiKey string
	client *http.Client
}

func NewClient(log logger.Logger, tp trace.TracerProvider, apiKey string) *Client {
	return &Client{
		log:    log,
		apiKey: apiKey,
		client: tracing.NewHTTPClient(tp),
	}
}

func (c *Client) sendPostRequest(ctx context.Context, url string, data interface{}, hubspotResp hubspotResponseInterface) error {
	return c.sendRequest(ctx, url, data, hubspotResp, http.MethodPost)
}

func (c *Client) sendPatchRequest(ctx context.Context, url string, data interface{}, hubspotResp hubspotResponseInterface) error {
	return c.sendRequest(ctx, url, data, hubspotResp, http.MethodPatch)
}

func (c *Client) sendRequest(ctx context.Context, url string, data interface{}, hubspotResp hubspotResponseInterface, method string) error {
	payload := new(bytes.Buffer)
	enc := json.NewEncoder(payload)
	enc.SetEscapeHTML(false)

	err := enc.Encode(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	// Add Bearer token to the header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if hubspotResp == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, hubspotResp)
	if err != nil {
		return err
	}

	if resp.StatusCode > 300 && hubspotResp.GetStatus() != hubspotStatusComplete {
		// print the response

		c.log.Info(string(body))

		return errors.New(fmt.Sprintf("HubSpot API returned non-200 status code: %d, message: %s", resp.StatusCode, hubspotResp.GetMessage()))
	}

	return nil
}

func (c *Client) sendDeleteRequest(url string) error {
	req, err := http.NewRequest("DELETE", url, io.Reader(nil))
	if err != nil {
		return err
	}

	// Add Bearer token to the header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 300 {
		return errors.New(fmt.Sprintf("HubSpot API returned non-200 status code: %d", resp.StatusCode))
	}

	return nil
}
