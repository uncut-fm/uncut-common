package model

import (
	"time"
)

type NFTCollection struct {
	ID              int
	Name            string
	ContractAddress string
	CreatorAddress  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UpdatedOnBlock  int
	Network         string
	TokenType       string
	Origin          string
	NFTs            []*NFT
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
