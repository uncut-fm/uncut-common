package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"math/big"
	"time"
)

const requestTimeout = 5 * time.Second

type Client struct {
	log            logger.Logger
	alchemyRpcURL  string
	restyClient    *resty.Client
	currencies     config.Web3Currencies
	cachedBalances map[string]cachedBalancesStruct
}

type cachedBalancesStruct struct {
	balances      []model.Balance
	retrievedTime time.Time
}

func (c cachedBalancesStruct) isOlderThan5min() bool {
	return c.retrievedTime.Add(5 * time.Minute).Before(time.Now())
}

func NewClient(log logger.Logger, currencies config.Web3Currencies, alchemyRpcUrl string) (*Client, error) {
	client := &Client{
		log:            log,
		alchemyRpcURL:  alchemyRpcUrl,
		currencies:     currencies,
		cachedBalances: make(map[string]cachedBalancesStruct),
		restyClient:    createRestyClient(),
	}

	return client, nil
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout)

	return client
}

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
		var currency model.CurrencyType
		switch tokenBalance.ContractAddress {
		case c.currencies.Weth.ContractAddress:
			currency = model.CurrencyTypeWeth
		case c.currencies.Cdols.ContractAddress:
			currency = model.CurrencyTypeCdols
		}

		balance := hexWeiStringToFloat(tokenBalance.TokenBalance)

		balances = append(balances, model.Balance{
			Currency: currency,
			Balance:  balance,
		})
	}

	return balances, nil
}

func (c Client) makeGetTokenBalancesRequest(ctx context.Context, walletAddress string) (*getTokenBalancesResponse, error) {
	request := &getTokenBalancesRequest{
		Jsonrpc: "2.0",
		Method:  "alchemy_getTokenBalances",
		Params:  []interface{}{walletAddress, []string{c.currencies.Weth.ContractAddress, c.currencies.Cdols.ContractAddress}},
		Id:      42,
	}

	response := new(getTokenBalancesResponse)

	_, err := c.restyClient.R().EnableTrace().
		SetBody(request).
		SetResult(response).
		Post(c.alchemyRpcURL)

	return response, c.log.CheckError(err, c.makeGetTokenBalancesRequest)
}

func (c *Client) getCachedBalancePerWallet(walletHexAddress string) ([]model.Balance, bool) {
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
	c.cachedBalances[walletHexAddress] = cachedBalancesStruct{
		balances:      balances,
		retrievedTime: time.Now(),
	}
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

func hexWeiStringToFloat(hexString string) float64 {
	balanceBigInt := new(big.Int)
	balanceBigInt.SetString(hexString, 0)

	return wei2Eth(balanceBigInt)
}
