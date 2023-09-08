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
	response, err := a.userClient.GetOrCreateUserAsCreator(a.addAdminTokenToGrpcCtx(ctx), &proto_user.EmailRequest{Email: email})
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
	protoUsers, err := a.userClient.ListNftCreators(a.addAdminTokenToGrpcCtx(ctx), &proto_user.Empty{})

	if a.log.CheckError(err, a.GetNftCreators) != nil {
		return nil, err
	}

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), nil
}

func (a API) ListAll(ctx context.Context) ([]*model.User, error) {
	protoUsers, err := a.userClient.ListAll(a.addAdminTokenToGrpcCtx(ctx), &proto_user.Empty{})

	if a.log.CheckError(err, a.ListAll) != nil {
		return nil, err
	}

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), nil
}

func (a API) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	protoUser, err := a.userClient.GetUserByEmail(a.addAdminTokenToGrpcCtx(ctx), &proto_user.EmailRequest{Email: email})
	if a.log.CheckError(err, a.GetUserByEmail) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), nil
}

func (a API) GetUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	return a.getUserByWalletAddress(ctx, walletAddress)
}

func (a API) getUserByWalletAddress(ctx context.Context, walletAddress string) (*model.User, error) {
	protoUser, err := a.userClient.GetUserByWalletAddress(a.addAdminTokenToGrpcCtx(ctx), &proto_user.WalletAddressRequest{WalletAddress: walletAddress})
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
		protoUser, err = a.userClient.GetUserByID(a.addAdminTokenToGrpcCtx(ctx), &proto_user.IDRequest{Id: uint64(userID)})
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
	user, err := a.getUserByWalletAddress(a.addAdminTokenToGrpcCtx(ctx), walletAddress)
	if a.log.CheckError(err, a.GetUserEmailByWalletAddress) != nil {
		return "", err
	}

	return user.Email, nil
}

func (a API) SearchUsers(ctx context.Context, keyword string, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
	protoUsersInfo, err := a.userClient.SearchByKeyword(a.addAdminTokenToGrpcCtx(ctx), &proto_user.SearchRequest{
		Keyword:    keyword,
		Pagination: model.ParseOffsetPaginationToProto(pagination)})

	if a.log.CheckError(err, a.SearchUsers) != nil {
		return nil, err
	}

	response := &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}

	return response, nil
}

func (a API) ListUsersByWalletAddresses(ctx context.Context, walletAddresses []string) ([]*model.User, error) {
	protoUsers, err := a.userClient.ListUsersByWalletAddresses(a.addAdminTokenToGrpcCtx(ctx), &proto_user.WalletAddressesRequest{WalletAddresses: walletAddresses})
	if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
		return nil, err
	}

	return model.ParseProtoUsersResponseToCommonUsers(protoUsers), nil
}

func (a API) ListUsersByIDs(ctx context.Context, userIDs []int) ([]*model.User, error) {
	protoUsers, err := a.userClient.ListUsersByIDs(a.addAdminTokenToGrpcCtx(ctx), &proto_user.IDsRequest{Ids: model.IntToUInt64Slice(userIDs)})
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
		protoWalletsResponse, err = a.userClient.ListWalletsByUserID(a.addAdminTokenToGrpcCtx(ctx), &proto_user.IDRequest{Id: uint64(userID)})
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

func (a API) GetUserSessionByWalletAddress(ctx context.Context, walletAddress string) (*model.UserSession, error) {
	var (
		err                  error
		protoSessionResponse *proto_user.UserSessionResponse
	)

	operation := func() error {
		protoSessionResponse, err = a.authClient.GetUserSessionByWalletAddress(a.addAdminTokenToGrpcCtx(ctx), &proto_user.WalletAddressRequest{WalletAddress: walletAddress})
		if a.log.CheckError(err, a.GetUserSessionByWalletAddress) != nil {
			return err
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 2 * time.Second

	_ = backoff.Retry(operation, b)

	if a.log.CheckError(err, a.ListWalletsByUserID) != nil {
		return nil, err
	}

	return model.ParseProtoUserSessionToUserSession(protoSessionResponse), nil
}

func (a API) addAdminTokenToGrpcCtx(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "admin-token", a.authAdminToken)
}
