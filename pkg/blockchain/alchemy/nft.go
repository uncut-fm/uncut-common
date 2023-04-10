package alchemy

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/uncut-fm/uncut-common/model"
	"sync"
	"time"
)

type OwnedNfts struct {
	Network model.BlockchainNetwork
	NFTs    []Nft
}

type TokenType string

var (
	TokenTypeErc1155 TokenType = "ERC1155"
	TokenTypeErc721  TokenType = "ERC721"
)

type Nft struct {
	Contract struct {
		Address string `json:"address"`
	} `json:"contract"`
	ID struct {
		TokenID       string `json:"tokenId"`
		TokenMetadata struct {
			TokenType TokenType `json:"tokenType"`
		} `json:"tokenMetadata"`
	} `json:"id"`
	Balance     string `json:"balance"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TokenURI    struct {
		Gateway string `json:"gateway"`
		Raw     string `json:"raw"`
	} `json:"tokenUri"`
	Media []struct {
		Gateway   string `json:"gateway"`
		Thumbnail string `json:"thumbnail"`
		Raw       string `json:"raw"`
		Format    string `json:"format"`
		Bytes     int    `json:"bytes"`
	} `json:"media"`
	Metadata struct {
		Image        string `json:"image"`
		ExternalURL  string `json:"external_url"`
		AnimationURL string `json:"animation_url"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Attributes   []struct {
			Value     interface{} `json:"value"`
			TraitType string      `json:"trait_type"`
		} `json:"attributes"`
	} `json:"metadata"`
	TimeLastUpdated  time.Time `json:"timeLastUpdated"`
	ContractMetadata struct {
		Name                string `json:"name"`
		Symbol              string `json:"symbol"`
		TokenType           string `json:"tokenType"`
		ContractDeployer    string `json:"contractDeployer"`
		DeployedBlockNumber int    `json:"deployedBlockNumber"`
		OpenSea             *struct {
			CollectionName        *string   `json:"collectionName,omitempty"`
			SafelistRequestStatus string    `json:"safelistRequestStatus"`
			ImageURL              *string   `json:"imageUrl,omitempty"`
			Description           *string   `json:"description,omitempty"`
			ExternalURL           string    `json:"externalUrl"`
			LastIngestedAt        time.Time `json:"lastIngestedAt"`
			FloorPrice            *float64  `json:"floorPrice,omitempty"`
		} `json:"openSea,omitempty"`
	} `json:"contractMetadata"`
	SpamInfo struct {
		IsSpam          string   `json:"isSpam"`
		Classifications []string `json:"classifications"`
	} `json:"spamInfo"`
	Error string `json:"error"`
}

func (a Nft) GetTokenID() int {
	return hexToInt(a.ID.TokenID)
}

func (a Nft) GetTokenIDString() string {
	return hexToNumString(a.ID.TokenID)
}

func (c Client) ListNftsOwnedByWalletAddress(ctx context.Context, walletAddress string) ([]OwnedNfts, error) {
	ownedNFTs := []OwnedNfts{}

	wg := new(sync.WaitGroup)

	for _, network := range c.networks {
		wg.Add(1)
		go func(n model.BlockchainNetwork) {
			defer wg.Done()

			ownedNetworkNfts := OwnedNfts{Network: n}

			var (
				pageKey *string
				nfts    []Nft
				err     error
			)

			for {
				nfts, pageKey, err = c.makeGetOwnedNftsRequest(ctx, walletAddress, n, pageKey)
				_ = c.log.CheckError(err, c.ListNftsOwnedByWalletAddress)

				ownedNetworkNfts.NFTs = append(ownedNetworkNfts.NFTs, nfts...)
				if pageKey == nil {
					break
				}
			}

			ownedNFTs = append(ownedNFTs, ownedNetworkNfts)
		}(network)
	}

	wg.Wait()

	return ownedNFTs, nil
}

func (c Client) makeGetOwnedNftsRequest(ctx context.Context, walletAddress string, network model.BlockchainNetwork, pageKey *string) ([]Nft, *string, error) {
	var err error

	response := new(getOwnedNftsResponse)
	operation := func() error {
		req := c.restyClient.R().EnableTrace().
			SetResult(response).
			SetQueryParams(map[string]string{
				"owner":        walletAddress,
				"withMetadata": "true",
			})

		if pageKey != nil {
			req.SetQueryParam("pageKey", *pageKey)
		}

		_, err = req.Get(c.getNftOwnersReqURL(network))

		return c.log.CheckError(err, c.makeGetOwnedNftsRequest)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 5 * time.Second

	err = backoff.Retry(operation, b)

	return response.OwnedNfts, response.PageKey, c.log.CheckError(err, c.makeGetOwnedNftsRequest)
}

func (c Client) getNftOwnersReqURL(network model.BlockchainNetwork) string {
	return fmt.Sprintf(getOwnedNftsURLPattern, network, c.alchemyAPIKey)
}
