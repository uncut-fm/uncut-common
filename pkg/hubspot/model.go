package hubspot

type (
	HubspotObjectType      string
	HubspotAssociationType string
)

var (
	ContactsObjectType          HubspotObjectType      = "contacts"
	TagsObjectType              HubspotObjectType      = "tags"
	ContactToTagAssociationType HubspotAssociationType = "user_tags"

	BadgesObjectType              HubspotObjectType      = "badges"
	EAPBadgeName                                         = "EAP"
	ContactToBadgeAssociationType HubspotAssociationType = "user_badges"
)

type HubspotSimplePublicObjectInput struct {
	ID         string            `json:"id,omitempty"`
	Properties map[string]string `json:"properties"`
}

type AssociationObject struct {
	From struct {
		ID string `json:"id"`
	} `json:"from"`
	To []struct {
		ID   string                 `json:"id"`
		Type HubspotAssociationType `json:"type"`
	}
}

type AssociationPair struct {
	From struct {
		ID string `json:"id"`
	} `json:"from"`
	To struct {
		ID string `json:"id"`
	} `json:"to"`
	Type HubspotAssociationType `json:"type"`
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
)
