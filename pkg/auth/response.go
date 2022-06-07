package auth

import "github.com/uncut-fm/uncut-common/model"

type GetUserResponse struct {
	User *model.User `json:"user"`
}

type ListUsersResponse struct {
	Users []*model.User `json:"users"`
}

type GetOrCreateUserResponse struct {
	User          *model.User `json:"user"`
	ExistedBefore bool        `json:"existed_before"`
}
