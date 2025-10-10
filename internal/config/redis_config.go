package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisConfig(v *viper.Viper) *RedisConfig {
	return &RedisConfig{
		Host:     v.GetString("REDIS_HOST"),
		Port:     v.GetString("REDIS_PORT"),
		Password: v.GetString("REDIS_PASSWORD"),
		DB:       v.GetInt("REDIS_DB"),
	}
}

func (rc *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}

func (rc *RedisConfig) NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.GetAddr(),
		Password: rc.Password,
		DB:       rc.DB,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return client
}
