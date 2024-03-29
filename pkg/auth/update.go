package auth

import (
	"context"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a API) UpdateUser(ctx context.Context, input *UpdateUserAuthRequest) (*model.User, error) {
	protoUser, err := a.userClient.UpdateUser(a.addAdminTokenToGrpcCtx(ctx), &user.UpdateUserRequest{
		Id:                 uint64(input.ID),
		Name:               input.Name,
		Title:              input.Title,
		Email:              input.Email,
		ProfileImageUrl:    input.ProfileImageURL,
		WalletAddress:      input.WalletAddress,
		TwitterHandle:      input.TwitterHandle,
		IsNftCreator:       input.IsNftCreator,
		Bio:                input.Bio,
		BannerImageUrl:     input.BannerImageURL,
		Location:           input.Location,
		InstagramHandle:    input.InstagramHandle,
		FacebookHandle:     input.FacebookHandle,
		LinkedinHandle:     input.LinkedinHandle,
		DiscordHandle:      input.DiscordHandle,
		WebsiteUrl:         input.WebsiteURL,
		VerificationStatus: input.VerificationStatus,
		Type:               input.Type,
		Karma:              input.Karma,
		KarmaIn_30Days:     input.KarmaIn30Days,
		ThemeColors:        model.ParseThemeColorsToProto(input.ThemeColors),
	})

	if a.log.CheckError(err, a.UpdateUser) != nil {
		return nil, err
	}

	return model.ParseProtoUserToUser(protoUser), err
}

func (a API) UpdateWallet(ctx context.Context, input *UpdateWalletRequest) (*model.Wallet, error) {
	request := &user.UpdateWalletRequest{
		UserId:      uint64(input.UserID),
		WalletId:    uint64(input.WalletID),
		Name:        input.Name,
		Description: input.Description,
		Primary:     input.Primary,
	}

	if input.LastSyncedAt != nil {
		request.LastSyncedAt = timestamppb.New(*input.LastSyncedAt)
	}

	protoWallet, err := a.userClient.UpdateWallet(a.addAdminTokenToGrpcCtx(ctx), request)

	if a.log.CheckError(err, a.UpdateWallet) != nil {
		return nil, err
	}

	return model.ParseProtoWalletToWallet(protoWallet), err
}

func (a API) UpdateKarmasByUserIDs(ctx context.Context, userIDs []uint64, karmas, karmasIn30Days []int32) error {
	_, err := a.userClient.UpdateKarmasByUserIDs(a.addAdminTokenToGrpcCtx(ctx), &user.UpdateKarmasByUserIDsRequest{
		UserIds:         userIDs,
		Karmas:          karmas,
		KarmasIn_30Days: karmasIn30Days,
	})

	return a.log.CheckError(err, a.UpdateKarmasByUserIDs)
}
