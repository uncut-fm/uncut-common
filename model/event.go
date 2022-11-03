package model

type Event struct {
	EventType EventType
	ObjectID  int
	SubjectID *string
}

type EventType string

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
	NFTSoldEvent                     EventType = "NftSold"
	MintedNftUpdateEvent             EventType = "MintedNftUpdate"
	ScheduledEmailEvent              EventType = "NewScheduledEmail"
	ShowCreatedEvent                 EventType = "NewShowCreatedEvent"
)
