package rediscli

import (
	"context"
	"fliqt/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

const (
	//host = "0.0.0.0"
	host     = "docker.for.mac.localhost"
	port     = "6379"
	password = ""
	db       = 0
)

func New() *RedisClient {
	cli := &RedisClient{
		config: config.New(),
	}
	cli.connect()
	return cli
}

type RedisClient struct {
	redis.UniversalClient
	config *config.Config
}

func (r *RedisClient) connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.config.Redis.Host, r.config.Redis.Port),
		Password: r.config.Redis.Password,
		DB:       r.config.Redis.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("Failed to connect to Redis: %v", err)
	}

	r.UniversalClient = client
}
