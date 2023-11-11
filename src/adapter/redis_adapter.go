package adapter

import (
	"fmt"
	"mcx002/filestoreService/src"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	NewRedis(appConfig *src.AppConfig) *redis.Client
}

type RedisAdapter struct{}

func NewRedisAdapter() *RedisAdapter {
	return &RedisAdapter{}
}

func (RedisAdapter) NewRedis(appConfig *src.AppConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", appConfig.RedisHost, appConfig.RedisPort),
		Password: appConfig.RedisPassword, // no password set
		DB:       0,                       // use default DB
	})
}
