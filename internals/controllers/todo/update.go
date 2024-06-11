package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
)

func (c Controller) Update(ctx context.Context, request *message.UpdateRequest) (*message.UpdateResponse, error) {
	return &message.UpdateResponse{}, nil
}
