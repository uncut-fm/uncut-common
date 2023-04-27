package auth

import "github.com/uncut-fm/uncut-common/model"

type GetOrCreateUserResponse struct {
	User          *model.User `json:"user"`
	ExistedBefore bool        `json:"existed_before"`
}

type UsersInfoResponse struct {
	TotalCount int
	Users      []*model.User
}
