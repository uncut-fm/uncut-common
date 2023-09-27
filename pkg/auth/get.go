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

type UsersFilters struct {
	WalletProviders []string
}

func (a API) ListAll(ctx context.Context, filters *UsersFilters, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
	req := &proto_user.ListAllUsersRequest{}
	if filters != nil {
		req.Filters = &proto_user.UserFilters{
			WalletProviders: filters.WalletProviders,
		}
	}

	if orderBy != nil {
		req.Order = model.ParseUserOrderToProto(orderBy)
	}

	if pagination != nil {
		req.Pagination = model.ParseOffsetPaginationToProto(pagination)
	}

	protoUsersInfo, err := a.userClient.ListAll(a.addAdminTokenToGrpcCtx(ctx), req)
	if a.log.CheckError(err, a.ListAll) != nil {
		return nil, err
	}

	response := &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}

	return response, nil
}

func (a API) ListUsersWithOutdatedKarma(ctx context.Context) ([]*model.User, error) {
	protoUsers, err := a.userClient.ListUsersWithOutdatedKarma(a.addAdminTokenToGrpcCtx(ctx), &proto_user.Empty{})
	if a.log.CheckError(err, a.ListUsersWithOutdatedKarma) != nil {
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

func (a API) SearchUsers(ctx context.Context, keyword string, pagination *model.OffsetPaginationInput, filters *UsersFilters) (*UsersInfoResponse, error) {
	req := &proto_user.SearchRequest{
		Keyword:    keyword,
		Pagination: model.ParseOffsetPaginationToProto(pagination)}

	if filters != nil {
		req.Filters = &proto_user.UserFilters{
			WalletProviders: filters.WalletProviders,
		}
	}
	protoUsersInfo, err := a.userClient.SearchByKeyword(a.addAdminTokenToGrpcCtx(ctx), req)

	if a.log.CheckError(err, a.SearchUsers) != nil {
		return nil, err
	}

	response := &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}

	return response, nil
}

func (a API) ListUsersByWalletAddresses(ctx context.Context, walletAddresses []string, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
	req := &proto_user.WalletAddressesRequest{WalletAddresses: walletAddresses}
	if orderBy != nil {
		req.Order = model.ParseUserOrderToProto(orderBy)
	}

	if pagination != nil {
		req.Pagination = model.ParseOffsetPaginationToProto(pagination)
	}

	protoUsersInfo, err := a.userClient.ListUsersByWalletAddresses(a.addAdminTokenToGrpcCtx(ctx), req)
	if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
		return nil, err
	}

	response := &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}

	return response, nil
}

func (a API) ListUsersByIDs(ctx context.Context, userIDs []int, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
	req := &proto_user.IDsRequest{Ids: model.IntToUInt64Slice(userIDs)}
	if orderBy != nil {
		req.Order = model.ParseUserOrderToProto(orderBy)
	}

	if pagination != nil {
		req.Pagination = model.ParseOffsetPaginationToProto(pagination)
	}

	protoUsersInfo, err := a.userClient.ListUsersByIDs(a.addAdminTokenToGrpcCtx(ctx), req)
	if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
		return nil, err
	}

	response := &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}

	return response, nil
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
