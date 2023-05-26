package model

import (
	"github.com/uncut-fm/uncut-common/pkg/proto/graph"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NFT struct {
	ID                     int     `json:"id"`
	ContractAddress        string  `json:"contractAddress"`
	Price                  float64 `json:"price"`
	MintedOn               int64   `json:"mintedOn"`
	Status                 string  `json:"status"`
	CreatedAt              int64   `json:"createdAt"`
	UpdatedAt              int64   `json:"updatedAt"`
	UpdatedOnBlock         int     `json:"updatedOnBlock"`
	Currency               string  `json:"currency"`
	TokenID                string  `json:"tokenId"`
	StoreID                int     `json:"storeId"`
	Fee                    float64 `json:"fee"`
	CreatorAddress         string  `json:"creatorAddress"`
	Supply                 int     `json:"supply"`
	Balance                int     `json:"balance"`
	Name                   string  `json:"name"`
	Description            string  `json:"description"`
	BlockchainDescription  string  `json:"blockchainDescription"`
	Perks                  string  `json:"perks"`
	ImageURL               string  `json:"imageUrl"`
	BlockchainImageURL     string  `json:"blockchainImageUrl"`
	AnimationURL           string  `json:"animationUrl"`
	BlockchainAnimationURL string  `json:"blockchainAnimationUrl"`
	Type                   string  `json:"type"`
	Royalties              int     `json:"royalties"`
	ShowOnWebsite          bool    `json:"showOnWebsite"`
	Password               string  `json:"password"`
	DropAt                 string  `json:"dropAt"`
	DropAtTime             int64   `json:"dropAtTime"`
	DropType               string  `json:"dropType"`
	MintTransaction        string  `json:"mintTransaction"`
	TemplateType           string  `json:"templateType"`
	FeaturedByCreator      bool    `json:"featuredByCreator"`
	NFTOwners              []*NFTOwner
	NFTCollection          *NFTCollection
	CreatedBy              *User
	Transactions           []*Transaction
}

// convert NFT to graph.NFT
func (n *NFT) ToProto() *graph.Nft {
	return &graph.Nft{
		Id:                     int64(n.ID),
		ContractAddress:        n.ContractAddress,
		Price:                  n.Price,
		MintedOn:               timestamppb.New(UnixTimeToTime(n.MintedOn)),
		Status:                 n.Status,
		CreatedAt:              timestamppb.New(UnixTimeToTime(n.CreatedAt)),
		UpdatedAt:              timestamppb.New(UnixTimeToTime(n.UpdatedAt)),
		UpdatedOnBlock:         int32(n.UpdatedOnBlock),
		Currency:               n.Currency,
		TokenId:                n.TokenID,
		StoreId:                int32(n.StoreID),
		Fee:                    n.Fee,
		CreatorAddress:         n.CreatorAddress,
		Supply:                 int32(n.Supply),
		Balance:                int32(n.Balance),
		Name:                   n.Name,
		Description:            n.Description,
		BlockchainDescription:  n.BlockchainDescription,
		Perks:                  n.Perks,
		ImageUrl:               n.ImageURL,
		BlockchainImageUrl:     n.BlockchainImageURL,
		AnimationUrl:           n.AnimationURL,
		BlockchainAnimationUrl: n.BlockchainAnimationURL,
		Type:                   n.Type,
		Royalties:              int32(n.Royalties),
		ShowOnWebsite:          n.ShowOnWebsite,
		Password:               n.Password,
		DropAt:                 n.DropAt,
		DropAtTime:             timestamppb.New(UnixTimeToTime(n.DropAtTime)),
		MintTransaction:        n.MintTransaction,
		TemplateType:           n.TemplateType,
	}
}
