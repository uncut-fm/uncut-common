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
	apiUrlPattern             = "https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s"
	WAXP          TokenSymbol = "WAXP"
	ETH           TokenSymbol = "ETH"
)

type TokenSymbol string

var fallbackPricePerToken = map[TokenSymbol]float64{
	ETH:  2000,
	WAXP: 0.059,
}

type ExchangerAPI struct {
	log              logger.Logger
	restyClient      *resty.Client
	cachedTokenPrice map[TokenSymbol]*cachedPriceStruct
}

func NewCryptoExchanger(log logger.Logger) *ExchangerAPI {
	return &ExchangerAPI{
		log:              log,
		restyClient:      createRestyClient(),
		cachedTokenPrice: make(map[TokenSymbol]*cachedPriceStruct),
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

func (c *ExchangerAPI) ETHEquivalentInUSD(ethQuantity float64) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(ethQuantity, ETH)

	return usdPrice, c.log.CheckError(err, c.ETHEquivalentInUSD)
}

func (c *ExchangerAPI) WaxEquivalentInUSD(waxQuantity float64) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(waxQuantity, WAXP)

	return usdPrice, c.log.CheckError(err, c.WaxEquivalentInUSD)
}

func (c *ExchangerAPI) getTokenEquivalentInUSD(tokenQuantity float64, token TokenSymbol) (float64, error) {
	tokenPrice := c.getTokenPrice(token)

	usdEquivalent := calculateExchangeValue(tokenPrice, tokenQuantity)

	return usdEquivalent, nil
}

func (c *ExchangerAPI) getTokenPrice(token TokenSymbol) float64 {
	tokenPrice, inCache, isPriceRelevant := c.getTokenPriceFromCache(token)
	if inCache {
		if !isPriceRelevant {
			go func() {
				tokenPrice, err := c.getTokenPriceInUSDFromAPI(token)
				if c.log.CheckError(err, c.getTokenPrice) != nil {
					return
				}

				c.setUSDEquivalentOfTokenCache(token, tokenPrice)
			}()

			return tokenPrice
		}
	}

	tokenPrice, err := c.getTokenPriceInUSDFromAPI(token)
	if c.log.CheckError(err, c.getTokenPrice) != nil {
		return fallbackPricePerToken[token]
	}

	c.setUSDEquivalentOfTokenCache(token, tokenPrice)

	return tokenPrice
}

func (c *ExchangerAPI) getTokenPriceFromCache(token TokenSymbol) (float64, bool, bool) {
	cachedPrice := c.cachedTokenPrice[token]
	if cachedPrice == nil {
		return 0, false, false
	}

	if cachedPrice.isOlderThan10min() {
		return cachedPrice.price, true, false
	}

	return cachedPrice.price, true, true
}

func (c *ExchangerAPI) setUSDEquivalentOfTokenCache(token TokenSymbol, priceInUSD float64) {
	c.cachedTokenPrice[token] = &cachedPriceStruct{
		price:         priceInUSD,
		retrievedTime: time.Now(),
	}
}

func calculateExchangeValue(pricePerItem, items float64) float64 {
	usdEquivalent := pricePerItem * items

	return usdEquivalent
}

func (c ExchangerAPI) getTokenPriceInUSDFromAPI(token TokenSymbol) (float64, error) {
	apiRequest := getRequest(string(token), "USD")

	responseStruct := struct {
		USD float64 `json:"USD"`
	}{}

	err := c.makeRequest(apiRequest, &responseStruct)

	return responseStruct.USD, c.log.CheckError(err, c.getTokenPriceInUSDFromAPI)
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
