package infrastructure

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Subscriber interface {
	Subscribe(context.Context, string) *redis.PubSub
}

type redisClient struct {
	client *redis.Client
}

func NewSubscriber() Subscriber {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	return &redisClient{
		client: client,
	}
}

func (r *redisClient) Subscribe(ctx context.Context, topic string) *redis.PubSub {
	return r.client.Subscribe(ctx, topic)
}
