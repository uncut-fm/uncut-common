package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

type User struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	ProfileImageUrl       string `json:"profileImageUrl"`
	IsNftCreator          bool   `json:"isNftCreator"`
	CreatedAt             int64  `json:"createdAt"`
	UpdatedAt             int64  `json:"updatedAt"`
	Title                 string `json:"title"`
	ThemeColorsAccent     string `json:"themeColorsAccent"`
	ThemeColorsBackground string `json:"themeColorsBackground"`
	IsAdmin               bool   `json:"isAdmin"`

	Wallets      []*Wallet
	NFTsCreated  []*NFT
	Transactions []*Transaction
}

// NewUsersListFromCommonUsers converts a slice of common_model.User to a slice of *User
func NewUsersListFromCommonUsers(commonUsers []*common_model.User) []*User {
	users := make([]*User, len(commonUsers))

	for i := range commonUsers {
		users[i] = NewUserFromCommonUser(commonUsers[i])
	}

	return users
}

func NewUserFromCommonUser(commonUser *common_model.User) *User {
	user := &User{
		ID:              commonUser.ID,
		Name:            commonUser.Name,
		Email:           commonUser.Email,
		ProfileImageUrl: commonUser.ProfileImageUrl,
		IsNftCreator:    commonUser.IsNftCreator,
		Title:           commonUser.Title,
		IsAdmin:         commonUser.IsAdmin,
	}

	user.Wallets = NewWalletsListFromCommonWallets(commonUser.Edges.Wallets, *user)

	if commonUser.ThemeColors != nil {
		user.ThemeColorsAccent = commonUser.ThemeColors.Accent
		user.ThemeColorsBackground = commonUser.ThemeColors.Background
	}

	return user
}

// convert User to proto proto.User
func (u *User) ToProto() *proto_user.User {
	return &proto_user.User{
		Id:              uint64(u.ID),
		Name:            u.Name,
		Email:           u.Email,
		ProfileImageUrl: u.ProfileImageUrl,
		IsNftCreator:    u.IsNftCreator,
		Title:           u.Title,
		ThemeColors: &proto_user.ThemeColors{
			Accent:     u.ThemeColorsAccent,
			Background: u.ThemeColorsBackground,
		},
		IsAdmin: u.IsAdmin,
		Edges:   &proto_user.UserEdges{Wallets: NewProtoWalletsListFromWallets(u.Wallets)},
	}
}
