package model

import "github.com/mindstand/gogm/v2"

type NFTOwner struct {
	gogm.BaseNode

	ID                int    `gogm:"name=id,pk"`
	UserWalletAddress string `gogm:"name=user_wallet_address"`
	Balance           int    `gogm:"name=balance"`
	CreatedAt         int    `gogm:"name=created_at"`
	UpdatedAt         int    `gogm:"name=updated_at"`
	IsHidden          bool   `gogm:"name=is_hidden"`

	Wallet       *Wallet        `gogm:"direction=incoming;relationship=CONTAINS"`
	NFTs         []*NFT         `gogm:"direction=outgoing;relationship=OWNS"`
	Transactions []*Transaction `gogm:"direction=outgoing;relationship=HAS"`
}
