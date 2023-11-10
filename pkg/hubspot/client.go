package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"io"
	"net/http"
)

const (
	batchObjectsReadBaseAPIUrl   = "https://api.hubapi.com/crm/v3/objects/%s/batch/read"
	batchObjectsCreateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/batch/create"
	batchObjectsUpdateBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/batch/update"

	objectDeleteBaseAPIUrl = "https://api.hubapi.com/crm/v3/objects/%s/%s" // /crm/v3/objects/:object_type/:id

	propertiesBaseAPIUrl = "https://api.hubapi.com/crm/v3/properties"

	batchAssociationsReadBaseAPIUrl    = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/read"    // /crm/v3/associations/:from-object-type/:to-object-type/batch/read
	batchAssociationsCreateBaseAPIUrl  = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/create"  // /crm/v3/associations/:from-object-type/:to-object-type/batch/create
	batchAssociationsArchiveBaseAPIUrl = "https://api.hubapi.com/crm/v3/associations/%s/%s/batch/archive" // /crm/v3/associations/:from-object-type/:to-object-type/batch/archive
)

type Client struct {
	log    logger.Logger
	apiKey string
}

func NewClient(log logger.Logger, apiKey string) *Client {
	return &Client{
		log:    log,
		apiKey: apiKey,
	}
}

type batchReadRequest struct {
	Properties []string            `json:"properties"`
	IdProperty *string             `json:"idProperty,omitempty"`
	Inputs     []map[string]string `json:"inputs"`
}

func (c *Client) ListObjects(ctx context.Context, objectType HubspotObjectType, idPropertyField *string, ids, properties []string) ([]HubspotSimplePublicObjectInput, error) {
	endpoint := fmt.Sprintf(batchObjectsReadBaseAPIUrl, objectType)
	payload := batchReadRequest{
		Properties: properties,
	}

	if idPropertyField != nil {
		payload.IdProperty = idPropertyField
	}

	for _, id := range ids {
		payload.Inputs = append(payload.Inputs, map[string]string{
			"id": id,
		})
	}

	resp := &hubspotObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

type batchUpdateInput struct {
	Properties map[string]string `json:"properties"`
	IdProperty string            `json:"idProperty"`
	ID         string            `json:"id"`
}

// UpdateByEmail updates a contact in HubSpot by email
func (c *Client) UpdateContactsByEmailsBatch(ctx context.Context, properties []HubspotSimplePublicObjectInput) error {
	endpoint := fmt.Sprintf(batchObjectsUpdateBaseAPIUrl, ContactsObjectType)

	var inputs []batchUpdateInput

	for _, property := range properties {
		inputs = append(inputs, batchUpdateInput{
			Properties: property.Properties,
			ID:         property.ID,
		})
	}

	payload := map[string][]batchUpdateInput{
		"inputs": inputs,
	}

	resp := &hubspotObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)

	return err
}

func (c *Client) CreateProperties(ctx context.Context, newProperties []NewProperty) error {
	endpoint := fmt.Sprintf("%s/contacts/batch/create", propertiesBaseAPIUrl)

	payload := map[string][]NewProperty{
		"inputs": newProperties,
	}

	resp := &hubspotObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)

	return err
}

func (c *Client) CreateObjectsBatch(ctx context.Context, objectType HubspotObjectType, objects []HubspotSimplePublicObjectInput) ([]HubspotSimplePublicObjectInput, error) {
	endpoint := fmt.Sprintf(batchObjectsCreateBaseAPIUrl, string(objectType))
	payload := map[string][]HubspotSimplePublicObjectInput{
		"inputs": objects,
	}

	resp := &hubspotObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, err
}

func (c *Client) DeleteObjectByID(ctx context.Context, objectType HubspotObjectType, id string) error {
	endpoint := fmt.Sprintf(objectDeleteBaseAPIUrl, string(objectType), id)

	err := c.sendDeleteRequest(endpoint)

	return err
}

func (c *Client) ListAssociationsByContactIDs(ctx context.Context, toObjectType HubspotObjectType, contactIDs []string) ([]AssociationObject, error) {
	endpoint := fmt.Sprintf(batchAssociationsReadBaseAPIUrl, ContactsObjectType, toObjectType)

	type input struct {
		ID string `json:"id"`
	}

	inputs := make([]input, len(contactIDs))

	for i, id := range contactIDs {
		inputs[i] = input{
			ID: id,
		}
	}

	payload := map[string][]input{
		"inputs": inputs,
	}

	resp := &hubspotReadAssociationsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Errors) > 0 {
		resp.setResultsForAssociationsWithErrors()
	}

	return resp.Results, nil
}

func (c *Client) CreateAssociationsBatch(ctx context.Context, toObjectType HubspotObjectType, associations []AssociationPair) error {
	endpoint := fmt.Sprintf(batchAssociationsCreateBaseAPIUrl, ContactsObjectType, toObjectType)

	payload := map[string][]AssociationPair{
		"inputs": associations,
	}

	resp := &hubspotCreateAssociationsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)

	return err
}

func (c *Client) ArchiveAssociationsBatch(ctx context.Context, toObjectType HubspotObjectType, associations []AssociationPair) error {
	endpoint := fmt.Sprintf(batchAssociationsArchiveBaseAPIUrl, ContactsObjectType, toObjectType)

	payload := map[string][]AssociationPair{
		"inputs": associations,
	}

	err := c.sendPostRequest(endpoint, payload, nil)

	return err
}

func (c *Client) sendPostRequest(url string, data interface{}, hubspotResp hubspotResponseInterface) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	// Add Bearer token to the header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 300 {
		return errors.New(fmt.Sprintf("HubSpot API returned non-200 status code: %d", resp.StatusCode))
	}

	return nil
}

type hubspotObjectsResponseV3 struct {
	Status        hubspotStatus                    `json:"status"`
	Message       string                           `json:"message"`
	CorrelationID string                           `json:"correlationId"`
	Category      string                           `json:"category"`
	Results       []HubspotSimplePublicObjectInput `json:"results"`
}

func (h hubspotObjectsResponseV3) GetStatus() hubspotStatus {
	return h.Status
}

func (h hubspotObjectsResponseV3) GetMessage() string {
	return h.Message
}

type associationResponseErr struct {
	Message string `json:"message"`
	Context struct {
		FromObjectID []string `json:"fromObjectId"`
	}
}

type hubspotReadAssociationsResponseV3 struct {
	Status  hubspotStatus            `json:"status"`
	Errors  []associationResponseErr `json:"errors"`
	Results []AssociationObject      `json:"results"`
}

func (h hubspotReadAssociationsResponseV3) GetStatus() hubspotStatus {
	return h.Status
}

func (h hubspotReadAssociationsResponseV3) GetMessage() string {
	if len(h.Errors) > 0 {
		return h.Errors[0].Message
	}

	return ""
}

func (h *hubspotReadAssociationsResponseV3) setResultsForAssociationsWithErrors() {
	for _, err := range h.Errors {
		if len(err.Context.FromObjectID) > 0 {
			h.Results = append(h.Results, AssociationObject{
				From: struct {
					ID string `json:"id"`
				}{
					ID: err.Context.FromObjectID[0],
				},
			})
		}
	}
}

type hubspotCreateAssociationsResponseV3 struct {
	Status  hubspotStatus            `json:"status"`
	Errors  []associationResponseErr `json:"errors"`
	Results []AssociationPair        `json:"results"`
}

func (h hubspotCreateAssociationsResponseV3) GetStatus() hubspotStatus {
	return h.Status
}

func (h hubspotCreateAssociationsResponseV3) GetMessage() string {
	if len(h.Errors) > 0 {
		return h.Errors[0].Message
	}

	return ""
}

type hubspotResponseInterface interface {
	GetStatus() hubspotStatus
	GetMessage() string
}

type hubspotStatus string

const (
	hubspotStatusError    hubspotStatus = "ERROR"
	hubspotStatusComplete hubspotStatus = "COMPLETE"
)
