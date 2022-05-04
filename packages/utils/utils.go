package utils

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

//NewRedisClient create a new instace of client redis
func NewRedisClient() (*redis.ClusterClient, error) {

	client := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:     []string{
				fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
			},
		})

	_, err := client.Ping().Result()
	return client, err
}
