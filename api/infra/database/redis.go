package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(ctx context.Context, connectionString string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: connectionString,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
