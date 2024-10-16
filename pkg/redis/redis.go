package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"log"
	"time"
)

// New instantiates new Redis client
func New(ctx context.Context, configs config.RedisConfigs) (*redis.Client, error) {
	opts := redis.Options{
		Addr:         configs.Address + ":" + configs.Port,
		Password:     configs.Password,
		DB:           configs.DB,
		DialTimeout:  time.Duration(configs.DialTimeoutInSeconds) * time.Second,
		ReadTimeout:  time.Duration(configs.ReadTimeoutInSeconds) * time.Second,
		WriteTimeout: time.Duration(configs.WriteTimeoutInSeconds) * time.Second,
	}

	client := redis.NewClient(&opts)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to Redis Addr %v, Port %v Reason %v", configs.Address, configs.Port, err)
	}

	go monitorContext(ctx, client)

	return client, nil
}

func monitorContext(ctx context.Context, client *redis.Client) {
	<-ctx.Done()
	log.Println("closing redis client")
	client.Close()
}
