package model

import (
	"errors"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"strconv"
	"strings"
)

type BlockchainNetwork string

func (b BlockchainNetwork) String() string {
	return string(b)
}

var (
	PolygonNetwork                  BlockchainNetwork = "polygon"
	PolygonMainnetBlockchainNetwork BlockchainNetwork = "polygon-mainnet"
	PolygonMumbaiBlockchainNetwork  BlockchainNetwork = "polygon-mumbai"
	PolygonAmoyBlockchainNetwork    BlockchainNetwork = "polygon-amoy"

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
		return []BlockchainNetwork{PolygonAmoyBlockchainNetwork, EthereumGoerliBlockchainNetwork, ArbitrumGoerliBlockchainNetwork, OptimismGoerliBlockchainNetwork}
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

// Name ensures that the name is a valid eosio name
func WaxName(str string) (string, error) {
	str = strings.ToLower(str)

	if len(str) > 12 {
		return "", errors.New("Name too long")
	}

	charmap := ".12345abcdefghijklmnopqrstuvwxyz"

	for _, c := range str {
		if strings.IndexRune(charmap, c) == -1 {
			return "", errors.New("Invalid character in name")
		}
	}

	// Last character can't be a dot
	if str[len(str)-1] == '.' {
		return "", errors.New("Last character can't be a dot")
	}

	return str, nil
}

// CollectionName ensures that the collection name is 12 characters long and a valid eosio name
func WaxCollectionName(str string) (string, error) {
	if len(str) != 12 {
		return "", errors.New("Collection name must be 12 characters while this one is " + strconv.Itoa(len(str)))
	}

	return WaxName(str)
}

type WaxAttribute struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type WaxFormat struct {
	Key  string  `json:"key"`
	Type WaxType `json:"type"`
}

type WaxType string

var (
	StringWaxAttributeType WaxType = "string"
	FloatWaxAttributeType  WaxType = "float"
	Int32WaxAttributeType  WaxType = "int32_t"
)
