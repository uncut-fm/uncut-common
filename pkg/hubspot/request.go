package hubspot

type batchReadRequest struct {
	Properties []string            `json:"properties"`
	IdProperty *string             `json:"idProperty,omitempty"`
	Inputs     []map[string]string `json:"inputs"`
}

type searchRequest struct {
	Properties   []string            `json:"properties,omitempty"`
	FilterGroups []searchFilterGroup `json:"filterGroups"`
}

type searchFilter struct {
	PropertyName string `json:"propertyName"`
	Operator     string `json:"operator"`
	Value        string `json:"value"`
}

type searchFilterGroup struct {
	Filters []searchFilter `json:"filters"`
}

func getDealStageFromDealProperties(properties map[string]string) string {
	return properties["dealstage"]
}

func IsFeaturedArtistDealProcessed(properties map[string]string) bool {
	return getDealStageFromDealProperties(properties) == FeaturedArtistSignupUpStageID
}
