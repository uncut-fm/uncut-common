package mocks

import (
	"github.com/dghubble/go-twitter/twitter"
	common_twitter "github.com/uncut-fm/uncut-common/pkg/twitter"
	"net/http"
)

type TwitterUsers struct {
	GetUserInfoFn func(username string) (*common_twitter.UserInfo, error)
}

func (t *TwitterUsers) GetUserInfo(username string) (*common_twitter.UserInfo, error) {
	return t.GetUserInfoFn(username)
}

type TwitterUsersService struct {
	LookupFn func(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error)
}

func (t *TwitterUsersService) Lookup(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error) {
	return t.LookupFn(params)
}
