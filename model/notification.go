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

func NewGrantArtxMetadata(userID, artxAmount int) map[string]interface{} {
	return map[string]interface{}{
		"userId":      userID,
		"artxGranted": artxAmount,
	}
}

func GetArtxAmountFromMetadata(metadata map[string]interface{}) (int, bool) {
	if artxGranted, ok := metadata["artxGranted"]; ok {
		return int(artxGranted.(float64)), true
	}

	return 0, false
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
	NftTransferCompletedNotification NotificationTemplateType = "NFT_TRANSFER_COMPLETED"
	NftSoldNotification              NotificationTemplateType = "NFT_SOLD"
	NftWelcomeNotification           NotificationTemplateType = "NFT_WELCOME"
	NftAirdropNotification           NotificationTemplateType = "NFT_AIRDROP"

	UsersNftMintedNotification           NotificationTemplateType = "USERS_NFT_MINTED"
	CreatorMintedNftNotification         NotificationTemplateType = "CREATOR_MINTED_NFT"
	CreatorCreatedCollectionNotification NotificationTemplateType = "CREATOR_CREATED_COLLECTION"

	OnboardingDay0Notification    NotificationTemplateType = "DAY_0"
	OnboardingNewUserNotification NotificationTemplateType = "NEW_USER"

	ConversationSpaceNotification            NotificationTemplateType = "SPACE_CONVERSATION"
	ConversationCommentNotification          NotificationTemplateType = "CONVERSATION_COMMENT"
	ConversationReplyCommentatorNotification NotificationTemplateType = "CONVERSATION_REPLY_TO_COMMENTATOR"
	ConversationReplyHostNotification        NotificationTemplateType = "CONVERSATION_REPLY_TO_HOST"

	FeedPostNotification                     NotificationTemplateType = "FEED_POST"
	FeedPostUserMentionNotification          NotificationTemplateType = "FEED_POST_USER_MENTION"
	FeedCommentUserMentionNotification       NotificationTemplateType = "FEED_COMMENT_USER_MENTION"
	FeedPostNftMentionNotification           NotificationTemplateType = "FEED_POST_NFT_MENTION"
	FeedCommentNftMentionNotification        NotificationTemplateType = "FEED_COMMENT_NFT_MENTION"
	FeedPostCollectionMentionNotification    NotificationTemplateType = "FEED_POST_COLLECTION_MENTION"
	FeedCommentCollectionMentionNotification NotificationTemplateType = "FEED_COMMENT_COLLECTION_MENTION"
	FeedPostReplyNotification                NotificationTemplateType = "FEED_POST_REPLY"
	FeedCommentReplyNotification             NotificationTemplateType = "FEED_COMMENT_REPLY"
	FeedPostRepostNotification               NotificationTemplateType = "FEED_POST_REPOST"
	FeedNftHolderPostNotification            NotificationTemplateType = "FEED_NFT_HOLDER_POST"
	FeedCollectionHolderPostNotification     NotificationTemplateType = "FEED_COLLECTION_HOLDER_POST"

	SocialPostLikeNotification         NotificationTemplateType = "FEED_POST_LIKE"
	SocialCommentLikeNotification      NotificationTemplateType = "FEED_COMMENT_LIKE"
	SocialCommentReplyLikeNotification NotificationTemplateType = "FEED_COMMENT_REPLY_LIKE"
	SocialNftLikeNotification          NotificationTemplateType = "NFT_LIKE"
	SocialFollowNotification           NotificationTemplateType = "USER_FOLLOW"

	UserVerificationStartedNotification  NotificationTemplateType = "USER_VERIFICATION_STARTED"
	UserVerificationAcceptedNotification NotificationTemplateType = "USER_VERIFICATION_ACCEPTED"
	UserVerificationRefusedNotification  NotificationTemplateType = "USER_VERIFICATION_REFUSED"

	WelcomeEmailSetupNotification NotificationTemplateType = "WELCOME_EMAIL_SETUP"
	WelcomeOwnerUncutNotification NotificationTemplateType = "WELCOME_OWNER_UNCUT"
	MintFirstTokenTipNotification NotificationTemplateType = "MINT_FIRST_TOKEN"

	FraudCollectionOwnerNotification NotificationTemplateType = "FRAUD_COLLECTION_OWNER"

	GamificationNewLevelReachedNotification NotificationTemplateType = "NEW_LEVEL_REACHED"

	ArtxIntroductionProductUpdateNotification NotificationTemplateType = "UPDATE_ARTX_0_MODAL"
	ArtxPlatformGrantNotification             NotificationTemplateType = "ARTX_PLATFORM_GRANT"
	WelcomeArtxBonusNotification              NotificationTemplateType = "WELCOME_ARTX_BONUS"
)

var (
	BlogNotificationTemplates        = []NotificationTemplateType{ConversationSpaceNotification, ConversationCommentNotification, ConversationReplyCommentatorNotification, ConversationReplyHostNotification}
	FeedNotificationTemplates        = []NotificationTemplateType{FeedPostNotification, FeedPostUserMentionNotification, FeedCommentUserMentionNotification, FeedPostNftMentionNotification, FeedPostReplyNotification, FeedCommentReplyNotification, FeedCommentNftMentionNotification}
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
