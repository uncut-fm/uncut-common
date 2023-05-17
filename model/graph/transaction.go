package model

import (
	"github.com/mindstand/gogm/v2"
	"time"
)

type Transaction struct {
	gogm.BaseNode

	ID                  int       `gogm:"name=id,pk"`
	Event               string    `gogm:"name=event"`
	PaperTransactionId  string    `gogm:"name=paper_transaction_id"`
	PaperCheckoutId     string    `gogm:"name=paper_checkout_id"`
	BuyerWalletAddress  string    `gogm:"name=buyer_wallet_address"`
	BuyerEmail          string    `gogm:"name=buyer_email"`
	NetworkFeeUsd       float64   `gogm:"name=network_fee_usd"`
	TotalPriceUsd       float64   `gogm:"name=total_price_usd"`
	PaymentCompletedAt  time.Time `gogm:"name=payment_completed_at"`
	TransferCompletedAt time.Time `gogm:"name=transfer_completed_at"`
	PaymentMethod       string    `gogm:"name=payment_method"`
	CreatedAt           time.Time `gogm:"name=created_at"`
	PaperCreatedAt      time.Time `gogm:"name=paper_created_at"`
	UpdatedAt           time.Time `gogm:"name=updated_at"`
	WalletType          string    `gogm:"name=wallet_type"`

	NFTOwner      *NFTOwner `gogm:"direction=incoming;relationship=HAS"`
	NFT           *NFT      `gogm:"direction=incoming;relationship=INVOLVES"`
	User          *User     `gogm:"direction=incoming;relationship=PERFORMED"`
	TransferredTo *Wallet   `gogm:"direction=incoming;relationship=TRANSFERRED_TO"`
}
