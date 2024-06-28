package exchanger

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"math"
	"sync"
	"time"
)

const (
	apiUrlPattern           = "https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s"
	DefaultFreeNftArtxPrice = 5
)

var fallbackPricePerToken = map[model.TokenSymbol]float64{
	model.ETHTokenSymbol:  2000,
	model.WAXPTokenSymbol: 0.059,
}

type ExchangerAPI struct {
	log                   logger.Logger
	restyClient           *resty.Client
	cachedTokenPrice      map[model.TokenSymbol]*cachedPriceStruct
	cachedTokenPriceMutex *sync.RWMutex
}

func NewCryptoExchanger(log logger.Logger) *ExchangerAPI {
	return &ExchangerAPI{
		log:                   log,
		restyClient:           createRestyClient(),
		cachedTokenPrice:      make(map[model.TokenSymbol]*cachedPriceStruct),
		cachedTokenPriceMutex: new(sync.RWMutex),
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
	usdPrice, err := c.getTokenEquivalentInUSD(ethQuantity, model.ETHTokenSymbol)

	return usdPrice, c.log.CheckError(err, c.ETHEquivalentInUSD)
}

func (c *ExchangerAPI) TokenEquivalentInUSD(tokenQuantity float64, token model.TokenSymbol) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(tokenQuantity, token)

	return usdPrice, c.log.CheckError(err, c.TokenEquivalentInUSD)
}

func (c *ExchangerAPI) convertTokenCurrencyToArtxPrice(tokenAmount float64, currency model.TokenSymbol) (float64, error) {
	usdPrice, err := c.TokenEquivalentInUSD(tokenAmount, currency)
	if err != nil {
		return 0, err
	}

	return math.Ceil(model.ConvertUsdToArtx(usdPrice)), nil
}

func (c *ExchangerAPI) ConvertCurrencyAmountToArtx(currencyAmount float64, currency model.CurrencySymbol) (float64, error) {
	var (
		artxAmount float64
		err        error
	)

	switch currency {
	case model.CurrencySymbolUsdc:
		artxAmount = model.ConvertUsdToArtx(currencyAmount)
	case model.CurrencySymbolWaxp:
		artxAmount, err = c.convertTokenCurrencyToArtxPrice(currencyAmount, model.WAXPTokenSymbol)
	case model.CurrencySymbolWEth:
		artxAmount, err = c.convertTokenCurrencyToArtxPrice(currencyAmount, model.ETHTokenSymbol)
	case model.CurrencySymbolArtx:
		return currencyAmount, nil
	default:
		return currencyAmount, nil
	}

	if err != nil {
		return 0, err
	}

	if artxAmount < DefaultFreeNftArtxPrice {
		return DefaultFreeNftArtxPrice, nil
	}

	return float64(roundToNearestFive(int(artxAmount))), err
}

// RoundToNearestFive rounds a number to the nearest 5
func roundToNearestFive(num int) int {
	remainder := num % 5
	if remainder >= 3 {
		return num + (5 - remainder)
	}
	return num - remainder
}

func (c *ExchangerAPI) getTokenEquivalentInUSD(tokenQuantity float64, token model.TokenSymbol) (float64, error) {
	tokenPrice := c.getTokenPrice(token)

	usdEquivalent := calculateExchangeValue(tokenPrice, tokenQuantity)

	return usdEquivalent, nil
}

func (c *ExchangerAPI) getTokenPrice(token model.TokenSymbol) float64 {
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
		}
		return tokenPrice
	}

	tokenPrice, err := c.getTokenPriceInUSDFromAPI(token)
	if c.log.CheckError(err, c.getTokenPrice) != nil {
		return fallbackPricePerToken[token]
	}

	c.setUSDEquivalentOfTokenCache(token, tokenPrice)

	return tokenPrice
}

func (c *ExchangerAPI) getTokenPriceFromCache(token model.TokenSymbol) (float64, bool, bool) {
	c.cachedTokenPriceMutex.RLock()
	defer c.cachedTokenPriceMutex.RUnlock()

	cachedPrice := c.cachedTokenPrice[token]
	if cachedPrice == nil {
		return 0, false, false
	}

	if cachedPrice.isOlderThan10min() {
		return cachedPrice.price, true, false
	}

	return cachedPrice.price, true, true
}

func (c *ExchangerAPI) setUSDEquivalentOfTokenCache(token model.TokenSymbol, priceInUSD float64) {
	c.cachedTokenPriceMutex.Lock()
	defer c.cachedTokenPriceMutex.Unlock()

	c.cachedTokenPrice[token] = &cachedPriceStruct{
		price:         priceInUSD,
		retrievedTime: time.Now(),
	}
}

func calculateExchangeValue(pricePerItem, items float64) float64 {
	usdEquivalent := pricePerItem * items

	return usdEquivalent
}

func (c ExchangerAPI) getTokenPriceInUSDFromAPI(token model.TokenSymbol) (float64, error) {
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
