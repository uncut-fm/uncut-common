package model

import (
	"github.com/uncut-fm/uncut-common/pkg/config"
)

type BlockchainNetwork string

func (b BlockchainNetwork) String() string {
	return string(b)
}

var (
	PolygonNetwork                  BlockchainNetwork = "polygon"
	PolygonMainnetBlockchainNetwork BlockchainNetwork = "polygon-mainnet"
	PolygonMumbaiBlockchainNetwork  BlockchainNetwork = "polygon-mumbai"

	EthereumMainnetBlockchainNetwork BlockchainNetwork = "eth-mainnet"
	EthereumGoerliBlockchainNetwork  BlockchainNetwork = "eth-goerli"

	ArbitrumMainnetBlockchainNetwork BlockchainNetwork = "arb-mainnet"
	ArbitrumGoerliBlockchainNetwork  BlockchainNetwork = "arb-goerli"

	OptimismMainnetBlockchainNetwork BlockchainNetwork = "opt-mainnet"
	OptimismGoerliBlockchainNetwork  BlockchainNetwork = "opt-goerli"

	WaxBlockchainNetwork BlockchainNetwork = "wax"

	WAXPTokenSymbol TokenSymbol = "WAXP"
	ETHTokenSymbol  TokenSymbol = "ETH"
)

type TokenSymbol string

func GetBlockchainNetworksByEnvironment(env string) []BlockchainNetwork {
	switch env {
	case config.LocalEnvironment, config.DevEnvironment, config.TestEnvironment:
		return []BlockchainNetwork{PolygonMumbaiBlockchainNetwork, EthereumGoerliBlockchainNetwork, ArbitrumGoerliBlockchainNetwork, OptimismGoerliBlockchainNetwork}
	case config.StageEnvironment, config.ProdEnvironment:
		return []BlockchainNetwork{PolygonMainnetBlockchainNetwork, EthereumMainnetBlockchainNetwork, ArbitrumMainnetBlockchainNetwork, OptimismMainnetBlockchainNetwork}
	default:
		return []BlockchainNetwork{PolygonMainnetBlockchainNetwork, EthereumMainnetBlockchainNetwork, ArbitrumMainnetBlockchainNetwork, OptimismMainnetBlockchainNetwork}
	}
}

type WaxCollectionInfo struct {
	CollectionName     string   `json:"collection_name"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	CreatedAtBlock     string   `json:"created_at_block"`
	CreatedAtTime      string   `json:"created_at_time"`
	Image              string   `json:"img"`
	Author             string   `json:"author"`
	AuthorizedAccounts []string `json:"authorized_accounts"`
	NotifyAccounts     []string `json:"notify_accounts"`
	MarketFee          float64  `json:"market_fee"`
	AllowNotify        bool     `json:"allow_notify"`
}

func (w WaxCollectionInfo) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"collection_name":     w.CollectionName,
		"name":                w.Name,
		"description":         w.Description,
		"created_at_block":    w.CreatedAtBlock,
		"created_at_time":     w.CreatedAtTime,
		"img":                 w.Image,
		"author":              w.Author,
		"authorized_accounts": w.AuthorizedAccounts,
		"notify_accounts":     w.NotifyAccounts,
		"market_fee":          w.MarketFee,
		"allow_notify":        w.AllowNotify,
	}
}
