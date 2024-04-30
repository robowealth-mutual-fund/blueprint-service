package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
)

func (r *Repository) Set(ctx context.Context, key string, value any, expiration string) error {
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		return err
	}

	if err = r.redis.RedisClient.Set(ctx, key, value, duration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, key string, entity interface{}) error {

	result, err := r.redis.RedisClient.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return errors.New("key does not exist")
	} else if err != nil {
		return err
	}

	err = utils.Unmarshal([]byte(result), &entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, key string) error {

	if err := r.redis.RedisClient.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
