package model

import (
	"github.com/mindstand/gogm/v2"
	"time"
)

type NFT struct {
	gogm.BaseNode

	ID                int            `gogm:"name=id,pk"`
	ContractAddress   string         `gogm:"name=contract_address"`
	Price             float64        `gogm:"name=price"`
	MintedOn          time.Time      `gogm:"name=minted_on"`
	Status            string         `gogm:"name=status"`
	CreatedAt         time.Time      `gogm:"name=created_at"`
	UpdatedAt         time.Time      `gogm:"name=updated_at"`
	UpdatedOnBlock    int            `gogm:"name=updated_on_block"`
	Currency          string         `gogm:"name=currency"`
	TokenID           string         `gogm:"name=token_id"`
	StoreID           int            `gogm:"name=store_id"`
	Fee               float64        `gogm:"name=fee"`
	CreatorAddress    string         `gogm:"name=creator_address"`
	Supply            int            `gogm:"name=supply"`
	Balance           int            `gogm:"name=balance"`
	Name              string         `gogm:"name=name"`
	Description       string         `gogm:"name=description"`
	BlockchainDesc    string         `gogm:"name=blockchain_description"`
	Perks             string         `gogm:"name=perks"`
	ImageURL          string         `gogm:"name=image_url"`
	BlockchainImgURL  string         `gogm:"name=blockchain_image_url"`
	AnimationURL      string         `gogm:"name=animation_url"`
	BlockchainAnimURL string         `gogm:"name=blockchain_animation_url"`
	Type              string         `gogm:"name=type"`
	WebsiteDesc       string         `gogm:"name=website_description"`
	TagLine           string         `gogm:"name=tag_line"`
	DropOfWeek        bool           `gogm:"name=drop_of_the_week"`
	Royalties         int            `gogm:"name=royalties"`
	ShowOnWebsite     bool           `gogm:"name=show_on_website"`
	Password          string         `gogm:"name=password"`
	DropAt            string         `gogm:"name=drop_at"`
	MintTransaction   string         `gogm:"name=mint_transaction"`
	TemplateType      string         `gogm:"name=template_type"`
	NFTOwner          *NFTOwner      `gogm:"direction=incoming;relationship=OWNS"`
	Collection        *NFTCollection `gogm:"direction=incoming;relationship=BELONGS_TO"`
	CreatedBy         *User          `gogm:"direction=incoming;relationship=CREATED"`
	Transactions      []*Transaction `gogm:"direction=outgoing;relationship=INVOLVES"`
}
