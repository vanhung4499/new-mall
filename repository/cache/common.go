package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"new-mall/config"
)

// RedisClient Redis cache client singleton
var RedisClient *redis.Client
var RedisContext = context.Background()

// InitCache initializes the redis link in the middleware
func InitCache() {
	rConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Username: rConfig.RedisUsername,
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		log.Info(err)
		panic(err)
	}
	RedisClient = client
}
