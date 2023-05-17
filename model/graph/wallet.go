package model

import (
	"github.com/mindstand/gogm/v2"
	common_model "github.com/uncut-fm/uncut-common/model"
	"time"
)

type Wallet struct {
	gogm.BaseNode

	ID              int       `gogm:"name=id,pk"`
	Name            string    `gogm:"name=name"`
	Description     string    `gogm:"name=description"`
	WalletAddress   string    `gogm:"name=wallet_address"`
	Provider        string    `gogm:"name=provider"`
	CreatedAt       time.Time `gogm:"name=created_at"`
	UpdatedAt       time.Time `gogm:"name=updated_at"`
	BecamePrimaryAt time.Time `gogm:"name=became_primary_at"`
	LastSyncedAt    time.Time `gogm:"name=last_synced_at"`

	User         *User          `gogm:"direction=incoming;relationship=OWNS"`
	NFTOwners    []*NFTOwner    `gogm:"direction=outgoing;relationship=CONTAINS"`
	Transactions []*Transaction `gogm:"direction=outgoing;relationship=TRANSFERRED_TO"`
}

func NewWalletsListFromCommonWallets(commonWallets []*common_model.Wallet, user *User) []*Wallet {
	wallets := make([]*Wallet, len(commonWallets))

	for i := range commonWallets {
		wallets[i] = NewWalletFromCommonWallet(commonWallets[i], user)
	}

	return wallets
}

func NewWalletFromCommonWallet(commonWallet *common_model.Wallet, user *User) *Wallet {
	return &Wallet{
		ID:            commonWallet.ID,
		Name:          commonWallet.Name,
		Description:   commonWallet.Description,
		WalletAddress: commonWallet.WalletAddress,
		Provider:      commonWallet.Provider,
		CreatedAt:     commonWallet.CreatedAt,
		UpdatedAt:     commonWallet.UpdatedAt,
		LastSyncedAt:  commonWallet.LastSyncedAt,
		User:          user,
	}
}
