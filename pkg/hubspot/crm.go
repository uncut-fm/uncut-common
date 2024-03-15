package hubspot

import (
	"context"
	"fmt"
)

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

	resp := &hubspotBatchObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

type batchUpdateInput struct {
	Properties map[string]string `json:"properties"`
	IdProperty *string           `json:"idProperty,omitempty"`
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

	resp := &hubspotBatchObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)

	return err
}

func (c *Client) CreateProperties(ctx context.Context, newProperties []NewProperty) error {
	endpoint := fmt.Sprintf("%s/contacts/batch/create", propertiesBaseAPIUrl)

	payload := map[string][]NewProperty{
		"inputs": newProperties,
	}

	resp := &hubspotBatchObjectsResponseV3{}

	err := c.sendPostRequest(endpoint, payload, resp)

	return err
}

func (c *Client) CreateObjectsBatch(ctx context.Context, objectType HubspotObjectType, objects []HubspotSimplePublicObjectInput) ([]HubspotSimplePublicObjectInput, error) {
	endpoint := fmt.Sprintf(batchObjectsCreateBaseAPIUrl, string(objectType))
	payload := map[string][]HubspotSimplePublicObjectInput{
		"inputs": objects,
	}

	resp := &hubspotBatchObjectsResponseV3{}

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

func (c *Client) CreateObject(ctx context.Context, objectType HubspotObjectType, object CreateObjectInput) (HubspotSimplePublicObjectInput, error) {
	endpoint := fmt.Sprintf(objectCreateBaseAPIUrl, string(objectType))

	resp := &hubspotObjectResponseV3{}

	err := c.sendPostRequest(endpoint, object, resp)
	if err != nil {
		return HubspotSimplePublicObjectInput{}, err
	}

	return resp.HubspotSimplePublicObjectInput, nil
}

func (c *Client) ListAssociationsByContactIDs(ctx context.Context, toObjectType HubspotObjectType, contactIDs []string) ([]AssociationPairToArrayObject, error) {
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
