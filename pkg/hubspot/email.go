package hubspot

import (
	"context"
)

var sendEmailEndpoint = "https://api.hubapi.com/marketing/v3/transactional/single-email/send"

type NewEmailRequest struct {
	HubspotEmailID int `json:"emailId"`
	Message        struct {
		From    *string  `json:"from"`
		To      *string  `json:"to"`
		SendID  *string  `json:"sendId"`
		ReplyTo []string `json:"replyTo"`
	} `json:"message"`
	CustomProperties map[string]interface{} `json:"customProperties"`
}

func (c *Client) SendEmail(ctx context.Context, request *NewEmailRequest) error {
	resp := &sendEmailResponse{}

	err := c.sendPostRequest(sendEmailEndpoint, request, resp)
	if err != nil {
		return err
	}

	return nil
}
