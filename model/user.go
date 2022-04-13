package model

import "time"

type User struct {
	ID                  int      `json:"id"`      // used to parse from ent model
	UserId              int      `json:"user_id"` // used to parse from jwt token
	Name                string   `json:"name,omitempty"`
	Email               string   `json:"email"`
	ProfileImageUrl     string   `json:"profile_image_url,omitempty"`
	HasAdminPanelAccess bool     `json:"has_admin_panel_access"`
	WalletAddresses     []string `json:"wallet_addresses"`
	Faucet              Faucet   `json:"faucet"`
	TwitterHandle       string   `json:"twitter_handle"`
	IsNftCreator        bool     `json:"is_nft_creator"`
}

type Faucet struct {
	MaticAllowed bool      `json:"matic_allowed"`
	LastUsed     time.Time `json:"last_used"`
}
