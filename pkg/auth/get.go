package auth

import (
	"context"
	"github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

func (a API) GetOrCreateUser(ctx context.Context, email string) (*GetOrCreateUserResponse, error) {
	response, err := a.drpcClient.GetOrCreateUserAsCreator(ctx, &proto_user.EmailRequest{Email: email})

	return getGetOrCreateUserResponse(response), a.log.CheckError(err, a.GetOrCreateUser)
}

func getGetOrCreateUserResponse(resp *proto_user.GetOrCreateUserResponse) *GetOrCreateUserResponse {
	return &GetOrCreateUserResponse{
		User:          model.ParseProtoUserToUser(resp.GetUser()),
		ExistedBefore: resp.GetExistedBefore(),
	}
}

func (a API) GetNftCreators(ctx context.Context) ([]*model.User, error) {
	protoUsers, err := a.drpcClient.ListNftCreators(ctx, &proto_user.Empty{})

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), a.log.CheckError(err, a.GetNftCreators)
}

func (a API) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	protoUser, err := a.drpcClient.GetUserByEmail(ctx, &proto_user.EmailRequest{Email: email})

	return model.ParseProtoUserToUser(protoUser), a.log.CheckError(err, a.GetUserByEmail)
}

func (a API) GetUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	return a.getUserByWalletAddress(ctx, walletAddress)
}

func (a API) getUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	protoUser, err := a.drpcClient.GetUserByWalletAddress(ctx, &proto_user.WalletAddressRequest{WalletAddress: walletAddress})

	return model.ParseProtoUserToUser(protoUser), a.log.CheckError(err, a.getUserByWalletAddress)
}

func (a API) GetUserByID(ctx context.Context, userID int) (*model.User, error) {
	protoUser, err := a.drpcClient.GetUserByID(ctx, &proto_user.IDRequest{Id: uint64(userID)})

	return model.ParseProtoUserToUser(protoUser), a.log.CheckError(err, a.GetUserByID)
}

func (a API) GetUserEmailByWalletAddress(ctx context.Context, walletAddress string) (string, error) {
	email, exist := a.cache.GetEmailByWalletCache(walletAddress)
	if exist {
		return email, nil
	}

	user, err := a.getUserByWalletAddress(ctx, walletAddress)
	if a.log.CheckError(err, a.GetUserEmailByWalletAddress) != nil {
		return "", err
	}

	return user.Email, nil
}

func (a API) ListUsersByWalletAddresses(ctx context.Context, walletAddresses []string) ([]*model.User, error) {
	protoUsers, err := a.drpcClient.ListUsersByWalletAddresses(ctx, &proto_user.WalletAddressesRequest{WalletAddresses: walletAddresses})

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), a.log.CheckError(err, a.ListUsersByWalletAddresses)
}

func (a API) ListWalletsByUserID(ctx context.Context, userID int) ([]*model.Wallet, error) {
	protoWalletsResponse, err := a.drpcClient.ListWalletsByUserID(ctx, &proto_user.IDRequest{Id: uint64(userID)})

	return model.ParseProtoWalletsToWallets(protoWalletsResponse.Wallets), a.log.CheckError(err, a.ListWalletsByUserID)
}
