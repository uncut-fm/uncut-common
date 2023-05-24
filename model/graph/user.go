package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	"time"
)

type User struct {
	ID                    int
	Name                  string
	Email                 string
	ProfileImageUrl       string
	IsNftCreator          bool
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Title                 string
	ThemeColorsAccent     string
	ThemeColorsBackground string
	IsAdmin               bool

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

// SetUpdatedFields sets the fields that differ between the two users
func (u *User) SetUpdatedFields(srcUser *User) bool {
	return setUpdatedFields(u, srcUser)
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
