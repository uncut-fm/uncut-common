package model

import (
	"time"

	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

type User struct {
	ID              int          `json:"id"`      // used to parse from ent model
	UserId          int          `json:"user_id"` // used to parse from jwt token
	Name            string       `json:"name,omitempty"`
	Title           string       `json:"title,omitempty"`
	Email           string       `json:"email"`
	ProfileImageUrl string       `json:"profile_image_url,omitempty"`
	WalletAddresses []string     `json:"wallet_addresses"`
	TwitterHandle   string       `json:"twitter_handle"`
	IsNftCreator    bool         `json:"is_nft_creator"`
	ThemeColors     *ThemeColors `json:"theme_colors"`
	IsAdmin         bool         `json:"is_admin"`
	Edges           UserEdges    `json:"edges"`
}

type ThemeColors struct {
	Accent     string `json:"accent"`
	Background string `json:"background"`
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
	DeletedAt     time.Time `json:"deleted_at,omitempty"`
	UserID        int       `json:"userID"`
	Primary       bool      `json:"primary"`
	LastSyncedAt  time.Time `json:"last_synced_at"`
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

func ParseProtoUsersResponseToCommonUsers(protoResponse *proto_user.UsersResponse) []*User {
	protoUsers := protoResponse.GetUsers()

	return ParseProtoUsersToCommonUsers(protoUsers)
}

func ParseProtoUsersToCommonUsers(protoUsers []*proto_user.User) []*User {
	users := make([]*User, len(protoUsers))

	for i, protoUser := range protoUsers {
		users[i] = ParseProtoUserToUser(protoUser)
	}

	return users
}

func ParseProtoUserToUser(protoUser *proto_user.User) *User {
	user := &User{
		ID:              int(protoUser.Id),
		UserId:          int(protoUser.Id),
		Name:            protoUser.Name,
		Title:           protoUser.Title,
		Email:           protoUser.Email,
		ProfileImageUrl: protoUser.ProfileImageUrl,
		TwitterHandle:   protoUser.TwitterHandle,
		IsNftCreator:    protoUser.IsNftCreator,
		ThemeColors:     ParseProtoThemeColors(protoUser.ThemeColors),
		IsAdmin:         protoUser.IsAdmin,
		Edges:           UserEdges{Wallets: ParseProtoWalletsToWallets(protoUser.Edges.Wallets)},
	}

	user.SetWalletAddressesStringListFromEdges()
	return user
}

func ParseProtoWalletsToWallets(protoWallets []*proto_user.Wallet) []*Wallet {
	wallets := make([]*Wallet, len(protoWallets))

	for i, protoWallet := range protoWallets {
		wallets[i] = ParseProtoWalletToWallet(protoWallet)
	}

	return wallets
}

func ParseProtoWalletToWallet(protoWallet *proto_user.Wallet) *Wallet {
	primary := protoWallet.BecamePrimaryAt != nil
	return &Wallet{
		ID:            int(protoWallet.Id),
		Name:          protoWallet.Name,
		Description:   protoWallet.Description,
		CreatedAt:     protoWallet.CreatedAt.AsTime(),
		UpdatedAt:     protoWallet.UpdatedAt.AsTime(),
		WalletAddress: protoWallet.WalletAddress,
		Provider:      protoWallet.Provider,
		Primary:       primary,
		LastSyncedAt:  protoWallet.LastSyncedAt.AsTime(),
	}
}

func ParseProtoThemeColors(protoThemeColors *proto_user.ThemeColors) *ThemeColors {
	if protoThemeColors == nil {
		return nil
	}

	return &ThemeColors{
		Accent:     protoThemeColors.Accent,
		Background: protoThemeColors.Background,
	}
}

func ParseThemeColorsToProto(protoThemeColors *ThemeColors) *proto_user.ThemeColors {
	if protoThemeColors == nil {
		return nil
	}

	return &proto_user.ThemeColors{
		Accent:     protoThemeColors.Accent,
		Background: protoThemeColors.Background,
	}
}
