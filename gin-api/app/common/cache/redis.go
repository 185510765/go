package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// func InitRedis() {
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	ping, err := RedisClient.Ping().Result()
	if err == redis.Nil {
		fmt.Print("Redis异常")
	} else if err != nil {
		fmt.Print("失败:", err)
	} else {
		fmt.Print(ping)
	}
}
