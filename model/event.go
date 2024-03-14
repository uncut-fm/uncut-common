package model

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

	NewCreatorEvent   EventType = "NewCreator"
	NewUserEvent      EventType = "NewUser"
	UserDeletedEvent  EventType = "UserDeleted"
	UserUpdatedEvent  EventType = "UserUpdated"
	UserLoggedInEvent EventType = "UserLoggedIn"

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

	NewPaperCheckoutEvent EventType = "NewPaperCheckout"

	ScheduledEmailEvent EventType = "NewScheduledEmail"
	ShowCreatedEvent    EventType = "NewShowCreated"
	ShowPublicEvent     EventType = "ShowBecomePublic"

	NftCuratedEvent         EventType = "NftCurated"
	NftCurationUpdatedEvent EventType = "NftCurationUpdated"
)
