package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"net/url"
	"strconv"
)

func (a API) GetOrCreateUser(email string) (*GetOrCreateUserResponse, error) {
	response, err := a.makeGetOrCreateUserRequest(email)
	if a.log.CheckError(err, a.GetOrCreateUser) != nil {
		return nil, err
	}

	getOrCreateUserResponse, err := a.getGetOrCreateUserResponse(response)
	return getOrCreateUserResponse, a.log.CheckError(err, a.GetOrCreateUser)
}

func (a API) makeGetOrCreateUserRequest(email string) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetBody(map[string]string{
			"email": email,
		}).
		SetHeader("admin-token", a.authAdminToken).
		Post(fmt.Sprintf(getOrCreateUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetOrCreateUserRequest)

}

func (a API) getGetOrCreateUserResponse(resp *resty.Response) (*GetOrCreateUserResponse, error) {
	responseStruct := new(GetOrCreateUserResponse)
	err := json.Unmarshal(resp.Body(), responseStruct)

	return responseStruct, a.log.CheckError(err, a.getGetOrCreateUserResponse)
}

func (a API) GetNftCreators() ([]*model.User, error) {
	commonUsers, err := a.getNftCreators()

	return commonUsers, a.log.CheckError(err, a.GetNftCreators)
}

func (a API) getNftCreators() ([]*model.User, error) {
	response, err := a.makeGetCreatorsRequest()
	if a.log.CheckError(err, a.getNftCreators) != nil {
		return nil, err
	}

	var users []*model.User

	err = json.Unmarshal(response.Body(), &users)

	return users, a.log.CheckError(err, a.getNftCreators)
}

func (a API) makeGetCreatorsRequest() (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetHeader("admin-token", a.authAdminToken).
		Get(fmt.Sprintf(getCreatorsEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetCreatorsRequest)
}

func (a API) GetUserByEmail(email string) (*model.User, error) {
	response, err := a.makeGetUserByEmailRequest(email)
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	getUserResponse, err := a.getGetUserResponse(response)
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	if getUserResponse.User == nil {
		return nil, errors.New("user not found")
	}

	return getUserResponse.User, a.log.CheckError(err, a.GetUserByEmail)
}

func (a API) makeGetUserByEmailRequest(email string) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetQueryParam("email", email).
		SetHeader("admin-token", a.authAdminToken).
		Get(fmt.Sprintf(getUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetUserByEmailRequest)
}

func (a API) GetUserByWalletAddress(walletAddress string) (*model.User, error) {
	commonUser, err := a.getUserByWalletAddress(walletAddress)
	if a.log.CheckError(err, a.GetUserByWalletAddress) != nil {
		return nil, err
	}

	return commonUser, a.log.CheckError(err, a.GetUserByWalletAddress)
}

func (a API) getUserByWalletAddress(walletAddress string) (*model.User, error) {
	response, err := a.makeGetUserByWalletRequest(walletAddress)
	if a.log.CheckError(err, a.getUserByWalletAddress) != nil {
		return nil, err
	}

	getUserResponse, err := a.getGetUserResponse(response)
	if a.log.CheckError(err, a.getUserByWalletAddress) != nil {
		return nil, err
	}

	if getUserResponse.User == nil {
		return nil, errors.New("user not found")
	}

	return getUserResponse.User, nil
}

func (a API) makeGetUserByWalletRequest(walletAddress string) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetQueryParam("wallet", walletAddress).
		SetHeader("admin-token", a.authAdminToken).
		Get(fmt.Sprintf(getUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetUserByWalletRequest)
}

func (a API) GetUserByID(userID int) (*model.User, error) {
	commonUser, err := a.getUserByID(userID)
	if a.log.CheckError(err, a.GetUserByID) != nil {
		return nil, err
	}

	return commonUser, a.log.CheckError(err, a.GetUserByID)
}

func (a API) getUserByID(userID int) (*model.User, error) {
	response, err := a.makeGetUserByIDRequest(userID)
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	getUserResponse, err := a.getGetUserResponse(response)
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	if getUserResponse.User == nil {
		return nil, errors.New("user not found")
	}

	go a.cache.SetUserEmailToCache(getUserResponse.User)

	return getUserResponse.User, nil
}

func (a API) makeGetUserByIDRequest(userID int) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetQueryParam("id", strconv.Itoa(userID)).
		SetHeader("admin-token", a.authAdminToken).
		Get(fmt.Sprintf(getUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeGetUserByIDRequest)
}

func (a API) GetUserEmailByWalletAddress(walletAddress string) (string, error) {
	email, exist := a.cache.GetEmailByWalletCache(walletAddress)
	if exist {
		return email, nil
	}

	user, err := a.getUserByWalletAddress(walletAddress)
	if a.log.CheckError(err, a.GetUserEmailByWalletAddress) != nil {
		return "", err
	}

	return user.Email, nil
}

func (a API) ListUsersByWalletAddresses(walletAddresses []string) ([]*model.User, error) {
	if len(walletAddresses) == 0 {
		return []*model.User{}, nil
	}

	commonUsers, err := a.listUsersByWalletAddresses(walletAddresses)
	if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
		return nil, err
	}

	return commonUsers, a.log.CheckError(err, a.ListUsersByWalletAddresses)
}

func (a API) listUsersByWalletAddresses(walletAddresses []string) ([]*model.User, error) {
	response, err := a.makeListUserByWalletsRequest(walletAddresses)
	if a.log.CheckError(err, a.listUsersByWalletAddresses) != nil {
		return nil, err
	}

	listUsersResponse, err := a.getListUsersResponse(response)
	if a.log.CheckError(err, a.listUsersByWalletAddresses) != nil {
		return nil, err
	}

	if listUsersResponse.Users == nil {
		return nil, errors.New("users not found")
	}

	for _, user := range listUsersResponse.Users {
		go func(user *model.User) { a.cache.SetUserEmailToCache(user) }(user)
	}

	return listUsersResponse.Users, nil
}

func (a API) makeListUserByWalletsRequest(walletAddresses []string) (*resty.Response, error) {
	resp, err := a.restyClient.R().EnableTrace().
		SetQueryParamsFromValues(url.Values{"wallets[]": walletAddresses}).
		SetHeader("admin-token", a.authAdminToken).
		Get(fmt.Sprintf(getUserEndpoint, a.authApiUrl))

	return resp, a.log.CheckError(err, a.makeListUserByWalletsRequest)
}

func (a API) getGetUserResponse(resp *resty.Response) (*GetUserResponse, error) {
	responseStruct := new(GetUserResponse)
	err := json.Unmarshal(resp.Body(), responseStruct)

	return responseStruct, a.log.CheckError(err, a.getGetUserResponse)
}

func (a API) getListUsersResponse(resp *resty.Response) (*ListUsersResponse, error) {
	responseStruct := new(ListUsersResponse)
	err := json.Unmarshal(resp.Body(), responseStruct)

	return responseStruct, a.log.CheckError(err, a.getGetUserResponse)
}
