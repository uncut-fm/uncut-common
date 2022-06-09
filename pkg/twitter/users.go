package twitter

import (
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
	"strings"
)

var (
	UsernameMissingErr = errors.New("username is missing")
	UserinfoErr        = errors.New("failed retrieving usersInfo")
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
		return nil, UsernameMissingErr
	}

	params := &twitter.UserLookupParams{ScreenName: []string{username}}
	users, resp, err := c.client.Lookup(params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, UserinfoErr
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
