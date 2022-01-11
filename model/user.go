package model

type User struct {
	ID                  int    `json:"id"`      // used to parse from ent model
	UserId              int    `json:"user_id"` // used to parse from jwt token
	Name                string `json:"name,omitempty"`
	Email               string `json:"email"`
	ProfileImageUrl     string `json:"profile_image_url,omitempty"`
	HasAdminPanelAccess bool   `json:"has_admin_panel_access"`
	WalletAddress       string `json:"wallet_address"`
}
