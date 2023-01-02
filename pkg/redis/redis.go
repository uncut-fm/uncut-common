package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
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

	go monitorContext(ctx, client)

	return client, nil
}

func monitorContext(ctx context.Context, client *redis.Client) {
	<-ctx.Done()
	log.Println("closing redis client")
	client.Close()
}
