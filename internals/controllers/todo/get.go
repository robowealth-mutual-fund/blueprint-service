package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (c Controller) Get(ctx context.Context, request *message.GetRequest) (*message.GetResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, c.config.Trace.OtelServiceName, "Controller.Todo.Get")
	defer span.End()

	response, err := c.service.Get(ctx, &model.GetRequest{Id: request.Id})
	if err != nil {
		span.SetStatus(codes.Error, "Error Get Todo")
		span.RecordError(err)
		return nil, err
	}

	return &message.GetResponse{
		Id:        response.Id,
		TaskName:  response.TaskName,
		Status:    response.Status,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}, nil
}
