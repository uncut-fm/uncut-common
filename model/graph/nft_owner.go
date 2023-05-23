package model

import "time"

type NFTOwner struct {
	ID                  int
	UserWalletAddress   string
	Balance             int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	IsHidden            bool
	TransactionsStrings []string

	Wallet       *Wallet
	NFTs         []*NFT
	Transactions []*Transaction
}

// GetPropertiesInMap returns a map of the properties of the NFTOwner; keys are in camelCase
func (n *NFTOwner) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                n.ID,
		"userWalletAddress": n.UserWalletAddress,
		"balance":           n.Balance,
		"createdAt":         n.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":         n.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"isHidden":          n.IsHidden,
	}
}
