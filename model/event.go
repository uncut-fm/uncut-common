package model

import (
	"strconv"
	"time"
)

type Event struct {
	EventType EventType
	ObjectID  int
	SubjectID *string
	Metadata  interface{}
}

type EventType string

func (e EventType) String() string {
	return string(e)
}

func (e Event) GetUserID() int {
	if e.SubjectID == nil || *e.SubjectID == "" {
		return 0
	}

	userID, _ := strconv.Atoi(*e.SubjectID)
	return userID
}

var (
	NewSpaceConversationEvent        EventType = "NewSpaceConversation"
	NewNFTConversationEvent          EventType = "NewNFTConversation"
	FreeNftTransferEvent             EventType = "FreeNftTransfer"
	PaidNftTransferEvent             EventType = "PaidNftTransfer"
	FaucetWalletBalanceAlertEvent    EventType = "FaucetWalletBalanceAlert"
	WelcomingEmailEvent              EventType = "WelcomingEmail"
	NftAirdropEmailEvent             EventType = "AirdropEmail"
	NewSpaceConversationCommentEvent EventType = "NewSpaceConversationComment"
	NewSpaceConversationReplyEvent   EventType = "NewSpaceConversationReply"
	NewMentionFeedPostReplyEvent     EventType = "NewMentionFeedPost"
	NewFeedConversationCommentEvent  EventType = "NewFeedConversationComment"

	NewPostEvent         EventType = "NewPost"
	PostDeletedEvent     EventType = "PostDeleted"
	NewBlogPostEvent     EventType = "NewBlogPost"
	BlogPostDeletedEvent EventType = "BlogPostDeleted"
	NewCommentEvent      EventType = "NewComment"
	CommentDeletedEvent  EventType = "CommentDeleted"

	NewCreatorEvent      EventType = "NewCreator"
	NewUserEvent         EventType = "NewUser"
	NewReferredUserEvent EventType = "NewReferredUser"
	UserDeletedEvent     EventType = "UserDeleted"
	UserUpdatedEvent     EventType = "UserUpdated"
	UserLoggedInEvent    EventType = "UserLoggedIn"

	UserLikedPostEvent      EventType = "UserLikedPost"
	UserUnlikedPostEvent    EventType = "UserUnlikedPost"
	UserLikedCommentEvent   EventType = "UserLikedComment"
	UserUnlikedCommentEvent EventType = "UserUnlikedComment"
	UserLikedNftEvent       EventType = "UserLikedNft"
	UserUnlikedNftEvent     EventType = "UserUnlikedNft"

	UserFollowedUserEvent   EventType = "UserFollowedUser"
	UserUnfollowedUserEvent EventType = "UserUnfollowedUser"

	WalletDeletedEvent EventType = "WalletDeleted"
	WalletAddedEvent   EventType = "WalletAdded"

	UserAddedOrRemovedTagsEvent EventType = "UserAddedOrRemovedTags"

	NewNFTMintedEvent        EventType = "NewNFTMinted"
	MintedNftUpdateEvent     EventType = "MintedNftUpdate"
	NewScheduledNftMintEvent EventType = "NewScheduledNftMint"
	ScheduledNftUpdateEvent  EventType = "ScheduledNftUpdate"
	ScheduledNftDeleteEvent  EventType = "ScheduledNftDelete"
	NFTSoldEvent             EventType = "NftSold"
	NftDeletedEvent          EventType = "NftDeleted"
	NftAirdropEvent          EventType = "NftAirdrop"

	NewPaperCheckoutEvent EventType = "NewPaperCheckout"

	ScheduledEmailEvent EventType = "NewScheduledEmail"
	ShowCreatedEvent    EventType = "NewShowCreated"
	ShowPublicEvent     EventType = "ShowBecomePublic"

	NftCuratedEvent           EventType = "NftCurated"
	NftCurationUpdatedEvent   EventType = "NftCurationUpdated"
	NftCurationDeletedEvent   EventType = "NftCurationDeleted"
	CuratedListFinalizedEvent EventType = "CuratedListFinalized"

	NewArtxTransferEvent         EventType = "NewArtxTransfer"
	ReportNFTEvent               EventType = "ReportNFT"
	OnboardingStepCompletedEvent EventType = "OnboardingStepCompleted"
	SubscriptionUpdatedEvent     EventType = "SubscriptionUpdated"
)

type OnboardingStepCompletedEventMetadata struct {
	UserEmail           string                   `json:"user_email"`
	TemplateType        NotificationTemplateType `json:"template_type"`
	StepNumber          int                      `json:"step_number"`
	TotalStepsCompleted int                      `json:"total_steps_completed"`
	OnboardingStatus    OnboardingStatus         `json:"onboarding_status"`
}

type LastLoggedInEventMetadata struct {
	LastLoggedInAt time.Time `json:"last_logged_in_at"`
}

type AirdropEventMetadata struct {
	UserEmails    []string          `json:"user_emails"`
	UserIDs       []int             `json:"user_ids"`
	CollectibleID int               `json:"collectible_id"`
	Blockchain    BlockchainNetwork `json:"blockchain"`
}
