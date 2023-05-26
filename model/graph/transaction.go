package model

type Transaction struct {
	ID                  int     `json:"id"`
	Event               string  `json:"event"`
	PaperTransactionId  string  `json:"paperTransactionId"`
	PaperCheckoutId     string  `json:"paperCheckoutId"`
	BuyerWalletAddress  string  `json:"buyerWalletAddress"`
	BuyerEmail          string  `json:"buyerEmail"`
	NetworkFeeUsd       float64 `json:"networkFeeUsd"`
	TotalPriceUsd       float64 `json:"totalPriceUsd"`
	PaymentCompletedAt  int64   `json:"paymentCompletedAt"`
	TransferCompletedAt int64   `json:"transferCompletedAt"`
	PaymentMethod       string  `json:"paymentMethod"`
	CreatedAt           int64   `json:"createdAt"`
	PaperCreatedAt      int64   `json:"paperCreatedAt"`
	UpdatedAt           int64   `json:"updatedAt"`
	WalletType          string  `json:"walletType"`

	NFTOwner      *NFTOwner `json:"-"`
	NFT           *NFT      `json:"-"`
	User          *User     `json:"-"`
	TransferredTo *Wallet   `json:"-"`
}
