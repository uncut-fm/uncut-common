package hubspot

import (
	"fmt"
	"time"
)

type (
	HubspotObjectType          string
	HubspotAssociationType     string
	HubspotAssociationCategory string
	HubpsotAssociationTypeID   int
)

var (
	ContactsObjectType HubspotObjectType = "contacts"
	TagsObjectType     HubspotObjectType = "tags"
	BadgesObjectType   HubspotObjectType = "badges"
	DealsObjectType    HubspotObjectType = "deals"

	ContactToTagAssociationType   HubspotAssociationType = "user_tags"
	EAPBadgeName                                         = "EAP"
	ContactToBadgeAssociationType HubspotAssociationType = "user_badges"

	DealToContactAssociationTypeID HubpsotAssociationTypeID = 3

	HubspotDefinedAssociationCategory HubspotAssociationCategory = "HUBSPOT_DEFINED"

	DefaultHubspotDealsOwnerID          = "562479474"
	FeaturedArtOfTheDayPipelineID       = "88439581"
	FeaturedArtOfTheDayScheduledStageID = "164285780"
)

type HubspotSimplePublicObjectInput struct {
	ID         string            `json:"id,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type CreateObjectInput struct {
	Properties   map[string]string     `json:"properties"`
	Associations []AssociationToObject `json:"associations,omitempty"`
}

type UpdateObjectInput struct {
	ObjectID   string            `json:"-"`
	Properties map[string]string `json:"properties"`
}

type HubspotAssociationTypeObject struct {
	AssociationCategory HubspotAssociationCategory `json:"associationCategory"`
	AssociationTypeId   HubpsotAssociationTypeID   `json:"associationTypeId"`
}

type AssociationPairToArrayObject struct {
	From IDObject `json:"from"`
	To   []struct {
		ID   string                 `json:"id"`
		Type HubspotAssociationType `json:"type"`
	}
}

type IDObject struct {
	ID string `json:"id"`
}

type AssociationPair struct {
	From IDObject               `json:"from"`
	To   IDObject               `json:"to"`
	Type HubspotAssociationType `json:"type"`
}

type AssociationToObject struct {
	Types []HubspotAssociationTypeObject `json:"types"`
	To    IDObject                       `json:"to"`
}

type NewProperty struct {
	Name      string            `json:"name"`
	Label     string            `json:"label"`
	Type      PropertyType      `json:"type"`
	FieldType PropertyFieldType `json:"fieldType"`
	GroupName PropertyGroupName `json:"groupName"`
	Options   []PropertyOption  `json:"options,omitempty"`
}

type PropertyType string

const (
	PropertyTypeString      PropertyType = "string"
	PropertyTypeNumber      PropertyType = "number"
	PropertyTypeDatetime    PropertyType = "datetime"
	PropertyTypeEnumeration PropertyType = "enumeration"
	PropertyTypeBool        PropertyType = "bool"
)

type PropertyFieldType string

const (
	PropertyFieldTypeNumber          PropertyFieldType = "number"
	PropertyFieldTypeText            PropertyFieldType = "text"
	PropertyFieldTypeDate            PropertyFieldType = "date"
	PropertyFieldTypeRadio           PropertyFieldType = "radio"
	PropertyFieldTypeFile            PropertyFieldType = "file"
	PropertyFieldTypeBooleanCheckbox PropertyFieldType = "booleancheckbox"
)

type PropertyOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PropertyGroupName string

const (
	PropertyGroupNameSocialMedia PropertyGroupName = "socialmediainformation"
	PropertyGroupNameContactInfo PropertyGroupName = "contactinformation"
	PropertyGroupArtxInfo        PropertyGroupName = "artx_information"
)

func GetTimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}

type FeaturedArtOfTheDayDealInput struct {
	UserName      string
	UserHubspotID string
	ScheduledDate time.Time
}

var featuredArtOfTheDayDealNamePattern = "%s - %s" // {user_name} - {scheduled_data}

func NewFeaturedArtOfTheDayDealCreateObjectInput(input FeaturedArtOfTheDayDealInput) CreateObjectInput {
	return CreateObjectInput{
		Associations: []AssociationToObject{
			{
				Types: []HubspotAssociationTypeObject{
					{
						AssociationCategory: HubspotDefinedAssociationCategory,
						AssociationTypeId:   DealToContactAssociationTypeID,
					},
				},
				To: IDObject{
					ID: input.UserHubspotID,
				},
			},
		},
		Properties: newFeaturedArtOfTheDayDealPropertiesInput(input),
	}
}

func NewFeaturedArtOfTheDayDealUpdateObjectInput(input FeaturedArtOfTheDayDealInput, dealID string) UpdateObjectInput {
	return UpdateObjectInput{
		ObjectID:   dealID,
		Properties: newFeaturedArtOfTheDayDealPropertiesInput(input),
	}
}

func newFeaturedArtOfTheDayDealPropertiesInput(input FeaturedArtOfTheDayDealInput) map[string]string {
	timeString := GetTimeString(input.ScheduledDate)

	return map[string]string{
		"dealname":         fmt.Sprintf(featuredArtOfTheDayDealNamePattern, input.UserName, timeString),
		"pipeline":         FeaturedArtOfTheDayPipelineID,
		"dealstage":        FeaturedArtOfTheDayScheduledStageID,
		"hubspot_owner_id": DefaultHubspotDealsOwnerID,
		"closedate":        timeString,
	}
}
