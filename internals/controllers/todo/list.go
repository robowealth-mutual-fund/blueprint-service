package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
)

func (c Controller) List(ctx context.Context, request *message.ListRequest) (*message.ListResponse, error) {
	return &message.ListResponse{}, nil
}
