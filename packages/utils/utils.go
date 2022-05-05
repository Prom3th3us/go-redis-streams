package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

// NewRedisClient create a new instace of client redis
func NewRedisClient() (*redis.ClusterClient, error) {
	client := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{
				fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
			},
			PoolSize:     10000,
			PoolTimeout:  30 * time.Millisecond,
			DialTimeout:  1 * time.Second,
			ReadTimeout:  20 * time.Millisecond,
			WriteTimeout: 20 * time.Millisecond,
			IdleTimeout:  5 * time.Second,
		})

	_, err := client.Ping().Result()
	return client, err
}
