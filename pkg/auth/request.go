package auth

import "github.com/uncut-fm/uncut-common/model"

type UpdateUserAuthRequest struct {
	ID              int
	Name            *string            `json:"name"`
	Title           *string            `json:"title"`
	Email           *string            `json:"email"`
	ProfileImageURL *string            `json:"profile_image_url"`
	WalletAddress   *string            `json:"wallet_address"`
	TwitterHandle   *string            `json:"twitter_handle"`
	IsNftCreator    *bool              `json:"is_nft_creator"`
	ThemeColors     *model.ThemeColors `json:"theme_colors"`
}

type UpdateWalletRequest struct {
	UserID      int
	WalletID    int     `json:"wallet_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Primary     *bool   `json:"primary"`
}

type AddWalletRequest struct {
	UserID        int
	WalletAddress string
	Name          *string
	Description   *string
	Provider      string
}

type DeleteWalletRequest struct {
	UserID   int
	WalletID int
}
