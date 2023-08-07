package storage

type EntityType string

const (
	EntityTypeNft          EntityType = "NFT"
	EntityTypeConversation EntityType = "Conversation"
	EntityTypeUser         EntityType = "User"
	EntityTypeSpace        EntityType = "Space"
	EntityTypeCollection   EntityType = "Collection"
	EntityTypeShow         EntityType = "Show"
)
