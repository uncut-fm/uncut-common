package auth

type UpdateUserAuthRequest struct {
	ID              int
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	ProfileImageURL *string `json:"profile_image_url"`
	WalletAddress   *string `json:"wallet_address"`
	TwitterHandle   *string `json:"twitter_handle"`
	IsNftCreator    *bool   `json:"is_nft_creator"`
}

type UpdateWalletRequest struct {
	UserID      int
	WalletID    int     `json:"wallet_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Primary     *bool   `json:"primary"`
}
