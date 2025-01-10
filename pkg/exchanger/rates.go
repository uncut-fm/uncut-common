package exchanger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/go-resty/resty/v2"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/tracing"
	"go.opentelemetry.io/otel/trace"
	"math"
	"time"
)

const (
	apiUrlPattern           = "https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s"
	DefaultFreeNftArtxPrice = 5
)

var fallbackPricePerToken = map[model.TokenSymbol]float64{
	model.ETHTokenSymbol:  2500,
	model.WAXPTokenSymbol: 0.04,
}

type ExchangerAPI struct {
	log         logger.Logger
	cache       Cache
	restyClient *resty.Client
}

func NewCryptoExchanger(log logger.Logger, tp trace.TracerProvider, cache Cache) *ExchangerAPI {
	return &ExchangerAPI{
		log:         log,
		cache:       cache,
		restyClient: createRestyClient(tp),
	}
}

func createRestyClient(tp trace.TracerProvider) *resty.Client {
	client := resty.New().
		SetTransport(tracing.NewTransport(tp)).
		SetTimeout(5 * time.Second).
		SetRetryCount(5).
		SetRetryWaitTime(time.Second).
		AddRetryAfterErrorCondition()

	return client
}

type Cache interface {
	GetTokenPrice(ctx context.Context, token model.TokenSymbol) (float64, bool)
	SetTokenPrice(ctx context.Context, token model.TokenSymbol, price float64)
}

func (c *ExchangerAPI) ETHEquivalentInUSD(ctx context.Context, ethQuantity float64) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(ctx, ethQuantity, model.ETHTokenSymbol)

	return usdPrice, c.log.CheckError(err, c.ETHEquivalentInUSD)
}

func (c *ExchangerAPI) TokenEquivalentInUSD(ctx context.Context, tokenQuantity float64, token model.TokenSymbol) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(ctx, tokenQuantity, token)

	return usdPrice, c.log.CheckError(err, c.TokenEquivalentInUSD)
}

func (c *ExchangerAPI) convertTokenCurrencyToArtxPrice(ctx context.Context, tokenAmount float64, currency model.TokenSymbol) (float64, error) {
	usdPrice, err := c.getTokenEquivalentInUSD(ctx, tokenAmount, currency)
	if err != nil {
		return 0, err
	}

	return math.Ceil(model.ConvertUsdToArtx(usdPrice)), nil
}

func (c *ExchangerAPI) ConvertCurrencyAmountToArtx(ctx context.Context, currencyAmount float64, currency model.CurrencySymbol) (float64, error) {
	var (
		artxAmount float64
		err        error
	)

	switch currency {
	case model.CurrencySymbolUsdc:
		artxAmount = model.ConvertUsdToArtx(currencyAmount)
	case model.CurrencySymbolWaxp:
		artxAmount, err = c.convertTokenCurrencyToArtxPrice(ctx, currencyAmount, model.WAXPTokenSymbol)
	case model.CurrencySymbolWEth:
		artxAmount, err = c.convertTokenCurrencyToArtxPrice(ctx, currencyAmount, model.ETHTokenSymbol)
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

func (c *ExchangerAPI) convertUsdToTokenCurrencyPrice(ctx context.Context, usdAmount float64, currency model.TokenSymbol) (float64, error) {
	tokenPriceToUSD, err := c.TokenEquivalentInUSD(ctx, 1, currency)
	if err != nil {
		return 0, err
	}

	return roundFloat64To2DecimalPlaces(usdAmount / tokenPriceToUSD), nil
}

func (c *ExchangerAPI) ConvertArtxAmountToCurrency(ctx context.Context, artxAmount float64, currency model.CurrencySymbol) (float64, error) {
	usdAmount := model.ConvertArtxToUsd(artxAmount)

	switch currency {
	case model.CurrencySymbolUsdc:
		return usdAmount, nil
	case model.CurrencySymbolWaxp:
		waxPrice, err := c.convertUsdToTokenCurrencyPrice(ctx, usdAmount, model.WAXPTokenSymbol)
		if err != nil {
			return 0, err
		}

		return waxPrice, nil
	case model.CurrencySymbolWEth:
		ethPrice, err := c.convertUsdToTokenCurrencyPrice(ctx, usdAmount, model.ETHTokenSymbol)
		if err != nil {
			return 0, err
		}

		return ethPrice, nil
	case model.CurrencySymbolArtx:
		return artxAmount, nil
	default:
		return artxAmount, nil
	}
}

// RoundToNearestFive rounds a number to the nearest 5
func roundToNearestFive(num int) int {
	remainder := num % 5
	if remainder >= 3 {
		return num + (5 - remainder)
	}
	return num - remainder
}

func roundFloat64To2DecimalPlaces(input float64) float64 {
	return math.Round(input*100) / 100
}

func (c *ExchangerAPI) getTokenEquivalentInUSD(ctx context.Context, tokenQuantity float64, token model.TokenSymbol) (float64, error) {
	tokenPrice := c.getTokenPrice(ctx, token)

	usdEquivalent := calculateExchangeValue(tokenPrice, tokenQuantity)

	return usdEquivalent, nil
}

func (c *ExchangerAPI) getTokenPrice(ctx context.Context, token model.TokenSymbol) float64 {
	if token == model.ArtxTokenSymbol {
		return model.ArtxUsdRate
	}

	tokenPrice, isPriceRelevant := c.getTokenPriceFromCache(ctx, token)
	if tokenPrice > 0 {
		if !isPriceRelevant {
			go func() {
				tokenPrice, err := c.getTokenPriceInUSDFromAPI(token)
				if c.log.CheckError(err, c.getTokenPrice) != nil {
					return
				}

				c.setUSDEquivalentOfTokenCache(ctx, token, tokenPrice)
			}()
		}
		return tokenPrice
	}

	tokenPrice, err := c.getTokenPriceInUSDFromAPI(token)
	if c.log.CheckError(err, c.getTokenPrice) != nil {
		return fallbackPricePerToken[token]
	}

	c.setUSDEquivalentOfTokenCache(ctx, token, tokenPrice)

	return tokenPrice
}

func (c *ExchangerAPI) getTokenPriceFromCache(ctx context.Context, token model.TokenSymbol) (float64, bool) {
	return c.cache.GetTokenPrice(ctx, token)
}

func (c *ExchangerAPI) setUSDEquivalentOfTokenCache(ctx context.Context, token model.TokenSymbol, priceInUSD float64) {
	c.cache.SetTokenPrice(ctx, token, priceInUSD)
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
	if err != nil {
		return 0, err
	}

	if responseStruct.USD == 0 {
		return 0, fmt.Errorf("token price is zero")
	}

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
