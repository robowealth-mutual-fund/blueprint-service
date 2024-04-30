package redis

import (
	"context"
)

//go:generate mockery --name=Interface
type Interface interface {
	Set(ctx context.Context, key string, value any, expiration string) error
	Get(ctx context.Context, key string, entity interface{}) error
	Delete(ctx context.Context, key string) error
}
