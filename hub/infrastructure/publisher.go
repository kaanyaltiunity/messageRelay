package infrastructure

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Publisher interface {
	Publish([]string, string) error
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

func (r *redisClient) Publish(topics []string, message string) error {
	topics = []string{"test"} // hard coded for now
	for v, _ := range topics {
		fmt.Println(v)
	}
	return nil
}
