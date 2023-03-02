package exchanger

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"time"
)

const (
	apiUrlPattern       = "https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s"
	fallbackPricePerEth = 2000
)

type ExchangerAPI struct {
	log            logger.Logger
	restyClient    *resty.Client
	cachedEthPrice *cachedPriceStruct
}

func NewCryptoExchanger(log logger.Logger) *ExchangerAPI {
	return &ExchangerAPI{
		log:         log,
		restyClient: createRestyClient(),
	}
}

func createRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(5 * time.Second)
	client.SetRetryCount(5)
	client.SetRetryWaitTime(time.Second)
	client.AddRetryAfterErrorCondition()

	return client
}

type cachedPriceStruct struct {
	price         float64
	retrievedTime time.Time
}

func (c cachedPriceStruct) isOlderThan10min() bool {
	return c.retrievedTime.Add(10 * time.Minute).Before(time.Now())
}

func (c *ExchangerAPI) ETHEquivalentInUSD(ethPrice float64) (float64, error) {
	usdPrice, err := c.getETHEquivalentInUSD(ethPrice)

	return usdPrice, c.log.CheckError(err, c.ETHEquivalentInUSD)
}

func (c *ExchangerAPI) getETHEquivalentInUSD(ethQuantity float64) (float64, error) {
	var err error

	ethPrice, hasEthPriceInCache, isPriceRelevant := c.getETHPriceCache()
	if !isPriceRelevant {
		ethPrice, err = c.getETHToUSDFromAPI()
		if c.log.CheckError(err, c.getETHEquivalentInUSD) != nil {
			if !hasEthPriceInCache {
				ethPrice = fallbackPricePerEth
			}
		}
	}

	usdEquivalent := calculateExchangeValue(ethPrice, ethQuantity)

	if !isPriceRelevant {
		c.setUSDEquivalentETHCache(ethPrice)
	}

	return usdEquivalent, nil
}

func (c *ExchangerAPI) getETHPriceCache() (float64, bool, bool) {
	cachedPrice := c.cachedEthPrice
	if cachedPrice == nil {
		return 0, false, false
	}

	if cachedPrice.isOlderThan10min() {
		return cachedPrice.price, true, false
	}

	return cachedPrice.price, true, true
}

func (c *ExchangerAPI) setUSDEquivalentETHCache(ethPrice float64) {
	c.cachedEthPrice = &cachedPriceStruct{
		price:         ethPrice,
		retrievedTime: time.Now(),
	}
}

func calculateExchangeValue(pricePerItem, items float64) float64 {
	usdEquivalent := pricePerItem * items

	return usdEquivalent
}

func (c ExchangerAPI) getETHToUSDFromAPI() (float64, error) {
	apiRequest := getRequest("ETH", "USD")

	responseStruct := struct {
		USD float64 `json:"USD"`
	}{}

	err := c.makeRequest(apiRequest, &responseStruct)

	return responseStruct.USD, c.log.CheckError(err, c.getETHToUSDFromAPI)
}

func (c ExchangerAPI) makeRequest(apiURL string, responseStruct interface{}) error {
	var err error

	operation := func() error {
		var resp *resty.Response

		resp, err = c.restyClient.R().EnableTrace().
			Get(apiURL)
		if err != nil {
			return c.log.CheckError(err, c.makeRequest)
		}

		return json.Unmarshal(resp.Body(), responseStruct)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 2 * time.Second

	err = backoff.Retry(operation, b)

	return c.log.CheckError(err, c.makeRequest)
}

func getRequest(from, to string) string {
	return fmt.Sprintf(apiUrlPattern, from, to)
}
