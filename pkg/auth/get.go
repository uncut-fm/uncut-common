package auth

import (
	"context"
	"github.com/cenkalti/backoff"
	"github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/grpc/metadata"
	"net"
	"sync"
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
	WalletProviders   []string
	IncludeEmptyUsers bool
}

func (a API) ListAll(ctx context.Context, filters *UsersFilters, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
	req := &proto_user.ListAllUsersRequest{}
	if filters != nil {
		req.Filters = &proto_user.UserFilters{
			WalletProviders:   filters.WalletProviders,
			IncludeEmptyUsers: true,
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
			WalletProviders:   filters.WalletProviders,
			IncludeEmptyUsers: filters.IncludeEmptyUsers,
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
	mainResponse := &UsersInfoResponse{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	// to collect all errors
	var errs []error
	var errMu sync.Mutex

	const batchSize = 1000

	// load in batches of 1000 to avoid grpc error "message too large"
	// run in parallel and merge results
	for i := 0; i < len(walletAddresses); i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + batchSize
			if end > len(walletAddresses) {
				end = len(walletAddresses)
			}

			response, err := a.listUsersByWalletAddresses(ctx, walletAddresses[start:end], orderBy, pagination)
			if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
				return
			}

			mu.Lock()
			mainResponse.TotalCount += response.TotalCount
			mainResponse.Users = append(mainResponse.Users, response.Users...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return mainResponse, nil
}

func (a API) listUsersByWalletAddresses(ctx context.Context, walletAddresses []string, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
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
	mainResponse := &UsersInfoResponse{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	// to collect all errors
	var errs []error
	var errMu sync.Mutex

	// Decide on a reasonable batch size. Here, I'm using 1000 as an example.
	const batchSize = 1000

	// load in batches of batchSize
	// run in parallel and merge results
	for i := 0; i < len(userIDs); i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()

			end := start + batchSize
			if end > len(userIDs) {
				end = len(userIDs)
			}

			partialResponse, err := a.listUsersByIDs(ctx, userIDs[start:end], orderBy, pagination)
			if a.log.CheckError(err, a.ListUsersByWalletAddresses) != nil {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
				return
			}

			mu.Lock()
			mainResponse.TotalCount += partialResponse.TotalCount
			mainResponse.Users = append(mainResponse.Users, partialResponse.Users...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(errs) > 0 {
		// handle the errors as you see fit, for simplicity, we're just returning the first one
		return nil, errs[0]
	}

	return mainResponse, nil
}

func (a API) listUsersByIDs(ctx context.Context, userIDs []int, orderBy *model.UserOrder, pagination *model.OffsetPaginationInput) (*UsersInfoResponse, error) {
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

	return &UsersInfoResponse{
		TotalCount: int(protoUsersInfo.TotalCount),
		Users:      model.ParseProtoUsersToCommonUsers(protoUsersInfo.Users),
	}, nil
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
