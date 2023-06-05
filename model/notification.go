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
	UserID        *int
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
	OnboardingDay4Notification               NotificationTemplateType = "DAY_4"
	OnboardingDay7Notification               NotificationTemplateType = "DAY_7"
	OnboardingDay10Notification              NotificationTemplateType = "DAY_10"
	OnboardingDay15Notification              NotificationTemplateType = "DAY_15"
	OnboardingDay20Notification              NotificationTemplateType = "DAY_20"
	OnboardingDay30Notification              NotificationTemplateType = "DAY_30"
	OnboardingUserHour1Notification          NotificationTemplateType = "HOUR_1"
	OnboardingUserHour2Notification          NotificationTemplateType = "HOUR_2"
	ConversationSpaceNotification            NotificationTemplateType = "SPACE_CONVERSATION"
	ConversationCommentNotification          NotificationTemplateType = "CONVERSATION_COMMENT"
	ConversationReplyCommentatorNotification NotificationTemplateType = "CONVERSATION_REPLY_TO_COMMENTATOR"
	ConversationReplyHostNotification        NotificationTemplateType = "CONVERSATION_REPLY_TO_HOST"

	FeedPostUserMentionNotification    NotificationTemplateType = "FEED_POST_USER_MENTION"
	FeedCommentUserMentionNotification NotificationTemplateType = "FEED_COMMENT_USER_MENTION"
	FeedPostNftMentionNotification     NotificationTemplateType = "FEED_POST_NFT_MENTION"
	FeedPostReplyNotification          NotificationTemplateType = "FEED_POST_REPLY"
	FeedCommentReplyNotification       NotificationTemplateType = "FEED_COMMENT_REPLY"

	WelcomeEmailSetupNotification NotificationTemplateType = "WELCOME_EMAIL_SETUP"
	WelcomeOwnerUncutNotification NotificationTemplateType = "WELCOME_OWNER_UNCUT"
	MintFirstTokenTipNotification NotificationTemplateType = "MINT_FIRST_TOKEN"
)

var (
	BlogNotificationTemplates        = []NotificationTemplateType{ConversationSpaceNotification, ConversationCommentNotification, ConversationReplyCommentatorNotification, ConversationReplyHostNotification}
	FeedNotificationTemplates        = []NotificationTemplateType{FeedPostUserMentionNotification, FeedCommentUserMentionNotification, FeedPostNftMentionNotification, FeedPostReplyNotification, FeedCommentReplyNotification}
	OnboardingUserSequenceTemplates  = []NotificationTemplateType{OnboardingUserHour1Notification, OnboardingUserHour2Notification, OnboardingDay1Notification, OnboardingDay2Notification, OnboardingDay4Notification, OnboardingDay7Notification, OnboardingDay10Notification, OnboardingDay15Notification, OnboardingDay20Notification, OnboardingDay30Notification}
	TransactionNotificationTemplates = []NotificationTemplateType{NftTransferCompletedNotification, NftSoldNotification, NftWelcomeNotification, NftAirdropNotification}
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
	UserID       *int   `json:"userId,omitempty"`
}
