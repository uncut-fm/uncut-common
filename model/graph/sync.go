package model

type GraphModelType string

var (
	NFTCollectionType GraphModelType = "NFTCollection"
	NFTOwnerType      GraphModelType = "NFTOwner"
	NFTType           GraphModelType = "NFT"
	TransactionType   GraphModelType = "Transaction"
	WalletType        GraphModelType = "Wallet"
	UserType          GraphModelType = "User"
)

type SyncEvent struct {
	ModelType GraphModelType
	Model     interface{}
}
