package model

import "time"

type BlockchainEvent struct {
	EventType       BlockchainEventType
	ParsedEventInfo interface{}
	BlockchainEvent interface{}
	PickedLive      bool
}

type BlockchainEventType string

var (
	TransferBlockchainEventType         BlockchainEventType = "NftTransferred"
	NftMintedBlockchainEventType        BlockchainEventType = "NftMinted"
	NftPriceChangedBlockchainEvent      BlockchainEventType = "NftPriceChanged"
	NftCollectionCreatedBlockchainEvent BlockchainEventType = "NftCollectionCreated"
	NewWaxBlockchainEvent               BlockchainEventType = "NewWaxEvent"
)

type CollectionCreatedEventInfo struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	CreatorAddress string `json:"creatorAddress"`
	BlockNumber    int    `json:"blockNumber"`
	ShowID         int    `json:"showId"`
}

type NftMintedEventInfo struct {
	NftID              int     `json:"nftId"`
	TokenID            int     `json:"tokenId"`
	StoreID            int     `json:"storeId"`
	Price              float64 `json:"price"`
	Fee                float64 `json:"fee"`
	BlockNumber        int     `json:"blockNumber"`
	Balance            int     `json:"balance"`
	Supply             int     `json:"supply"`
	NftContractAddress string  `json:"nftContractAddress"`
	CreatorAddress     string  `json:"creatorAddress"`
	MintedOn           int     `json:"mintedOn"`
	Currency           string  `json:"currency"`
	Description        string  `json:"description"`
	ImageURL           string  `json:"imageUrl"`
	AudioURL           string  `json:"audioUrl"`
	MetadataURL        string  `json:"metadataUrl"`
}

type NftPriceChangedEventInfo struct {
	NftID int     `json:"nftId"`
	Price float64 `json:"price"`
}

type TransferEventInfo struct {
	TokenID           int       `json:"tokenId"`
	CollectionAddress string    `json:"collectionAddress"`
	OldOwnerAddress   string    `json:"oldOwnerAddress"`
	NewOwnerAddress   string    `json:"newOwnerAddress"`
	BlockNumber       int       `json:"blockNumber"`
	Quantity          int       `json:"quantity"`
	Transaction       string    `json:"transaction"`
	IsBuyEvent        bool      `json:"isBuyEvent"`
	TransferredAt     time.Time `json:"transferredAt"`
}

type BlockchainRequest struct {
	RequestType BlockchainRequestType
	ObjectID    int
	SubjectID   *string
}

type BlockchainRequestType string

var (
	BlockchainRequestMintNft                 BlockchainRequestType = "MintNFT"
	BlockchainRequestBurnNft                 BlockchainRequestType = "BurnNFT"
	BlockchainRequestMintCollection          BlockchainRequestType = "MintCollection"
	BlockchainRequestMintWaxCollection       BlockchainRequestType = "MintWaxCollection"
	BlockchainRequestMintWaxSchema           BlockchainRequestType = "MintWaxSchema"
	BlockchainRequestUpdateWaxCollectionData BlockchainRequestType = "UpdateWaxCollectionData"
	BlockchainRequestMintWaxNft              BlockchainRequestType = "MintWaxNft"
)

type TranscoderPubsubRequest struct {
	RequestType        TranscoderPubsubRequestType
	NftTokenID         string
	NftContractAddress string
	VideoURL           string
	ImageURL           string
}

type TranscoderPubsubRequestType string

var (
	ExtractVideoFirstFrame TranscoderPubsubRequestType = "ExtractVideoFirstFrame"
	ProcessImageMetadata   TranscoderPubsubRequestType = "ImageMetadata"
)
