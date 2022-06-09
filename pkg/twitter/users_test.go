package twitter

import (
	"bytes"
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
	"github.com/uncut-fm/uncut-common/pkg/testutil/mocks"
	"io"
	"net/http"
	"testing"
)

// ErrGeneric holds error that is returned in mocks functions
var ErrGeneric = errors.New("generic error")

func TestGetUserInfoErr(t *testing.T) {
	usersService := &mocks.TwitterUsersService{LookupFn: func(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error) {
		return nil, nil, ErrGeneric
	}}

	users := New(usersService)

	_, err := users.GetUserInfo("test")

	assert.Equal(t, ErrGeneric, err)
}

func TestGetUserInfoUsername(t *testing.T) {
	usersService := &mocks.TwitterUsersService{LookupFn: func(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error) {
		return nil, nil, ErrGeneric
	}}

	users := New(usersService)

	userInfo, err := users.GetUserInfo("")

	assert.Nil(t, userInfo)
	assert.Equal(t, UsernameMissingErr, err.Error())
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func TestGetUserInfoMock(t *testing.T) {
	username := "test_username"

	expectedName := "test Name"
	expectedBio := "test Bio"
	expectedURL := "test-url_normal.jpg"

	usersService := &mocks.TwitterUsersService{LookupFn: func(params *twitter.UserLookupParams) ([]twitter.User, *http.Response, error) {
		if params.ScreenName[0] == username {
			return []twitter.User{{Name: expectedName, Description: expectedBio, ProfileImageURL: expectedURL}}, &http.Response{StatusCode: http.StatusOK, Body: nopCloser{bytes.NewBuffer([]byte{})}}, nil
		}

		return nil, nil, ErrGeneric
	}}

	users := New(usersService)

	userInfo, err := users.GetUserInfo(username)

	assert.Nil(t, err)

	assert.Equal(t, userInfo.Name, expectedName)
	assert.Equal(t, userInfo.Bio, expectedBio)
	assert.Equal(t, userInfo.ProfileImageURL, "test-url.jpg")
}
