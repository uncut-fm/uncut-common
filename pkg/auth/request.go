package auth

type UpdateUserAuthRequest struct {
	ID                  int
	Name                *string `json:"name"`
	Email               *string `json:"email"`
	ProfileImageURL     *string `json:"profile_image_url"`
	HasAdminPanelAccess *bool   `json:"has_admin_panel_access"`
	WalletAddress       *string `json:"wallet_address"`
	TwitterHandle       *string `json:"twitter_handle"`
	IsNftCreator        *bool   `json:"is_nft_creator"`
}
