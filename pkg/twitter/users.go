package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/uncut-fm/uncut-common/pkg/errors"
	"net/http"
	"strings"
)

type Users struct {
	client UsersService
}

func New(client UsersService) *Users {
	return &Users{client: client}
}

type UsersService interface {
	Lookup(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error)
}

type UserInfo struct {
	Name            string `json:"name"`
	Bio             string `json:"bio"`
	ProfileImageURL string `json:"profile_image"`
}

func (c *Users) GetUserInfo(username string) (*UserInfo, error) {
	if len(username) == 0 {
		return nil, errors.UsernameMissingErr
	}

	params := &twitter.UserLookupParams{ScreenName: []string{username}}
	users, resp, err := c.client.Lookup(params)
	if err != nil {
		return nil, errors.UserHandleNotFoundErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.UserinfoErr
	}

	originalProfileImage := getTwitterOriginalProfileImage(users[0].ProfileImageURL)

	userInfo := &UserInfo{
		Name:            users[0].Name,
		Bio:             users[0].Description,
		ProfileImageURL: originalProfileImage,
	}

	return userInfo, nil
}

func getTwitterOriginalProfileImage(profileImageURL string) string {
	return strings.ReplaceAll(profileImageURL, "_normal", "")
}
