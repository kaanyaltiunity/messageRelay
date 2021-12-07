package infrastructure

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
)

type Cache interface {
	Register(echo.Context, string) error
	GetAll(echo.Context, func([]string) []string) ([]string, error)
}

type redisClientCache struct {
	client *redis.Client
}

func NewCache() Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
	return &redisClientCache{
		client: client,
	}
}

func (r *redisClientCache) Register(ctx echo.Context, uuid string) error {
	err := r.client.Set(ctx.Request().Context(), uuid, true, 0).Err()
	return err
}

func (r *redisClientCache) GetAll(ctx echo.Context, filter func([]string) []string) ([]string, error) {
	userIds := r.client.Keys(ctx.Request().Context(), "*")
	err := userIds.Err()
	if err != nil {
		return nil, err
	}
	return filter(userIds.Val()), nil
}
