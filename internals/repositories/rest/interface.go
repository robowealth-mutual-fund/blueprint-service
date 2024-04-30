package rest

import "context"

//go:generate mockery --name=Interface
type Interface interface {
	RestyClientGet(ctx context.Context, host, path string) ([]byte, error)
	RestyClientPost(ctx context.Context, host, path string, request interface{}) ([]byte, error)
	ClientPostWithRetries(ctx context.Context, host, path string, headers map[string]string, request, entity, errorEntity interface{}) error
}
