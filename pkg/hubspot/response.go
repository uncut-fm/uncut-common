package hubspot

type hubspotBatchObjectsResponseV3 struct {
	Status        hubspotStatus                    `json:"status"`
	Message       string                           `json:"message"`
	CorrelationID string                           `json:"correlationId"`
	Category      string                           `json:"category"`
	Results       []HubspotSimplePublicObjectInput `json:"results"`
}

type hubspotSearchObjectsResponseV3 struct {
	Total   int                              `json:"total"`
	Results []HubspotSimplePublicObjectInput `json:"results"`
}

type hubspotObjectResponseV3 struct {
	Status        hubspotStatus `json:"status"`
	Message       string        `json:"message"`
	CorrelationID string        `json:"correlationId"`
	Category      string        `json:"category"`
	HubspotSimplePublicObjectInput
}

func (h hubspotBatchObjectsResponseV3) GetStatus() hubspotStatus {
	return h.Status
}

func (h hubspotBatchObjectsResponseV3) GetMessage() string {
	return h.Message
}

func (h hubspotObjectResponseV3) GetStatus() hubspotStatus {
	return h.Status
}

func (h hubspotObjectResponseV3) GetMessage() string {
	return h.Message
}

func (h hubspotSearchObjectsResponseV3) GetStatus() hubspotStatus {
	return hubspotStatusComplete
}

func (h hubspotSearchObjectsResponseV3) GetMessage() string {
	return ""
}

type associationResponseErr struct {
	Message string `json:"message"`
	Context struct {
		FromObjectID []string `json:"fromObjectId"`
	}
}

type hubspotReadAssociationsResponseV3 struct {
	Status  hubspotStatus                  `json:"status"`
	Errors  []associationResponseErr       `json:"errors"`
	Results []AssociationPairToArrayObject `json:"results"`
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
			h.Results = append(h.Results, AssociationPairToArrayObject{
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
	hubspotStatusPending  hubspotStatus = "PENDING"
	hubspotStatusComplete hubspotStatus = "COMPLETE"
)

type sendEmailResponse struct {
	Status     hubspotStatus `json:"status"`
	SendResult string        `json:"sendResult"`
}

func (s *sendEmailResponse) GetStatus() hubspotStatus {
	return s.Status
}

func (s *sendEmailResponse) GetMessage() string {
	return s.SendResult
}
