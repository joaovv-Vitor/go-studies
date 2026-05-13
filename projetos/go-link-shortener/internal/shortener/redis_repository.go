package shortener

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) CacheRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrURLNotFound
	}
	if err != nil {
		return "", fmt.Errorf("erro ao buscar no cache: %w", err)
	}
	return val, nil
}
