package model

type ConversationEmailRequest struct {
	PodcastName         string
	PodcastSlug         string
	CreatorName         string
	ObjectName          string
	ObjectIdentifier    string
	ConversationMessage string
}

type EmailReceiver struct {
	Email string
	Name  string
}
