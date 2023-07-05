package auth

import (
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

type UpdateUserAuthRequest struct {
	ID                 int
	Name               *string            `json:"name"`
	Title              *string            `json:"title"`
	Email              *string            `json:"email"`
	ProfileImageURL    *string            `json:"profile_image_url"`
	WalletAddress      *string            `json:"wallet_address"`
	TwitterHandle      *string            `json:"twitter_handle"`
	IsNftCreator       *bool              `json:"is_nft_creator"`
	ThemeColors        *model.ThemeColors `json:"theme_colors"`
	Bio                *string            `json:"bio"`
	BannerImageURL     *string            `json:"banner_image_url"`
	InstagramHandle    *string            `json:"instagram_handle"`
	FacebookHandle     *string            `json:"facebook_handle"`
	LinkedinHandle     *string            `json:"linkedin_handle"`
	DiscordHandle      *string            `json:"discord_handle"`
	WebsiteURL         *string            `json:"website_url"`
	Location           *string            `json:"location"`
	VerificationStatus *string            `json:"verification_status"`
}

type UpdateWalletRequest struct {
	UserID       int
	WalletID     int        `json:"wallet_id"`
	Name         *string    `json:"name"`
	Description  *string    `json:"description"`
	Primary      *bool      `json:"primary"`
	LastSyncedAt *time.Time `json:"last_synced_at"`
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
