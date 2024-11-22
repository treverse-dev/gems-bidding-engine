// pkg/redisclient/redisclient.go
package redisclient

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisClient{Client: client}
}

func (r *RedisClient) Ping() error {
	return r.Client.Ping(ctx).Err()
}
