package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) UpdateUser(input *UpdateUserAuthRequest) (*model.User, error) {
	commonUser, err := a.updateUser(input)
	if a.log.CheckError(err, a.UpdateUser) != nil {
		return nil, err
	}

	return commonUser, nil
}

func (a API) updateUser(input *UpdateUserAuthRequest) (*model.User, error) {
	response, err := a.makeUpdateUserRequest(input)
	if a.log.CheckError(err, a.updateUser) != nil {
		return nil, err
	}

	user := new(model.User)

	err = json.Unmarshal(response.Body(), user)

	return user, a.log.CheckError(err, a.updateUser)
}

func (a API) makeUpdateUserRequest(input *UpdateUserAuthRequest) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetBody(input).
		SetHeader("admin-token", a.authAdminToken).
		Patch(fmt.Sprintf(updateUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetOrCreateUserRequest)
}
