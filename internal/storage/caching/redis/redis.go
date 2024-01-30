package redis

import (
	"eshop-products-ms/internal/config"
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func New() (*Redis, error) {
	conf := config.Get()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Pass,
		DB:       conf.Redis.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Redis{client: client}, nil
}
