package model

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	NewWaxPurchasedSaleEvent            BlockchainEventType = "NewWaxPurchasedSaleEvent"
	NewWaxAssetTransferEvent            BlockchainEventType = "NewWaxAssetTransferEvent"
)

type PurchasedSaleEventInfo struct {
	SaleID string `json:"saleId"`
}

type WaxTransferTransactionEventInfo struct {
	TransactionID int `json:"transactionId"`
}

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
	Metadata    BlockchainRequestMetadata
}

type BlockchainRequestMetadata map[string]interface{}

func NewUpdateCollectionWalletBlockchainRequestMetadata(oldWallet, newWallet string) BlockchainRequestMetadata {
	return BlockchainRequestMetadata{
		"oldWallet": oldWallet,
		"newWallet": newWallet,
	}
}

func (b BlockchainRequestMetadata) GetUpdateCollectionWalletOldWallet() (string, bool) {
	oldWallet, ok := b["oldWallet"].(string)

	return oldWallet, ok
}

func (b BlockchainRequestMetadata) GetUpdateCollectionWalletNewWallet() (string, bool) {
	newWallet, ok := b["newWallet"].(string)

	return newWallet, ok
}

func NewMintWaxAssetBlockchainRequestMetadata(newOwner string, copies int) BlockchainRequestMetadata {
	return BlockchainRequestMetadata{
		"newOwner": newOwner,
		"copies":   copies,
	}
}

func (b BlockchainRequestMetadata) GetMintWaxAssetNewOwner() (string, bool) {
	newOwner, ok := b["newOwner"].(string)

	return newOwner, ok
}

func (b BlockchainRequestMetadata) GetMintWaxAssetCopies() (int, bool) {
	if copies, ok := b["copies"]; ok {
		switch copies.(type) {
		case int:
			return copies.(int), ok
		case float64:
			return int(copies.(float64)), ok
		case float32:
			return int(copies.(float32)), ok
		default:
			fmt.Printf("unexpected type: %T\n", copies)
			return 0, false
		}
	}

	return 0, false
}

func NewListAssetsOnMarketBlockchainRequestMetadata(assetIDs []int) BlockchainRequestMetadata {
	return BlockchainRequestMetadata{
		"assetIDs": assetIDs,
	}
}

func NewUpdateNftPriceBlockchainRequestMetadata(price float64) BlockchainRequestMetadata {
	return BlockchainRequestMetadata{
		"price": price,
	}
}

func (b BlockchainRequestMetadata) GetUpdateNftPricePrice() (float64, bool) {
	if price, ok := b["price"]; ok {
		switch price.(type) {
		case float64:
			return price.(float64), ok
		case float32:
			return float64(price.(float32)), ok
		default:
			fmt.Printf("unexpected type: %T\n", price)
			return 0, false
		}
	}

	return 0, false
}

func (b BlockchainRequestMetadata) GetListAssetsOnMarketAssetIDs() ([]int, bool) {
	metadataMap, err := json.Marshal(b)
	if err != nil {
		return nil, false
	}

	var assetIDs struct {
		AssetIDs []int `json:"assetIDs"`
	}

	err = json.Unmarshal(metadataMap, &assetIDs)
	if err != nil {
		return nil, false
	}

	return assetIDs.AssetIDs, true
}

type BlockchainRequestType string

var (
	BlockchainRequestMintNft                   BlockchainRequestType = "MintNFT"
	BlockchainRequestBurnNft                   BlockchainRequestType = "BurnNFT"
	BlockchainRequestBurnWaxNft                BlockchainRequestType = "BurnWaxNft"
	BlockchainRequestMintCollection            BlockchainRequestType = "MintCollection"
	BlockchainRequestMintWaxCollection         BlockchainRequestType = "MintWaxCollection"
	BlockchainRequestMintWaxSchema             BlockchainRequestType = "MintWaxSchema"
	BlockchainRequestUpdateWaxCollectionData   BlockchainRequestType = "UpdateWaxCollectionData"
	BlockchainRequestMintWaxNft                BlockchainRequestType = "MintWaxNft"
	BlockchainRequestUpdateWaxTemplateAssets   BlockchainRequestType = "UpdateWaxTemplateAssets"
	BlockchainRequestUpdateWaxCollectionWallet BlockchainRequestType = "UpdateWaxCollectionWallet"
	BlockchainRequestListWaxAssetOnMarket      BlockchainRequestType = "ListWaxAssetOnMarket"
	BlockchainRequestMintWaxAsset              BlockchainRequestType = "MintWaxAsset"
	BlockchainRequestUpdateWaxNftPrice         BlockchainRequestType = "UpdateWaxNftPrice"
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
