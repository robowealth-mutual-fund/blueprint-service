package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
)

func (c Controller) Get(ctx context.Context, request *message.GetRequest) (*message.GetResponse, error) {
	return &message.GetResponse{}, nil
}
