package model

import (
	"github.com/uncut-fm/uncut-common/pkg/config"
	"time"

	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

type User struct {
	ID                   int          `json:"id"`
	Name                 string       `json:"name,omitempty"`
	Title                string       `json:"title,omitempty"`
	Email                string       `json:"email"`
	ProfileImageUrl      string       `json:"profile_image_url,omitempty"`
	WalletAddresses      []string     `json:"wallet_addresses"`
	TwitterHandle        string       `json:"twitter_handle"`
	IsNftCreator         bool         `json:"is_nft_creator"`
	ThemeColors          *ThemeColors `json:"theme_colors"`
	IsAdmin              bool         `json:"is_admin"`
	BannerImageUrl       string       `json:"banner_image_url"`
	Location             string       `json:"location"`
	VerificationStatus   string       `json:"verification_status"`
	Bio                  string       `json:"bio"`
	InstagramHandle      string       `json:"instagram_handle"`
	FacebookHandle       string       `json:"facebook_handle"`
	LinkedinHandle       string       `json:"linkedin_handle"`
	DiscordHandle        string       `json:"discord_handle"`
	WebsiteUrl           string       `json:"website_url"`
	Type                 string       `json:"type"`
	Karma                int32        `json:"karma"`
	KarmaIn30Days        int32        `json:"karma_in_30_days"`
	LastKarmaProcessedAt time.Time    `json:"last_karma_processed_at"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
	LastLoggedInAt       time.Time    `json:"last_logged_in_at"`
	Edges                UserEdges    `json:"edges"`
}

func (u User) IsProfileComplete() bool {
	return u.Name != "" && u.ProfileImageUrl != "" && u.Title != "" && len(u.WalletAddresses) > 0 && u.countUserSocialLinks() >= 1
}

func (u User) countUserSocialLinks() int {
	var count int

	if len(u.FacebookHandle) > 0 {
		count++
	}

	if len(u.InstagramHandle) > 0 {
		count++
	}

	if len(u.TwitterHandle) > 0 {
		count++
	}

	if len(u.LinkedinHandle) > 0 {
		count++
	}

	if len(u.DiscordHandle) > 0 {
		count++
	}

	if len(u.WebsiteUrl) > 0 {
		count++
	}

	return count
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
		ID:                   int(protoUser.Id),
		Name:                 protoUser.Name,
		Title:                protoUser.Title,
		Email:                protoUser.Email,
		ProfileImageUrl:      protoUser.ProfileImageUrl,
		TwitterHandle:        protoUser.TwitterHandle,
		IsNftCreator:         protoUser.IsNftCreator,
		ThemeColors:          ParseProtoThemeColors(protoUser.ThemeColors),
		IsAdmin:              protoUser.IsAdmin,
		BannerImageUrl:       protoUser.BannerImageUrl,
		Location:             protoUser.Location,
		VerificationStatus:   protoUser.VerificationStatus,
		Bio:                  protoUser.Bio,
		InstagramHandle:      protoUser.InstagramHandle,
		FacebookHandle:       protoUser.FacebookHandle,
		LinkedinHandle:       protoUser.LinkedinHandle,
		DiscordHandle:        protoUser.DiscordHandle,
		WebsiteUrl:           protoUser.WebsiteUrl,
		Type:                 protoUser.Type,
		Karma:                protoUser.Karma,
		KarmaIn30Days:        protoUser.KarmaIn_30Days,
		LastKarmaProcessedAt: protoUser.LastKarmaProcessedAt.AsTime(),
		CreatedAt:            protoUser.CreatedAt.AsTime(),
		UpdatedAt:            protoUser.UpdatedAt.AsTime(),
		LastLoggedInAt:       protoUser.LastLoggedInAt.AsTime(),
		Edges:                UserEdges{Wallets: ParseProtoWalletsToWallets(protoUser.Edges.Wallets)},
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

type UserSession struct {
	User         *User
	AccessToken  string
	RefreshToken string
	IsNewUser    bool
}

func ParseProtoUserSessionToUserSession(protoUserSession *proto_user.UserSessionResponse) *UserSession {
	return &UserSession{
		User:         ParseProtoUserToUser(protoUserSession.User),
		AccessToken:  protoUserSession.AccessToken,
		RefreshToken: protoUserSession.RefreshToken,
		IsNewUser:    protoUserSession.IsNewUser,
	}
}

type UserOrder struct {
	Field UserOrderField
	Desc  bool
}

type UserOrderField string

func (u UserOrderField) String() string {
	return string(u)
}

var (
	UserOrderFieldCreatedAt     UserOrderField = "created_at"
	UserOrderFieldKarma         UserOrderField = "karma"
	UserOrderFieldKarmaIn30Days UserOrderField = "karma_in_30_days"
)

func ParseUserOrderToProto(order *UserOrder) *proto_user.UserOrder {
	if order == nil {
		return nil
	}

	return &proto_user.UserOrder{
		Field: order.Field.String(),
		Desc:  order.Desc,
	}
}

func ParseProtoUserOrder(protoOrder *proto_user.UserOrder) *UserOrder {
	if protoOrder == nil {
		return nil
	}

	return &UserOrder{
		Field: UserOrderField(protoOrder.Field),
		Desc:  protoOrder.Desc,
	}
}

const (
	MainProdUserID  = 4294967824
	MainTestUserID  = 8589934595
	MainLocalUserID = 8589934603
)

func GetMainUserIDByEnvironment(env string) int {
	switch env {
	case config.ProdEnvironment, config.StageEnvironment:
		return MainProdUserID
	case config.TestEnvironment:
		return MainTestUserID
	case config.LocalEnvironment, config.DevEnvironment:
		return MainLocalUserID
	default:
		return MainProdUserID
	}
}
