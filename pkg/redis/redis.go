package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// New instantiates new Redis client
func New(ctx context.Context, addr string, port string) (*redis.Client, error) {
	opts := redis.Options{
		Addr: addr + ":" + port,
	}

	client := redis.NewClient(&opts)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to Redis Addr %v, Port %v Reason %v", addr, port, err)
	}

	return client, nil
}
