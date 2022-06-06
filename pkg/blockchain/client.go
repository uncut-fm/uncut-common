package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-management-api-2/graph/model"
	"math/big"
	"time"
)

type Client struct {
	log            logger.Logger
	ethClient      *ethclient.Client
	currencies     config.Web3Currencies
	mainCurrency   model.CurrencyType
	wethClient     *Token
	cdolsClient    *Token
	cachedBalances map[string]cachedBalancesStruct
}

type cachedBalancesStruct struct {
	balances      []*model.Balance
	retrievedTime time.Time
}

func (c cachedBalancesStruct) isOlderThan5min() bool {
	return c.retrievedTime.Add(5 * time.Minute).Before(time.Now())
}

func NewClient(log logger.Logger, ethClient *ethclient.Client, currencies config.Web3Currencies, blockchainName string) (*Client, error) {
	client := &Client{
		log:            log,
		ethClient:      ethClient,
		currencies:     currencies,
		cachedBalances: make(map[string]cachedBalancesStruct),
	}

	err := client.newTokenClients()
	if err != nil {
		return nil, err
	}

	switch blockchainName {
	case "Polygon", "Mumbai":
		client.mainCurrency = model.CurrencyTypeMatic
	}

	return client, nil
}

func (c *Client) newTokenClients() error {
	wethAddress := common.HexToAddress(c.currencies.Weth.ContractAddress)
	cdolsAddress := common.HexToAddress(c.currencies.Cdols.ContractAddress)

	wethClient, err := NewToken(wethAddress, c.ethClient)
	if c.log.CheckError(err, c.newTokenClients) != nil {
		return err
	}

	cdolsClient, err := NewToken(cdolsAddress, c.ethClient)
	if c.log.CheckError(err, c.newTokenClients) != nil {
		return err
	}

	c.wethClient = wethClient
	c.cdolsClient = cdolsClient
	return nil
}

func (c *Client) GetBalanceByWalletHexAddress(ctx context.Context, walletHexAddress string) ([]*model.Balance, error) {
	if cachedBalance, ok := c.getCachedBalancePerWallet(walletHexAddress); ok {
		return cachedBalance, nil
	}

	walletAddress := common.HexToAddress(walletHexAddress)

	mainBalance, err := c.getMainCurrencyBalance(ctx, walletAddress)
	if c.log.CheckError(err, c.GetBalanceByWalletHexAddress) != nil {
		return nil, err
	}

	balances := []*model.Balance{mainBalance}

	otherTokenBalances, err := c.getOtherCurrenciesBalance(walletAddress)
	if c.log.CheckError(err, c.GetBalanceByWalletHexAddress) != nil {
		return nil, err
	}

	balances = append(balances, otherTokenBalances...)

	c.setCachedBalancePerWallet(balances, walletHexAddress)

	return balances, nil
}

func (c Client) getMainCurrencyBalance(ctx context.Context, walletAddress common.Address) (*model.Balance, error) {
	balanceWei, err := c.ethClient.BalanceAt(context.Background(), walletAddress, nil)
	if c.log.CheckError(err, c.getMainCurrencyBalance) != nil {
		return nil, err
	}

	balanceEth := wei2Eth(balanceWei)

	return &model.Balance{
		Currency: c.mainCurrency,
		Balance:  balanceEth,
	}, err
}

func (c Client) getOtherCurrenciesBalance(walletAddress common.Address) ([]*model.Balance, error) {
	wethBalance, err := c.getTokenBalance(c.wethClient, walletAddress, model.CurrencyTypeWeth)
	if err != nil {
		return nil, err
	}

	cdolsBalance, err := c.getTokenBalance(c.cdolsClient, walletAddress, model.CurrencyTypeCdols)
	if err != nil {
		return nil, err
	}

	return []*model.Balance{wethBalance, cdolsBalance}, nil
}

func (c Client) getTokenBalance(tokenClient *Token, walletAddress common.Address, currency model.CurrencyType) (*model.Balance, error) {
	balanceWei, err := tokenClient.BalanceOf(&bind.CallOpts{}, walletAddress)
	if c.log.CheckError(err, c.getTokenBalance) != nil {
		return nil, err
	}

	balanceEth := wei2Eth(balanceWei)

	return &model.Balance{
		Currency: currency,
		Balance:  balanceEth,
	}, nil
}

func (c *Client) getCachedBalancePerWallet(walletHexAddress string) ([]*model.Balance, bool) {
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

func (c *Client) setCachedBalancePerWallet(balances []*model.Balance, walletHexAddress string) {
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
