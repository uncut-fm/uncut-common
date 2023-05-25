package model

import (
	"time"
)

type NFTCollection struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	ContractAddress string    `json:"contractAddress"`
	CreatorAddress  string    `json:"creatorAddress"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedOnBlock  int       `json:"updatedOnBlock"`
	Network         string    `json:"network"`
	TokenType       string    `json:"tokenType"`
	Origin          string    `json:"origin"`
	NFTs            []*NFT    `json:"-"`
}

// GetPropertiesInMap returns a map of the properties of the NFTCollection; keys are in camelCase
func (n *NFTCollection) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              n.ID,
		"name":            n.Name,
		"contractAddress": n.ContractAddress,
		"creatorAddress":  n.CreatorAddress,
		"createdAt":       n.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":       n.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedOnBlock":  n.UpdatedOnBlock,
		"network":         n.Network,
		"tokenType":       n.TokenType,
		"origin":          n.Origin,
	}
}
