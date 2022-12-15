package model

type Event struct {
	EventType EventType
	ObjectID  int
	SubjectID *string
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
	NewCreatorEvent                  EventType = "NewCreator"
	NewNFTMintedEvent                EventType = "NewNFTMinted"
	MintedNftUpdateEvent             EventType = "MintedNftUpdate"
	NewScheduledNftMintEvent         EventType = "NewScheduledNftMint"
	ScheduledNftUpdateEvent          EventType = "ScheduledNftUpdate"
	ScheduledNftDeleteEvent          EventType = "ScheduledNftDelete"
	NFTSoldEvent                     EventType = "NftSold"
	ScheduledEmailEvent              EventType = "NewScheduledEmail"
	ShowCreatedEvent                 EventType = "NewShowCreated"
	ShowPublicEvent                  EventType = "ShowBecomePublic"
)
