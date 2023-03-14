package alchemy

import (
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"github.com/uncut-fm/uncut-common/pkg/logger"
)

const requestTimeout = 5 * time.Second

var (
	getOwnedNftsURLPattern     = "https://%s.g.alchemy.com/nft/v2/%s/getNFTs" // https://{blockchain_network}.g.alchemy.com/nft/v2/{apiKey}/getNFTs
	getTokenBalancesURLPattern = "https://%s.g.alchemy.com/v2/%s"             // https://{blockchain_network}.g.alchemy.com/v2/{apiKey}
)

type Client struct {
	log                 logger.Logger
	alchemyAPIKey       string
	polygonNetwork      model.BlockchainNetwork
	ethereumNetwork     model.BlockchainNetwork
	restyClient         *resty.Client
	currencies          config.Web3Currencies
	cachedBalances      map[string]cachedBalancesStruct
	cachedBalancesMutex *sync.RWMutex
}

type cachedBalancesStruct struct {
	balances      []model.Balance
	retrievedTime time.Time
}

func (c cachedBalancesStruct) isOlderThan5min() bool {
	return c.retrievedTime.Add(5 * time.Minute).Before(time.Now())
}

func NewClient(log logger.Logger, currencies config.Web3Currencies, alchemyAPIKey, env string) *Client {
	polygonNetwork, ethNetwork := model.GetBlockchainNetworksByEnvironment(env)
	client := &Client{
		log:                 log,
		alchemyAPIKey:       alchemyAPIKey,
		polygonNetwork:      polygonNetwork,
		ethereumNetwork:     ethNetwork,
		currencies:          currencies,
		cachedBalances:      make(map[string]cachedBalancesStruct),
		cachedBalancesMutex: new(sync.RWMutex),
		restyClient:         createRestyClient(),
	}

	return client
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(requestTimeout)

	return client
}
