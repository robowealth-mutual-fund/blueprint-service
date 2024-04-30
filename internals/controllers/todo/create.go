package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (c Controller) Create(ctx context.Context, request *message.CreateRequest) (*message.CreateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, c.config.Trace.OtelServiceName, "Controller.Todo.Create")
	defer span.End()

	response, err := c.service.Create(ctx, &model.CreateRequest{
		TaskName: request.GetTaskName(),
		Status:   request.GetStatus(),
	})

	if err != nil {
		span.SetStatus(codes.Error, "Error Create Todo")
		span.RecordError(err)
		return nil, err
	}

	return &message.CreateResponse{
		TaskName:  response.TaskName,
		Status:    response.Status,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}, nil
}
