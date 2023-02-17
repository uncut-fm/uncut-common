package model

import (
	"fmt"
	"strings"
)

var transactionStatusKeyPattern = "%s_status" // {txHash}_status

func GetTransactionStatusKey(txHash string) string {
	return fmt.Sprintf(transactionStatusKeyPattern, strings.ToLower(txHash))
}

type TransactionStatus int

const (
	TransactionStatusNotProcessed TransactionStatus = iota
	TransactionStatusProcessing
	TransactionStatusProcessed
	TransactionStatusTracked
)
