package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

var unsubscribeEndpointPattern = "%s/unsubscribe/?req=%s"

func GenerateUnsubscribeEndpointURL(managementBaseURL string, request UnsubscribeHTTPRequest) (string, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return "", err
	}

	requestEncrypted := base64.StdEncoding.EncodeToString(requestBytes)

	return fmt.Sprintf(unsubscribeEndpointPattern, managementBaseURL, requestEncrypted), nil
}

type NewNotification struct {
	ReceiverEmail string
	CategoryType  string
	TemplateType  NotificationTemplateType
	Status        string
	ScheduledFor  *time.Time
	SentAt        *time.Time
	Metadata      map[string]interface{}
}

type NewSubscriptionNotification struct {
	ReceiverEmail string
	CategoryType  string
	TemplateType  NotificationTemplateType
	SpaceID       *int
}

type NotificationTemplateType string

const (
	ConversationSpaceNotification            NotificationTemplateType = "SPACE_CONVERSATION"
	ConversationCommentNotification          NotificationTemplateType = "CONVERSATION_COMMENT"
	ConversationReplyCommentatorNotification NotificationTemplateType = "CONVERSATION_REPLY_TO_COMMENTATOR"
	ConversationReplyHostNotification        NotificationTemplateType = "CONVERSATION_REPLY_TO_HOST"
)

var ConversationNotificationTemplates = []NotificationTemplateType{ConversationSpaceNotification, ConversationCommentNotification, ConversationReplyCommentatorNotification, ConversationReplyHostNotification}

func (n NotificationTemplateType) String() string {
	return string(n)
}

func (n *NewNotification) SetMetadataFromTemplateVariables(templateVariables interface{}) error {
	templateBytes, err := json.Marshal(templateVariables)
	if err != nil {
		return err
	}

	err = json.Unmarshal(templateBytes, &n.Metadata)

	return err
}

type UnsubscribeHTTPRequest struct {
	JWT          string `json:"token"`
	CategoryType string `json:"categoryType"`
	SpaceID      *int   `json:"spaceId,omitempty"`
}
