package alchemy

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/params"
	"github.com/uncut-fm/uncut-common/model"
	"math/big"
	"time"
)

func (c *Client) GetBalanceByWalletHexAddress(ctx context.Context, walletHexAddress string) ([]model.Balance, error) {
	if cachedBalance, ok := c.getCachedBalancePerWallet(walletHexAddress); ok {
		return cachedBalance, nil
	}

	balances, err := c.getTokenBalances(ctx, walletHexAddress)
	if c.log.CheckError(err, c.GetBalanceByWalletHexAddress) != nil {
		return nil, err
	}

	c.setCachedBalancePerWallet(balances, walletHexAddress)

	return balances, nil
}

func (c Client) getTokenBalances(ctx context.Context, walletAddress string) ([]model.Balance, error) {
	alchemyBalances, err := c.makeGetTokenBalancesRequest(ctx, walletAddress)
	if c.log.CheckError(err, c.getTokenBalances) != nil {
		return nil, err
	}

	var balances []model.Balance
	for _, tokenBalance := range alchemyBalances.Result.TokenBalances {
		var currency model.CurrencySymbol
		switch tokenBalance.ContractAddress {
		case c.currencies.Weth.ContractAddress:
			currency = model.CurrencySymbolWEth
		case c.currencies.Cdols.ContractAddress:
			currency = model.CurrencySymbolCdol
		case c.currencies.Matic.ContractAddress:
			currency = model.CurrencySymbolMatic
		case c.currencies.Usdc.ContractAddress:
			currency = model.CurrencySymbolUsdc
		}

		balance := hexWeiStringToFloatEth(tokenBalance.TokenBalance)

		balances = append(balances, model.Balance{
			Currency: currency,
			Balance:  balance,
		})
	}

	return balances, nil
}

func (c Client) makeGetTokenBalancesRequest(ctx context.Context, walletAddress string) (*getTokenBalancesResponse, error) {
	tokenAddresses := c.currencies.GetAddresses()
	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "alchemy_getTokenBalances",
		Params:  []interface{}{walletAddress, tokenAddresses},
		Id:      42,
	}

	var err error

	response := new(getTokenBalancesResponse)
	operation := func() error {
		_, err = c.restyClient.R().EnableTrace().
			SetBody(request).
			SetResult(response).
			Post(c.getRpcUrl(c.polygonNetwork))

		return c.log.CheckError(err, c.makeGetTokenBalancesRequest)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 1 * time.Second

	err = backoff.Retry(operation, b)

	return response, c.log.CheckError(err, c.makeGetTokenBalancesRequest)
}

func (c *Client) getCachedBalancePerWallet(walletHexAddress string) ([]model.Balance, bool) {
	c.cachedBalancesMutex.RLock()
	defer c.cachedBalancesMutex.RUnlock()

	cachedBalance, ok := c.cachedBalances[walletHexAddress]
	if !ok {
		return nil, false
	}

	if cachedBalance.isOlderThan5min() {
		delete(c.cachedBalances, walletHexAddress)
		return nil, false
	}

	return cachedBalance.balances, true
}

func (c *Client) setCachedBalancePerWallet(balances []model.Balance, walletHexAddress string) {
	c.cachedBalancesMutex.Lock()
	c.cachedBalances[walletHexAddress] = cachedBalancesStruct{
		balances:      balances,
		retrievedTime: time.Now(),
	}
	c.cachedBalancesMutex.Unlock()
}

func (c Client) getRpcUrl(network model.BlockchainNetwork) string {
	return fmt.Sprintf(rpcURLPattern, network, c.alchemyAPIKey)
}

func wei2Eth(wei *big.Int) float64 {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	eth, _ := f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether)).Float64()
	return eth
}

func hexWeiStringToFloatEth(hexString string) float64 {
	balanceBigInt := hexToBigInt(hexString)

	return wei2Eth(balanceBigInt)
}

func hexToInt(hexString string) int {
	bigInt := hexToBigInt(hexString)

	return int(bigInt.Int64())
}
func hexToBigInt(hexString string) *big.Int {
	bigInt := new(big.Int)
	bigInt.SetString(hexString, 0)

	return bigInt
}
