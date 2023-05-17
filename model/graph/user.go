package model

import (
	"github.com/mindstand/gogm/v2"
	common_model "github.com/uncut-fm/uncut-common/model"
)

type User struct {
	gogm.BaseNode

	ID                    int    `gogm:"name=id,pk"`
	Name                  string `gogm:"name=name"`
	Email                 string `gogm:"name=email"`
	ProfileImageUrl       string `gogm:"name=profile_image_url"`
	IsNftCreator          bool   `gogm:"name=is_nft_creator"`
	CreatedAt             int    `gogm:"name=created_at"`
	UpdatedAt             int    `gogm:"name=updated_at"`
	Title                 string `gogm:"name=title"`
	ThemeColorsAccent     string `gogm:"name=theme_colors_accent"`
	ThemeColorsBackground string `gogm:"name=theme_colors_background"`
	IsAdmin               bool   `gogm:"name=is_admin"`

	Wallets      []*Wallet      `gogm:"direction=outgoing;relationship=OWNS"`
	NFTsCreated  []*NFT         `gogm:"direction=outgoing;relationship=CREATED"`
	Transactions []*Transaction `gogm:"direction=outgoing;relationship=PERFORMED"`
}

func NewUsersListFromCommonUsers(commonUsers []*common_model.User) []*User {
	users := make([]*User, len(commonUsers))

	for i := range commonUsers {
		users[i] = NewUserFromCommonUser(commonUsers[i])
	}

	return users
}

func NewUserFromCommonUser(commonUser *common_model.User) *User {
	user := &User{
		ID:              commonUser.UserId,
		Name:            commonUser.Name,
		Email:           commonUser.Email,
		ProfileImageUrl: commonUser.ProfileImageUrl,
		IsNftCreator:    commonUser.IsNftCreator,
		Title:           commonUser.Title,
		IsAdmin:         commonUser.IsAdmin,
	}

	user.Wallets = NewWalletsListFromCommonWallets(commonUser.Edges.Wallets, user)

	if commonUser.ThemeColors != nil {
		user.ThemeColorsAccent = commonUser.ThemeColors.Accent
		user.ThemeColorsBackground = commonUser.ThemeColors.Background
	}

	return user
}
