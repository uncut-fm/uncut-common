package model

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
}

type NftPriceChangedEventInfo struct {
	NftID int     `json:"nftId"`
	Price float64 `json:"price"`
}

type TransferEventInfo struct {
	TokenID           int    `json:"tokenId"`
	CollectionAddress string `json:"collectionAddress"`
	OldOwnerAddress   string `json:"oldOwnerAddress"`
	NewOwnerAddress   string `json:"newOwnerAddress"`
	BlockNumber       int    `json:"blockNumber"`
	Quantity          int    `json:"quantity"`
	Transaction       string `json:"transaction"`
}

type BlockchainRequest struct {
	RequestType BlockchainRequestType
	ObjectID    int
}

type BlockchainRequestType string

var (
	BlockchainRequestMintNft BlockchainRequestType = "MintNFT"
)
