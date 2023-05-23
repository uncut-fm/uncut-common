package model

import (
	"time"
)

type Transaction struct {
	ID                  int
	Event               string
	PaperTransactionId  string
	PaperCheckoutId     string
	BuyerWalletAddress  string
	BuyerEmail          string
	NetworkFeeUsd       float64
	TotalPriceUsd       float64
	PaymentCompletedAt  time.Time
	TransferCompletedAt time.Time
	PaymentMethod       string
	CreatedAt           time.Time
	PaperCreatedAt      time.Time
	UpdatedAt           time.Time
	WalletType          string

	NFTOwner      *NFTOwner
	NFT           *NFT
	User          *User
	TransferredTo *Wallet
}
