package auth

import (
	"context"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

func (a API) UpdateUser(ctx context.Context, input *UpdateUserAuthRequest) (*model.User, error) {
	protoUser, err := a.grpcClient.UpdateUser(ctx, &user.UpdateUserRequest{
		Id:              uint64(input.ID),
		Name:            input.Name,
		Email:           input.Email,
		ProfileImageUrl: input.ProfileImageURL,
		WalletAddress:   input.WalletAddress,
		TwitterHandle:   input.TwitterHandle,
		IsNftCreator:    input.IsNftCreator,
	})

	if a.log.CheckError(err, a.UpdateUser) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), err
}
