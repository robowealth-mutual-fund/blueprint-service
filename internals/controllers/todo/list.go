package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (c Controller) List(ctx context.Context, _ *message.ListRequest) (*message.ListResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, c.config.Trace.OtelServiceName, "Controller.Todo.List")
	defer span.End()

	todos, err := c.service.List(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "Error List Todo")
		span.RecordError(err)
		return nil, err
	}

	response := &message.ListResponse{}

	response.Count = int64(len(*todos))

	for _, todo := range *todos {
		dataItem := &message.ListResponse_TodoItem{
			Id:        todo.Id,
			TaskName:  todo.TaskName,
			Status:    todo.Status,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		}

		response.Data = append(response.Data, dataItem)
	}

	return response, nil
}
