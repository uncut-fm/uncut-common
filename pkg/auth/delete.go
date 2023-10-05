package auth

import (
	"context"
	"github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
)

func (a API) DeleteWallet(ctx context.Context, input *DeleteWalletRequest) error {
	_, err := a.userClient.DeleteWallet(a.addAdminTokenToGrpcCtx(ctx), &user.DeleteWalletRequest{
		UserId:   uint64(input.UserID),
		WalletId: uint64(input.WalletID),
	})

	return err
}

func (a API) DeleteUser(ctx context.Context, userID int) error {
	_, err := a.userClient.DeleteUser(a.addAdminTokenToGrpcCtx(ctx), &user.IDRequest{Id: uint64(userID)})
	return err
}
