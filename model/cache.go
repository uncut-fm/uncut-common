package model

import (
	"fmt"
	"strings"
)

var (
	transactionStatusKeyPattern         = "%s_status"    // {txHash}_status
	transactionTransferStatusKeyPattern = "%s_%s_status" // {txHash}_{toAddress}_status
)

func GetTransactionStatusKey(txHash string) string {
	return fmt.Sprintf(transactionStatusKeyPattern, strings.ToLower(txHash))
}

func GetTransferTransactionStatusKey(txHash, toAddress string) string {
	return fmt.Sprintf(transactionTransferStatusKeyPattern, strings.ToLower(txHash), toAddress)
}

type TransactionStatus int

const (
	TransactionStatusNotProcessed TransactionStatus = iota
	TransactionStatusProcessing
	TransactionStatusProcessed
	TransactionStatusTracked
)
