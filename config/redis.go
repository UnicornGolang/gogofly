package config

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	redisClient *redis.Client
	duration    = 30 * 24 * 60 * 60 * time.Second
)

type RedisClient struct {}

func InitRedis() (*RedisClient, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "redis123!@#",
		DB:       1,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

// 对于第三方的组件，是否需要对 API 进行封装，
// 取决于管理的要求
func (rc *RedisClient) Set(key string, value any) error {
	return redisClient.Set(context.Background(), key, value, duration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return redisClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Del(key ...string) (any, error) {
	return redisClient.Del(context.Background(), key...).Result()
}
