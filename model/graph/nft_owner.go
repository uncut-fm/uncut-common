package model

type NFTOwner struct {
	ID                  int      `json:"id"`
	UserWalletAddress   string   `json:"userWalletAddress"`
	Balance             int      `json:"balance"`
	CreatedAt           int64    `json:"createdAt"`
	UpdatedAt           int64    `json:"updatedAt"`
	IsHidden            bool     `json:"isHidden"`
	TransactionsStrings []string `json:"transactions"`
	AssetIDs            []string `json:"assetIDs"`

	Wallet       *Wallet
	NFTs         []*NFT
	Transactions []*Transaction
}
