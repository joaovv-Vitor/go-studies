package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient cria uma nova conexão com o Redis

func NewRedisClient(ctx context.Context, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao redis: %w", err)
	}

	client := redis.NewClient(opts)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao redis: %w", err)
	}
	return client, nil

}
