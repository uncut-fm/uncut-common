package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"time"
)

type User struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	ProfileImageUrl       string    `json:"profileImageUrl"`
	IsNftCreator          bool      `json:"isNftCreator"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	Title                 string    `json:"title"`
	ThemeColorsAccent     string    `json:"themeColorsAccent"`
	ThemeColorsBackground string    `json:"themeColorsBackground"`
	IsAdmin               bool      `json:"isAdmin"`

	Wallets      []*Wallet      `json:"-"`
	NFTsCreated  []*NFT         `json:"-"`
	Transactions []*Transaction `json:"-"`
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

// GetPropertiesInMap returns a map of the user's properties; keys are in camelCase
func (u *User) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                    u.ID,
		"name":                  u.Name,
		"email":                 u.Email,
		"profileImageUrl":       u.ProfileImageUrl,
		"isNftCreator":          u.IsNftCreator,
		"createdAt":             u.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":             u.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"title":                 u.Title,
		"themeColorsAccent":     u.ThemeColorsAccent,
		"themeColorsBackground": u.ThemeColorsBackground,
		"isAdmin":               u.IsAdmin,
	}
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
