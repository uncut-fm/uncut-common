package model

type NFTCollection struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	CreatorAddress  string `json:"creatorAddress"`
	CreatedAt       int64  `json:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt"`
	UpdatedOnBlock  int    `json:"updatedOnBlock"`
	Network         string `json:"network"`
	TokenType       string `json:"tokenType"`
	Origin          string `json:"origin"`
	NFTs            []*NFT `json:"-"`
}
