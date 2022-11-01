package model

import (
	"encoding/json"
	"github.com/uncut-fm/uncut-management-api-2/ent/notification"
	"time"
)

type NewNotification struct {
	ReceiverEmail string
	CategoryType  notification.CategoryType
	TemplateType  NotificationTemplateType
	Status        notification.Status
	ScheduledFor  time.Time
	SentAt        *time.Time
	Metadata      map[string]interface{}
}

type NotificationTemplateType string

const (
	NftTransferCompletedNotification NotificationTemplateType = "NFT_TRANSFER_COMPLETED"
	NftSoldNotification              NotificationTemplateType = "NFT_SOLD"
	NftWelcomeNotification           NotificationTemplateType = "NFT_WELCOME"
	NftAirdropNotification           NotificationTemplateType = "NFT_AIRDROP"

	ConversationSpaceNotification            NotificationTemplateType = "SPACE_CONVERSATION"
	ConversationCommentNotification          NotificationTemplateType = "CONVERSATION_COMMENT"
	ConversationReplyCommentatorNotification NotificationTemplateType = "CONVERSATION_REPLY_TO_COMMENTATOR"
	ConversationReplyHostNotification        NotificationTemplateType = "CONVERSATION_REPLY_TO_HOST"
)

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
