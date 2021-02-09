package main

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a cache of Redis
type RedisClient struct {
	client *redis.Client
}

// GetClient is the best method to create clients.
func (c *RedisClient) GetClient() *redis.Client {
	if c.client == nil {
		c.client = redis.NewClient(&redis.Options{
			Addr:        "localhost:6379",
			Password:    "",
			DB:          0,
			DialTimeout: 3 * time.Second,
			PoolSize:    300,
		})
	}

	return c.client
}
