package model

import (
	"time"
)

type User struct {
	ID              int       `json:"id"`      // used to parse from ent model
	UserId          int       `json:"user_id"` // used to parse from jwt token
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email"`
	ProfileImageUrl string    `json:"profile_image_url,omitempty"`
	WalletAddresses []string  `json:"wallet_addresses"`
	Faucet          Faucet    `json:"faucet"`
	TwitterHandle   string    `json:"twitter_handle"`
	IsNftCreator    bool      `json:"is_nft_creator"`
	Edges           UserEdges `json:"edges"`
}

type UserEdges struct {
	Wallets []*Wallet `json:"wallets"`
}

type Wallet struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Description   string    `json:"description,omitempty"`
	WalletAddress string    `json:"wallet_address,omitempty"`
	Provider      string    `json:"provider,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Faucet struct {
	MaticAllowed bool      `json:"matic_allowed"`
	LastUsed     time.Time `json:"last_used"`
}

func (u *User) SetWalletAddressesStringListFromEdges() {

	u.WalletAddresses = make([]string, len(u.Edges.Wallets))

	for i := range u.Edges.Wallets {
		u.WalletAddresses[i] = u.Edges.Wallets[i].WalletAddress
	}
}
