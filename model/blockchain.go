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
