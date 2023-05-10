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
	TokenID     string    `json:"tokenId"`
	TokenType   TokenType `json:"tokenType"`
	Balance     string    `json:"balance"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	TokenUri    string    `json:"tokenUri"`
	Image       struct {
		CachedUrl    string `json:"cachedUrl"`
		OriginalUrl  string `json:"originalUrl"`
		ThumbnailUrl string `json:"thumbnailUrl"`
		ContentType  string `json:"contentType"`
		Size         int    `json:"size"`
	} `json:"image"`
	TimeLastUpdated time.Time `json:"timeLastUpdated"`
	Raw             struct {
		TokenUri string `json:"tokenUri"`
		Metadata struct {
			AnimationURL *string `json:"animation_url,omitempty"`
		} `json:"metadata,omitempty"`
	} `json:"raw"`
	Contract struct {
		Address             string `json:"address"`
		Name                string `json:"name"`
		Symbol              string `json:"symbol"`
		TotalSupply         string `json:"totalSupply"`
		TokenType           string `json:"tokenType"`
		ContractDeployer    string `json:"contractDeployer"`
		DeployedBlockNumber int    `json:"deployedBlockNumber"`
		OpenSeaMetadata     struct {
			CollectionName        *string   `json:"collectionName"`
			SafelistRequestStatus *string   `json:"safelistRequestStatus"`
			ImageURL              *string   `json:"imageUrl"`
			Description           *string   `json:"description"`
			ExternalURL           string    `json:"externalUrl"`
			LastIngestedAt        time.Time `json:"lastIngestedAt"`
			FloorPrice            *float64  `json:"floorPrice"`
		} `json:"openSeaMetadata,omitempty"`
	} `json:"contract"`
	SpamInfo struct {
		IsSpam          string   `json:"isSpam"`
		Classifications []string `json:"classifications"`
	} `json:"spamInfo"`
	Error string `json:"error"`
}

func (a Nft) GetTokenID() int {
	return hexToInt(a.TokenID)
}

func (a Nft) GetTokenIDString() string {
	return hexToNumString(a.TokenID)
}

// ListNftsOwnedByWalletAddress fetches nfts owned by a wallet address in parallel, and sends them to the ownedNftsChan channel
func (c Client) ListNftsOwnedByWalletAddress(ctx context.Context, walletAddress string, ownedNftsChan chan<- OwnedNfts) error {
	wg := new(sync.WaitGroup)

	for _, network := range c.networks {
		wg.Add(1)
		go func(n model.BlockchainNetwork) {
			defer wg.Done()

			var (
				pageKey *string
				nfts    []Nft
				err     error
			)

			for {
				nfts, pageKey, err = c.makeGetOwnedNftsRequest(ctx, walletAddress, n, pageKey)
				_ = c.log.CheckError(err, c.ListNftsOwnedByWalletAddress)

				// send nfts to channel
				ownedNftsChan <- OwnedNfts{
					Network: n,
					NFTs:    nfts,
				}

				if pageKey == nil {
					break
				}

			}

		}(network)
	}

	wg.Wait()

	return nil
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
		if err != nil {
			return c.log.CheckError(err, c.makeGetOwnedNftsRequest)
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 5 * time.Second

	err = backoff.Retry(operation, b)

	return response.OwnedNfts, response.PageKey, c.log.CheckError(err, c.makeGetOwnedNftsRequest)
}

func (c Client) getNftOwnersReqURL(network model.BlockchainNetwork) string {
	return fmt.Sprintf(getOwnedNftsURLPattern, network, c.alchemyAPIKey)
}
