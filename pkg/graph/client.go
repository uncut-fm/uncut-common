package graph

import (
	"context"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/proto/graph"
	"google.golang.org/grpc"
)

type API struct {
	log             logger.Logger
	grpcClient      graph.NetworkClient
	graphAdminToken string
}

func NewAPI(l logger.Logger, graphAdminToken string, grpcConn *grpc.ClientConn) *API {
	return &API{
		log:             l,
		graphAdminToken: graphAdminToken,
		grpcClient:      graph.NewNetworkClient(grpcConn),
	}
}

func (a API) GetUsersNetworkByID(ctx context.Context, userID int) (*graph.NetworkMembersInfoResponse, error) {
	response, err := a.grpcClient.GetNetworkByUserID(ctx, &graph.GetNetworkByUserIDRequest{UserId: uint64(userID)})

	return response, a.log.CheckError(err, a.GetUsersNetworkByID)
}
