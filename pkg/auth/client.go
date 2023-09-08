package auth

import (
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/grpc"
	"time"
)

var (
	getOrCreateUserEndpoint = "%s/user/get-or-create"
	getCreatorsEndpoint     = "%s/user/creators"
	updateUserEndpoint      = "%s/user/admin"
	getUserEndpoint         = "%s/user/"
	getWalletsEndpoint      = "%s/user/wallets/%d"
)

const requestTimeout = 5 * time.Second

type API struct {
	log            logger.Logger
	restyClient    *resty.Client
	userClient     proto_user.UsersClient
	authClient     proto_user.AuthClient
	authApiUrl     string
	authAdminToken string
}

func NewAPI(l logger.Logger, authApiUrl, authAdminToken string, grpcConn *grpc.ClientConn) *API {
	return &API{
		log:            l,
		authApiUrl:     authApiUrl,
		authAdminToken: authAdminToken,
		restyClient:    createRestyClient(),
		userClient:     proto_user.NewUsersClient(grpcConn),
		authClient:     proto_user.NewAuthClient(grpcConn),
	}
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout)

	return client
}
