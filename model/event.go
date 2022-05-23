package model

type Event struct {
	EventType EventType
	ObjectID  int
	SubjectID *int
}

type EventType string

var (
	NewSpaceConversationEvent EventType = "NewSpaceConversation"
	FreeNftTransferEvent      EventType = "FreeNftTransfer"
	PaidNftTransferEvent      EventType = "PaidNftTransfer"
)
