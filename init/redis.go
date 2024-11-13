package init

import (
	"context"
	"easy_mall/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Username: config.RedisUsername,
		Password: config.RedisPassword,
		DB:       config.RedisDbName,
	})
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
