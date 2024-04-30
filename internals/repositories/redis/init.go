package redis

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/redis"
)

type Repository struct {
	redis *redis.Redis
}

func New(redis *redis.Redis) Interface {
	return &Repository{
		redis: redis,
	}
}
