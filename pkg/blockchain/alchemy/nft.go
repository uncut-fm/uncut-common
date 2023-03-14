package alchemy

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

type OwnedNfts struct {
	PolygonNFTs  []AlchemyNFT
	EthereumNFTs []AlchemyNFT
}

type AlchemyNFT struct {
	Contract struct {
		Address string `json:"address"`
	} `json:"contract"`
	ID struct {
		TokenID       string `json:"tokenId"`
		TokenMetadata struct {
			TokenType string `json:"tokenType"`
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
		Gateway string `json:"gateway"`
		Raw     string `json:"raw"`
	} `json:"media"`
	Metadata struct {
	} `json:"metadata"`
	TimeLastUpdated  time.Time `json:"timeLastUpdated"`
	ContractMetadata struct {
		Name                string `json:"name"`
		Symbol              string `json:"symbol"`
		TokenType           string `json:"tokenType"`
		ContractDeployer    string `json:"contractDeployer"`
		DeployedBlockNumber int    `json:"deployedBlockNumber"`
		OpenSea             struct {
			CollectionName        string    `json:"collectionName"`
			SafelistRequestStatus string    `json:"safelistRequestStatus"`
			ImageURL              string    `json:"imageUrl"`
			Description           string    `json:"description"`
			ExternalURL           string    `json:"externalUrl"`
			LastIngestedAt        time.Time `json:"lastIngestedAt"`
		} `json:"openSea"`
	} `json:"contractMetadata"`
	SpamInfo struct {
		IsSpam          string   `json:"isSpam"`
		Classifications []string `json:"classifications"`
	} `json:"spamInfo"`
}

func (a AlchemyNFT) GetTokenID() int {
	return int(hexToBigInt(a.ID.TokenID).Int64())
}

func (c Client) ListNftsOwnedByWalletAddress(ctx context.Context, walletAddress string) (*OwnedNfts, error) {
	polygonNFTs, err := c.makeGetOwnedNftsRequest(ctx, walletAddress, c.polygonNetwork)
	if c.log.CheckError(err, c.ListNftsOwnedByWalletAddress) != nil {
		return nil, err
	}

	ethereumNFTs, err := c.makeGetOwnedNftsRequest(ctx, walletAddress, c.ethereumNetwork)
	if c.log.CheckError(err, c.ListNftsOwnedByWalletAddress) != nil {
		return nil, err
	}

	return &OwnedNfts{
		PolygonNFTs:  polygonNFTs,
		EthereumNFTs: ethereumNFTs,
	}, nil
}

func (c Client) makeGetOwnedNftsRequest(ctx context.Context, walletAddress string, network model.BlockchainNetwork) ([]AlchemyNFT, error) {
	var err error

	response := new(getOwnedNftsResponse)
	operation := func() error {
		_, err = c.restyClient.R().EnableTrace().
			SetResult(response).
			SetQueryParams(map[string]string{
				"owner":        walletAddress,
				"withMetadata": "true",
			}).
			Get(c.getNftOwnersReqURL(network))

		return c.log.CheckError(err, c.makeGetOwnedNftsRequest)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 1 * time.Second

	err = backoff.Retry(operation, b)

	return response.OwnedNfts, c.log.CheckError(err, c.makeGetOwnedNftsRequest)
}

func (c Client) getNftOwnersReqURL(network model.BlockchainNetwork) string {
	return fmt.Sprintf(getOwnedNftsURLPattern, network, c.alchemyAPIKey)
}
