package model

type BlockchainEvent struct {
	EventType BlockchainEventType
	EventInfo interface{}
}

type BlockchainEventType string

var (
	TransferBlockchainEventType         BlockchainEventType = "NftTransferred"
	NftMintedBlockchainEventType        BlockchainEventType = "NftMinted"
	NftPriceChangedBlockchainEvent      BlockchainEventType = "NftPriceChanged"
	NftCollectionCreatedBlockchainEvent BlockchainEventType = "NftCollectionCreated"
)
