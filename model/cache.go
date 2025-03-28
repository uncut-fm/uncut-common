package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
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
	return fmt.Sprintf(transactionTransferStatusKeyPattern, strings.ToLower(txHash), strings.ToLower(toAddress))
}

type TransactionStatus int

const (
	TransactionStatusNotProcessed TransactionStatus = iota
	TransactionStatusProcessing
	TransactionStatusProcessed
	TransactionStatusTracked
)

func ListKeysByPatternFromRedis(ctx context.Context, tracer trace.Tracer, redisClient *redis.Client, pattern string) ([]string, error) {
	var err error

	ctx, span := tracer.Start(ctx, "ListKeysByPatternFromRedis")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}

		span.End()
	}()

	var cursor uint64
	var keys []string

	limit := int64(500)

	for {
		var (
			ks        []string
			newCursor uint64
		)

		// Scan returns a slice of keys, a new cursor, and an error.
		ks, newCursor, err = redisClient.Scan(ctx, cursor, pattern, limit).Result()
		if err != nil {
			return keys, err
		}

		keys = append(keys, ks...)
		// When the cursor returns to 0, the iteration is complete.
		if newCursor == 0 {
			break
		}
		cursor = newCursor
	}

	return keys, nil
}
