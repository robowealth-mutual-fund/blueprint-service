package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (c Controller) Update(ctx context.Context, request *message.UpdateRequest) (*message.UpdateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, c.config.Trace.OtelServiceName, "Controller.Todo.Update")
	defer span.End()

	response, err := c.service.Update(ctx, &model.UpdateRequest{
		Id:       request.Id,
		TaskName: request.TaskName,
		Status:   request.Status,
	})
	if err != nil {
		span.SetStatus(codes.Error, "Error Update Todo")
		span.RecordError(err)
		return nil, err
	}

	return &message.UpdateResponse{
		Id: response.Id,
	}, nil
}
