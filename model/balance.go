package model

type Balance struct {
	Currency CurrencyType `json:"currency"`
	Balance  float64      `json:"balance"`
}

type CurrencyType string

const (
	CurrencyTypeMatic CurrencyType = "MATIC"
	CurrencyTypeWeth  CurrencyType = "WETH"
	CurrencyTypeCdols CurrencyType = "CDOLS"
)
