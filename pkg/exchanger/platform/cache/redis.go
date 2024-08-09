package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	tokenPriceCacheKeyPattern = "token_price_%s" // token symbol
	tokenPriceExpiration      = time.Minute * 10
	tokenPriceTTL             = time.Hour * 24
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

type tokenPriceCache struct {
	Price     float64
	ExpiresAt time.Time
}

func (t tokenPriceCache) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

func (c *RedisCache) GetTokenPrice(ctx context.Context, token string) (float64, bool) {
	key := c.getTokenPriceCacheKey(token)

	tokenCache := tokenPriceCache{}

	cache, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return 0, false
	}

	err = json.Unmarshal(cache, &tokenCache)
	if err != nil {
		return 0, false
	}

	return tokenCache.Price, !tokenCache.IsExpired()
}

func (c *RedisCache) SetTokenPrice(ctx context.Context, token string, price float64) {
	key := c.getTokenPriceCacheKey(token)

	tokenCache := tokenPriceCache{
		Price:     price,
		ExpiresAt: time.Now().Add(tokenPriceExpiration),
	}

	cache, err := json.Marshal(tokenCache)
	if err != nil {
		return
	}

	_ = c.client.Set(ctx, key, cache, tokenPriceTTL).Err()

	return
}

func (c *RedisCache) getTokenPriceCacheKey(token string) string {
	return fmt.Sprintf(tokenPriceCacheKeyPattern, token)
}
