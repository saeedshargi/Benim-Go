package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Ping(context.Context) error
	Close() error
	Set(context.Context, string, interface{}) error
	Get(context.Context, string) (string, error)
}

type redisClient struct {
	cl *redis.Client
}

func NewClient(connection string, pass string, _ string) Client {

	time.Local = time.UTC
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: pass, // no password set
		DB:       0,    // use default DB
	})

	return &redisClient{cl: rdb}

}

func (rc *redisClient) Ping(ctx context.Context) error {
	_, err := rc.cl.Ping(ctx).Result()
	return err
}

func (rc *redisClient) Close() error {
	return rc.cl.Close()
}

func (rc *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	return rc.cl.Set(ctx, key, value, 0).Err()
}

func (rc *redisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.cl.Get(ctx, key).Result()
}
