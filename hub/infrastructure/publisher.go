package infrastructure

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
)

type Publisher interface {
	Publish(echo.Context, string, string) error
}

type redisClient struct {
	client *redis.Client
}

func NewPublisher() Publisher {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &redisClient{
		client: client,
	}
}

func (r *redisClient) Publish(ctx echo.Context, topic string, message string) error {
	err := r.client.Publish(ctx.Request().Context(), topic, message).Err()
	if err != nil {
		return err
	}
	return nil
}
