package model

import "time"

type User struct {
	ID                  int      `json:"id"` // used to parse from ent model
	Name                string   `json:"name,omitempty"`
	Email               string   `json:"email"`
	ProfileImageUrl     string   `json:"profile_image_url,omitempty"`
	HasAdminPanelAccess bool     `json:"has_admin_panel_access"`
	WalletAddresses     []string `json:"wallet_addresses"`
	Faucet              Faucet   `json:"faucet"`
	TwitterHandle       string   `json:"twitter_handle"`
}

type Faucet struct {
	MaticAllowed bool      `json:"matic_allowed"`
	LastUsed     time.Time `json:"last_used"`
}
