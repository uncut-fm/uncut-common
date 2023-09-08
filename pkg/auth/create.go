package auth

import (
	"context"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

func (a API) AddWallet(ctx context.Context, input *AddWalletRequest) (*model.Wallet, error) {
	protoWallet, err := a.userClient.AddWallet(a.addAdminTokenToGrpcCtx(ctx), &user.AddWalletRequest{
		UserId:        uint64(input.UserID),
		Name:          input.Name,
		Description:   input.Description,
		WalletAddress: input.WalletAddress,
		Provider:      input.Provider,
	})

	if a.log.CheckError(err, a.AddWallet) != nil {
		return nil, err
	}

	return model.ParseProtoWalletToWallet(protoWallet), err
}
