package auth

import (
	"context"
	"github.com/cenkalti/backoff"
	"github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/grpc/metadata"
	"net"
	"time"
)

func (a API) GetOrCreateUser(ctx context.Context, email string) (*GetOrCreateUserResponse, error) {
	response, err := a.grpcClient.GetOrCreateUserAsCreator(a.addAdminTokenToGrpcCtx(ctx), &proto_user.EmailRequest{Email: email})
	if a.log.CheckError(err, a.GetOrCreateUser) != nil {
		return nil, err
	}

	return getGetOrCreateUserResponse(response), nil
}

func getGetOrCreateUserResponse(resp *proto_user.GetOrCreateUserResponse) *GetOrCreateUserResponse {
	return &GetOrCreateUserResponse{
		User:          model.ParseProtoUserToUser(resp.GetUser()),
		ExistedBefore: resp.GetExistedBefore(),
	}
}

func (a API) GetNftCreators(ctx context.Context) ([]*model.User, error) {
	protoUsers, err := a.grpcClient.ListNftCreators(a.addAdminTokenToGrpcCtx(ctx), &proto_user.Empty{})

	if a.log.CheckError(err, a.GetNftCreators) != nil {
		return nil, err
	}

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), nil
}

func (a API) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	protoUser, err := a.grpcClient.GetUserByEmail(a.addAdminTokenToGrpcCtx(ctx), &proto_user.EmailRequest{Email: email})
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), nil
}

func (a API) GetUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	return a.getUserByWalletAddress(ctx, walletAddress)
}

func (a API) getUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	protoUser, err := a.grpcClient.GetUserByWalletAddress(a.addAdminTokenToGrpcCtx(ctx), &proto_user.WalletAddressRequest{WalletAddress: walletAddress})
	if a.log.CheckError(err, a.getUserByWalletAddress) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), nil
}

func (a API) GetUserByID(ctx context.Context, userID int) (*model.User, error) {
	var (
		err       error
		protoUser *proto_user.User
	)

	operation := func() error {
		protoUser, err = a.grpcClient.GetUserByID(a.addAdminTokenToGrpcCtx(ctx), &proto_user.IDRequest{Id: uint64(userID)})
		if opErr, ok := err.(*net.OpError); ok {
			if opErr.Err.Error() == "connection reset by peer" {
				// Connection closed by server
				// Wait for a short period before attempting to reconnect
				return err
			}
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 2 * time.Second

	_ = backoff.Retry(operation, b)

	if a.log.CheckError(err, a.GetUserByID) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), err
}

func (a API) GetUserEmailByWalletAddress(ctx context.Context, walletAddress string) (string, error) {
	email, exist := a.cache.GetEmailByWalletCache(walletAddress)
	if exist {
		return email, nil
	}

	user, err := a.getUserByWalletAddress(a.addAdminTokenToGrpcCtx(ctx), walletAddress)
	if a.log.CheckError(err, a.GetUserEmailByWalletAddress) != nil {
		return "", err
	}

	return user.Email, nil
}

func (a API) ListUsersByWalletAddresses(ctx context.Context, walletAddresses []string) ([]*model.User, error) {
	protoUsers, err := a.grpcClient.ListUsersByWalletAddresses(a.addAdminTokenToGrpcCtx(ctx), &proto_user.WalletAddressesRequest{WalletAddresses: walletAddresses})
	if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
		return nil, err
	}

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), nil
}

func (a API) ListWalletsByUserID(ctx context.Context, userID int) ([]*model.Wallet, error) {
	var (
		err                  error
		protoWalletsResponse *proto_user.WalletsResponse
	)

	operation := func() error {
		protoWalletsResponse, err = a.grpcClient.ListWalletsByUserID(a.addAdminTokenToGrpcCtx(ctx), &proto_user.IDRequest{Id: uint64(userID)})
		if opErr, ok := err.(*net.OpError); ok {
			if opErr.Err.Error() == "connection reset by peer" {
				// Connection closed by server
				// Wait for a short period before attempting to reconnect
				return err
			}
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 2 * time.Second

	_ = backoff.Retry(operation, b)

	if a.log.CheckError(err, a.ListWalletsByUserID) != nil {
		return nil, err
	}

	return model.ParseProtoWalletsToWallets(protoWalletsResponse.Wallets), nil
}

func (a API) addAdminTokenToGrpcCtx(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "admin-token", a.authAdminToken)
}
