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
	NewSpaceConversationCommentEvent EventType = "NewSpaceConversationComment"
	NewSpaceConversationReplyEvent   EventType = "NewSpaceConversationReply"
	NewCreatorEvent                  EventType = "NewCreator"
)
