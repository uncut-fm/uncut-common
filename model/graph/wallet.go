package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Wallet struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	WalletAddress   string `json:"walletAddress"`
	Provider        string `json:"provider"`
	CreatedAt       int64  `json:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt"`
	BecamePrimaryAt int64  `json:"becamePrimaryAt"`
	LastSyncedAt    int64  `json:"lastSyncedAt"`

	User         *User
	NFTOwners    []*NFTOwner
	Transactions []*Transaction
}

func NewWalletsListFromCommonWallets(commonWallets []*common_model.Wallet, user User) []*Wallet {
	wallets := make([]*Wallet, len(commonWallets))

	// remove the user's wallets to avoid infinite recursion
	user.Wallets = nil

	for i := range commonWallets {
		wallets[i] = NewWalletFromCommonWallet(commonWallets[i], user)
	}

	return wallets
}

func NewProtoWalletsListFromWallets(wallets []*Wallet) []*proto_user.Wallet {
	protoWallets := make([]*proto_user.Wallet, len(wallets))

	for i := range wallets {
		protoWallets[i] = wallets[i].ToProto()
	}

	return protoWallets
}

func (w *Wallet) ToProto() *proto_user.Wallet {
	return &proto_user.Wallet{
		Id:              uint64(w.ID),
		Name:            w.Name,
		Description:     w.Description,
		WalletAddress:   w.WalletAddress,
		Provider:        w.Provider,
		CreatedAt:       timestamppb.New(UnixTimeToTime(w.CreatedAt)),
		UpdatedAt:       timestamppb.New(UnixTimeToTime(w.UpdatedAt)),
		LastSyncedAt:    timestamppb.New(UnixTimeToTime(w.LastSyncedAt)),
		BecamePrimaryAt: timestamppb.New(UnixTimeToTime(w.BecamePrimaryAt)),
	}
}

func NewWalletFromCommonWallet(commonWallet *common_model.Wallet, user User) *Wallet {
	return &Wallet{
		ID:            commonWallet.ID,
		Name:          commonWallet.Name,
		Description:   commonWallet.Description,
		WalletAddress: commonWallet.WalletAddress,
		Provider:      commonWallet.Provider,
		CreatedAt:     commonWallet.CreatedAt.Unix(),
		UpdatedAt:     commonWallet.UpdatedAt.Unix(),
		LastSyncedAt:  commonWallet.LastSyncedAt.Unix(),
		User:          &user,
	}
}
