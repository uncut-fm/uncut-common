package alchemy

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (c *Client) GetTimestampByBlockNumberAndNetwork(ctx context.Context, blockNumber int, network model.BlockchainNetwork) (*time.Time, error) {
	block, err := c.makeGetBlockByNumberRequest(ctx, blockNumber, c.getBlockchainNetworkByCommonName(network))
	if c.log.CheckError(err, c.GetTimestampByBlockNumberAndNetwork) != nil {
		return nil, err
	}

	blockTimestampBigInt := hexToBigInt(block.Result.Timestamp)

	blockTimestamp := time.Unix(blockTimestampBigInt.Int64(), 0)

	return &blockTimestamp, nil
}

func (c Client) makeGetBlockByNumberRequest(ctx context.Context, blockNumber int, network model.BlockchainNetwork) (*getBlockResponse, error) {
	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{getHexByInt(blockNumber), true},
		Id:      0,
	}

	var err error

	response := new(getBlockResponse)
	operation := func() error {
		_, err = c.restyClient.R().SetContext(ctx).
			SetBody(request).
			SetResult(response).
			Post(c.getRpcUrl(network))

		return c.log.CheckError(err, c.makeGetBlockByNumberRequest)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 1 * time.Second

	err = backoff.Retry(operation, b)

	return response, c.log.CheckError(err, c.makeGetTokenBalancesRequest)
}

func getHexByInt(i int) string {
	return fmt.Sprintf("0x%x", i)
}
