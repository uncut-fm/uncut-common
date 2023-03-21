package alchemy

import (
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"github.com/uncut-fm/uncut-common/pkg/logger"
)

const requestTimeout = 5 * time.Second

var (
	getOwnedNftsURLPattern = "https://%s.g.alchemy.com/nft/v2/%s/getNFTs" // https://{blockchain_network}.g.alchemy.com/nft/v2/{apiKey}/getNFTs
	rpcURLPattern          = "https://%s.g.alchemy.com/v2/%s"             // https://{blockchain_network}.g.alchemy.com/v2/{apiKey}
)

type Client struct {
	log                 logger.Logger
	alchemyAPIKey       string
	networks            []model.BlockchainNetwork
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
	client := &Client{
		log:                 log,
		alchemyAPIKey:       alchemyAPIKey,
		networks:            model.GetBlockchainNetworksByEnvironment(env),
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

func (c Client) getBlockchainNetworkByCommonName(networkToFind model.BlockchainNetwork) model.BlockchainNetwork {
	for _, n := range c.networks {
		if strings.Contains(n.String(), networkToFind.String()) {
			return n
		}
	}

	return c.networks[0]
}
