package config

import (
	"github.com/go-redis/redis"
)

var (
	RedisCon *redis.Client
)

func NewRedisConnection() (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     getEnv(REDIS_HOST, "redis-server") + ":" + getEnv(REDIS_PORT, "6379"),
		Password: getEnv(REDIS_PASSWORD, ""), // no password set
		DB:       0,                          // use default DB
	})

	_, err := c.Ping().Result()
	return c, err
}
