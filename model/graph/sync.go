package model

import "encoding/json"

type GraphModelType string

var (
	GraphModelTypeNFTCollection GraphModelType = "NFTCollection"
	GraphModelTypeNFTOwner      GraphModelType = "NFTOwner"
	GraphModelTypeNFT           GraphModelType = "NFT"
	GraphModelTypeTransaction   GraphModelType = "Transaction"
	GraphModelTypeWallet        GraphModelType = "Wallet"
	GraphModelTypeUser          GraphModelType = "User"
)

type SyncEvent struct {
	ModelType GraphModelType
	NFTs      []*NFT
	Users     []*User
}

func (s SyncEvent) Marshal() ([]byte, error) {
	return json.Marshal(s)
}
