package cache

import (
	"context"
	"github.com/uncut-fm/uncut-common/model"
	"sync"
	"time"
)

type InmemoryCache struct {
	cachedTokenPrice      map[model.TokenSymbol]*cachedPriceStruct
	cachedTokenPriceMutex *sync.RWMutex
}

func NewInmemoryCache() *InmemoryCache {
	return &InmemoryCache{
		cachedTokenPrice:      make(map[model.TokenSymbol]*cachedPriceStruct),
		cachedTokenPriceMutex: new(sync.RWMutex),
	}
}

type cachedPriceStruct struct {
	price         float64
	retrievedTime time.Time
}

func (c cachedPriceStruct) isOlderThan10min() bool {
	return c.retrievedTime.Add(10 * time.Minute).Before(time.Now())
}

func (c *InmemoryCache) GetTokenPrice(ctx context.Context, token model.TokenSymbol) (float64, bool) {
	c.cachedTokenPriceMutex.RLock()
	defer c.cachedTokenPriceMutex.RUnlock()

	cachedPrice := c.cachedTokenPrice[token]
	if cachedPrice == nil {
		return 0, false
	}

	if cachedPrice.isOlderThan10min() {
		return cachedPrice.price, false
	}

	return cachedPrice.price, true
}

func (c *InmemoryCache) SetTokenPrice(ctx context.Context, token model.TokenSymbol, price float64) {
	c.cachedTokenPriceMutex.Lock()
	defer c.cachedTokenPriceMutex.Unlock()

	c.cachedTokenPrice[token] = &cachedPriceStruct{
		price:         price,
		retrievedTime: time.Now(),
	}
}
