package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	log "github.com/robowealth-mutual-fund/stdlog"
)

type Redis struct {
	RedisClient *redis.Client
}

func New(config config.Config) *Redis {

	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    config.Redis.MasterName,
		SentinelAddrs: strings.Split(config.Redis.SentinelHost, ","),
		Password:      config.Redis.RedisPassword,
		RouteRandomly: false,
	})

	log.Info("Connecting to Redis")
	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		panic(fmt.Errorf("error connecting to redis:: %w", err))
	}
	return &Redis{
		RedisClient: redisClient,
	}
}
