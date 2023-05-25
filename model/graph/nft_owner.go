package model

import "time"

type NFTOwner struct {
	ID                  int       `json:"id"`
	UserWalletAddress   string    `json:"userWalletAddress"`
	Balance             int       `json:"balance"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	IsHidden            bool      `json:"isHidden"`
	TransactionsStrings []string  `json:"transactions"`

	Wallet       *Wallet        `json:"-"`
	NFTs         []*NFT         `json:"-"`
	Transactions []*Transaction `json:"-"`
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
