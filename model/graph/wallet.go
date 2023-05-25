package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Wallet struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	WalletAddress   string    `json:"walletAddress"`
	Provider        string    `json:"provider"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	BecamePrimaryAt time.Time `json:"becamePrimaryAt"`
	LastSyncedAt    time.Time `json:"lastSyncedAt"`

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
		CreatedAt:       timestamppb.New(w.CreatedAt),
		UpdatedAt:       timestamppb.New(w.UpdatedAt),
		LastSyncedAt:    timestamppb.New(w.LastSyncedAt),
		BecamePrimaryAt: timestamppb.New(w.BecamePrimaryAt),
	}
}

func NewWalletFromCommonWallet(commonWallet *common_model.Wallet, user User) *Wallet {
	return &Wallet{
		ID:            commonWallet.ID,
		Name:          commonWallet.Name,
		Description:   commonWallet.Description,
		WalletAddress: commonWallet.WalletAddress,
		Provider:      commonWallet.Provider,
		CreatedAt:     commonWallet.CreatedAt,
		UpdatedAt:     commonWallet.UpdatedAt,
		LastSyncedAt:  commonWallet.LastSyncedAt,
		User:          &user,
	}
}

// GetPropertiesInMap returns a map of the wallet's properties; keys are in camelCase
func (w *Wallet) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              w.ID,
		"name":            w.Name,
		"description":     w.Description,
		"walletAddress":   w.WalletAddress,
		"provider":        w.Provider,
		"createdAt":       w.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":       w.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"lastSyncedAt":    w.LastSyncedAt.Format("2006-01-02 15:04:05 MST"),
		"becamePrimaryAt": w.BecamePrimaryAt.Format("2006-01-02 15:04:05 MST"),
	}
}
