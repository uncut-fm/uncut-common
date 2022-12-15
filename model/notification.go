package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

var unsubscribeEndpointPattern = "%s/user/unsubscribe/?req=%s"

func GenerateUnsubscribeURL(appBaseUrl string, request UnsubscribeHTTPRequest) (string, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return "", err
	}

	requestEncrypted := base64.StdEncoding.EncodeToString(requestBytes)

	return fmt.Sprintf(unsubscribeEndpointPattern, appBaseUrl, requestEncrypted), nil
}

type NewNotification struct {
	ReceiverEmail string
	CategoryType  string
	TemplateType  NotificationTemplateType
	Channel       string
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
	NftTransferCompletedNotification         NotificationTemplateType = "NFT_TRANSFER_COMPLETED"
	NftSoldNotification                      NotificationTemplateType = "NFT_SOLD"
	NftWelcomeNotification                   NotificationTemplateType = "NFT_WELCOME"
	NftAirdropNotification                   NotificationTemplateType = "NFT_AIRDROP"
	OnboardingDay0Notification               NotificationTemplateType = "DAY_0"
	OnboardingDay1Notification               NotificationTemplateType = "DAY_1"
	OnboardingDay2Notification               NotificationTemplateType = "DAY_2"
	OnboardingDay5Notification               NotificationTemplateType = "DAY_5"
	OnboardingDay10Notification              NotificationTemplateType = "DAY_10"
	ConversationSpaceNotification            NotificationTemplateType = "SPACE_CONVERSATION"
	ConversationCommentNotification          NotificationTemplateType = "CONVERSATION_COMMENT"
	ConversationReplyCommentatorNotification NotificationTemplateType = "CONVERSATION_REPLY_TO_COMMENTATOR"
	ConversationReplyHostNotification        NotificationTemplateType = "CONVERSATION_REPLY_TO_HOST"
	WelcomeEmailSetupNotification            NotificationTemplateType = "WELCOME_EMAIL_SETUP"
)

var (
	ConversationNotificationTemplates = []NotificationTemplateType{ConversationSpaceNotification, ConversationCommentNotification, ConversationReplyCommentatorNotification, ConversationReplyHostNotification}
	OnboardingSequenceTemplates       = []NotificationTemplateType{OnboardingDay1Notification, OnboardingDay2Notification, OnboardingDay5Notification, OnboardingDay10Notification}
	TransactionNotificationTemplates  = []NotificationTemplateType{NftTransferCompletedNotification, NftSoldNotification, NftWelcomeNotification, NftAirdropNotification}
)

func (n NotificationTemplateType) String() string {
	return string(n)
}

func (n NotificationTemplateType) GetPointer() *NotificationTemplateType {
	return &n
}

func (n *NewNotification) SetMetadataFromTemplateVariables(templateVariables interface{}) error {
	templateBytes, err := json.Marshal(templateVariables)
	if err != nil {
		return err
	}

	err = json.Unmarshal(templateBytes, &n.Metadata)

	return err
}

type UpdateNotification struct {
	ID       int
	Status   string
	SentAt   *time.Time
	Metadata map[string]interface{}
}

func (u *UpdateNotification) SetMetadataFromTemplateVariables(templateVariables interface{}) error {
	templateBytes, err := json.Marshal(templateVariables)
	if err != nil {
		return err
	}

	err = json.Unmarshal(templateBytes, &u.Metadata)

	return err
}

type UnsubscribeHTTPRequest struct {
	JWT          string `json:"token"`
	CategoryType string `json:"categoryType"`
	SpaceID      *int   `json:"spaceId,omitempty"`
}
