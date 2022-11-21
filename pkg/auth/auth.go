package auth

import (
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"storj.io/drpc"
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
	drpcClient     proto_user.DRPCUsersClient
	cache          Cache
	authApiUrl     string
	authAdminToken string
}

func NewAPI(l logger.Logger, cache Cache, authApiUrl, authAdminToken string, drpcConn drpc.Conn) *API {
	return &API{
		log:            l,
		cache:          cache,
		authApiUrl:     authApiUrl,
		authAdminToken: authAdminToken,
		restyClient:    createRestyClient(),
		drpcClient:     proto_user.NewDRPCUsersClient(drpcConn),
	}
}

type Cache interface {
	GetEmailByIDCache(userID int) (string, bool)
	GetEmailByWalletCache(wallet string) (string, bool)
	SetUserEmailToCache(user *model.User)
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout)

	return client
}
