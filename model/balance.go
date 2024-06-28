package model

type Balance struct {
	Currency CurrencySymbol `json:"currency"`
	Balance  float64        `json:"balance"`
}

type CurrencySymbol string

const (
	CurrencySymbolMatic CurrencySymbol = "MATIC"
	CurrencySymbolWEth  CurrencySymbol = "wETH"
	CurrencySymbolCdol  CurrencySymbol = "CDOL"
	CurrencySymbolUsdc  CurrencySymbol = "USDC"
	CurrencySymbolWaxp  CurrencySymbol = "WAXP"
	// Artx is the internal uncut currency
	CurrencySymbolArtx CurrencySymbol = "ARTX"
)
