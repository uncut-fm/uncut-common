package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-backoffice-api/graph/model"
	model_common "github.com/uncut-fm/uncut-common/model"
	mngmt_model "github.com/uncut-fm/uncut-management-api-2/graph/model"
)

func (a API) UpdateUser(input *model.UpdateUserAuthRequest) (*mngmt_model.User, error) {
	commonUser, err := a.updateUser(input)
	if a.log.CheckError(err, a.UpdateUser) != nil {
		return nil, err
	}

	user := mngmt_model.NewUserFromCommonModel(commonUser)

	return user, nil
}

func (a API) updateUser(input *model.UpdateUserAuthRequest) (*model_common.User, error) {
	response, err := a.makeUpdateUserRequest(input)
	if a.log.CheckError(err, a.updateUser) != nil {
		return nil, err
	}

	user := new(model_common.User)

	err = json.Unmarshal(response.Body(), user)

	return user, a.log.CheckError(err, a.updateUser)
}

func (a API) makeUpdateUserRequest(input *model.UpdateUserAuthRequest) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetBody(input).
		SetHeader("admin-token", a.authAdminToken).
		Patch(fmt.Sprintf(updateUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetOrCreateUserRequest)
}
