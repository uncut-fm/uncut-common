package model

import (
	"github.com/uncut-fm/uncut-common/pkg/proto/graph"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type NFT struct {
	ID                     int            `json:"id"`
	ContractAddress        string         `json:"contractAddress"`
	Price                  float64        `json:"price"`
	MintedOn               time.Time      `json:"mintedOn"`
	Status                 string         `json:"status"`
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
	UpdatedOnBlock         int            `json:"updatedOnBlock"`
	Currency               string         `json:"currency"`
	TokenID                string         `json:"tokenId"`
	StoreID                int            `json:"storeId"`
	Fee                    float64        `json:"fee"`
	CreatorAddress         string         `json:"creatorAddress"`
	Supply                 int            `json:"supply"`
	Balance                int            `json:"balance"`
	Name                   string         `json:"name"`
	Description            string         `json:"description"`
	BlockchainDescription  string         `json:"blockchainDescription"`
	Perks                  string         `json:"perks"`
	ImageURL               string         `json:"imageUrl"`
	BlockchainImageURL     string         `json:"blockchainImageUrl"`
	AnimationURL           string         `json:"animationUrl"`
	BlockchainAnimationURL string         `json:"blockchainAnimationUrl"`
	Type                   string         `json:"type"`
	Royalties              int            `json:"royalties"`
	ShowOnWebsite          bool           `json:"showOnWebsite"`
	Password               string         `json:"password"`
	DropAt                 string         `json:"dropAt"`
	DropAtTime             time.Time      `json:"dropAtTime"`
	DropType               string         `json:"dropType"`
	MintTransaction        string         `json:"mintTransaction"`
	TemplateType           string         `json:"templateType"`
	FeaturedByCreator      bool           `json:"featuredByCreator"`
	NFTOwners              []*NFTOwner    `json:"-"`
	NFTCollection          *NFTCollection `json:"-"`
	CreatedBy              *User          `json:"-"`
	Transactions           []*Transaction `json:"-"`
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
