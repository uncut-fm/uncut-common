package hubspot

import (
	"context"
	"fmt"
)

func (c *Client) ListObjects(ctx context.Context, objectType HubspotObjectType, idPropertyField *string, ids, properties []string) ([]HubspotSimplePublicObjectInput, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.ListObjects))
	defer span.End()

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

	err := c.sendPostRequest(ctx, endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func (c *Client) SearchObjectsByProperty(ctx context.Context, objectType HubspotObjectType, propertyName string, propertyValue string, properties []string) ([]HubspotSimplePublicObjectInput, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.SearchObjectsByProperty))
	defer span.End()

	endpoint := fmt.Sprintf(objectsSearchBaseAPIUrl, objectType)
	payload := searchRequest{
		FilterGroups: []searchFilterGroup{
			{
				Filters: []searchFilter{
					{
						PropertyName: propertyName,
						Operator:     "EQ",
						Value:        propertyValue,
					},
				},
			},
		},
	}

	if len(properties) > 0 {
		payload.Properties = properties
	}

	resp := &hubspotSearchObjectsResponseV3{}

	err := c.sendPostRequest(ctx, endpoint, payload, resp)
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
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.UpdateContactsByEmailsBatch))
	defer span.End()

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

	err := c.sendPostRequest(ctx, endpoint, payload, resp)

	return err
}

func (c *Client) CreateProperties(ctx context.Context, newProperties []NewProperty) error {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.CreateProperties))
	defer span.End()

	endpoint := fmt.Sprintf("%s/contacts/batch/create", propertiesBaseAPIUrl)

	payload := map[string][]NewProperty{
		"inputs": newProperties,
	}

	resp := &hubspotBatchObjectsResponseV3{}

	err := c.sendPostRequest(ctx, endpoint, payload, resp)

	return err
}

func (c *Client) CreateObjectsBatch(ctx context.Context, objectType HubspotObjectType, objects []HubspotSimplePublicObjectInput) ([]HubspotSimplePublicObjectInput, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.CreateObjectsBatch))
	defer span.End()

	endpoint := fmt.Sprintf(batchObjectsCreateBaseAPIUrl, string(objectType))
	payload := map[string][]HubspotSimplePublicObjectInput{
		"inputs": objects,
	}

	resp := &hubspotBatchObjectsResponseV3{}

	err := c.sendPostRequest(ctx, endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, err
}

func (c *Client) DeleteObjectByID(ctx context.Context, objectType HubspotObjectType, id string) error {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.DeleteObjectByID))
	defer span.End()

	endpoint := fmt.Sprintf(objectDeleteBaseAPIUrl, string(objectType), id)

	err := c.sendDeleteRequest(endpoint)

	return err
}

func (c *Client) CreateObject(ctx context.Context, objectType HubspotObjectType, object CreateObjectInput) (HubspotSimplePublicObjectInput, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.CreateObject))
	defer span.End()

	endpoint := fmt.Sprintf(objectCreateBaseAPIUrl, string(objectType))

	resp := &hubspotObjectResponseV3{}

	err := c.sendPostRequest(ctx, endpoint, object, resp)
	if err != nil {
		return HubspotSimplePublicObjectInput{}, err
	}

	return resp.HubspotSimplePublicObjectInput, nil
}

func (c *Client) UpdateObject(ctx context.Context, objectType HubspotObjectType, object UpdateObjectInput) (HubspotSimplePublicObjectInput, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.UpdateObject))
	defer span.End()

	endpoint := fmt.Sprintf(objectUpdateBaseAPIUrl, string(objectType), object.ObjectID)

	resp := &hubspotObjectResponseV3{}

	err := c.sendPatchRequest(ctx, endpoint, object, resp)
	if err != nil {
		return HubspotSimplePublicObjectInput{}, err
	}

	return resp.HubspotSimplePublicObjectInput, nil
}

func (c *Client) ListAssociationsByContactIDs(ctx context.Context, toObjectType HubspotObjectType, contactIDs []string) ([]AssociationPairToArrayObject, error) {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.ListAssociationsByContactIDs))
	defer span.End()

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

	err := c.sendPostRequest(ctx, endpoint, payload, resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Errors) > 0 {
		resp.setResultsForAssociationsWithErrors()
	}

	return resp.Results, nil
}

func (c *Client) CreateAssociationsBatch(ctx context.Context, toObjectType HubspotObjectType, associations []AssociationPair) error {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.CreateAssociationsBatch))
	defer span.End()

	endpoint := fmt.Sprintf(batchAssociationsCreateBaseAPIUrl, ContactsObjectType, toObjectType)

	payload := map[string][]AssociationPair{
		"inputs": associations,
	}

	resp := &hubspotCreateAssociationsResponseV3{}

	err := c.sendPostRequest(ctx, endpoint, payload, resp)

	return err
}

func (c *Client) ArchiveAssociationsBatch(ctx context.Context, toObjectType HubspotObjectType, associations []AssociationPair) error {
	ctx, span := c.tracer.Start(ctx, c.log.GetFunctionName(c.ArchiveAssociationsBatch))
	defer span.End()

	endpoint := fmt.Sprintf(batchAssociationsArchiveBaseAPIUrl, ContactsObjectType, toObjectType)

	payload := map[string][]AssociationPair{
		"inputs": associations,
	}

	err := c.sendPostRequest(ctx, endpoint, payload, nil)

	return err
}
