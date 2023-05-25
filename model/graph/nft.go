package model

import (
	"github.com/uncut-fm/uncut-common/pkg/proto/graph"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type NFT struct {
	ID                     int
	ContractAddress        string
	Price                  float64
	MintedOn               time.Time
	Status                 string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	UpdatedOnBlock         int
	Currency               string
	TokenID                string
	StoreID                int
	Fee                    float64
	CreatorAddress         string
	Supply                 int
	Balance                int
	Name                   string
	Description            string
	BlockchainDescription  string
	Perks                  string
	ImageURL               string
	BlockchainImageURL     string
	AnimationURL           string
	BlockchainAnimationURL string
	Type                   string
	Royalties              int
	ShowOnWebsite          bool
	Password               string
	DropAt                 string
	DropAtTime             time.Time
	DropType               string
	MintTransaction        string
	TemplateType           string
	FeaturedByCreator      bool
	NFTOwners              []*NFTOwner
	NFTCollection          *NFTCollection
	CreatedBy              *User
	Transactions           []*Transaction
}

func (n *NFT) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     n.ID,
		"contractAddress":        n.ContractAddress,
		"price":                  n.Price,
		"mintedOn":               n.MintedOn.Format("2006-01-02 15:04:05 MST"),
		"status":                 n.Status,
		"createdAt":              n.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":              n.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedOnBlock":         n.UpdatedOnBlock,
		"currency":               n.Currency,
		"tokenId":                n.TokenID,
		"storeId":                n.StoreID,
		"fee":                    n.Fee,
		"creatorAddress":         n.CreatorAddress,
		"supply":                 n.Supply,
		"balance":                n.Balance,
		"name":                   n.Name,
		"description":            n.Description,
		"blockchainDescription":  n.BlockchainDescription,
		"perks":                  n.Perks,
		"imageUrl":               n.ImageURL,
		"blockchainImageUrl":     n.BlockchainImageURL,
		"animationUrl":           n.AnimationURL,
		"blockchainAnimationUrl": n.BlockchainAnimationURL,
		"type":                   n.Type,
		"royalties":              n.Royalties,
		"showOnWebsite":          n.ShowOnWebsite,
		"password":               n.Password,
		"dropAt":                 n.DropAt,
		"dropAtTime":             n.DropAtTime.Format("2006-01-02 15:04:05 MST"),
		"mintTransaction":        n.MintTransaction,
		"templateType":           n.TemplateType,
	}
}

// convert NFT to graph.NFT
func (n *NFT) ToProto() *graph.Nft {
	return &graph.Nft{
		Id:                     int64(n.ID),
		ContractAddress:        n.ContractAddress,
		Price:                  n.Price,
		MintedOn:               timestamppb.New(n.MintedOn),
		Status:                 n.Status,
		CreatedAt:              timestamppb.New(n.CreatedAt),
		UpdatedAt:              timestamppb.New(n.UpdatedAt),
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
		DropAtTime:             timestamppb.New(n.DropAtTime),
		MintTransaction:        n.MintTransaction,
		TemplateType:           n.TemplateType,
	}
}
