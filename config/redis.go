package config

import (
	"github.com/go-redis/redis"
)

var (
	RedisCon *redis.Client
)

func NewRedisConnection() (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	return c, err
}
