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

	NewCreatorEvent    EventType = "NewCreator"
	NewUserEvent       EventType = "NewUser"
	UserDeletedEvent   EventType = "UserDeleted"
	WalletDeletedEvent EventType = "WalletDeleted"
	WalletAddedEvent   EventType = "WalletAdded"

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
)
