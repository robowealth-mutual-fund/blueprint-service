package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (c Controller) Delete(ctx context.Context, request *message.DeleteRequest) (*message.DeleteResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, c.config.Trace.OtelServiceName, "Controller.Todo.Delete")
	defer span.End()

	response, err := c.service.Delete(ctx, &model.DeleteRequest{
		Id: request.Id,
	})

	if err != nil {
		span.SetStatus(codes.Error, "Error Delete Todo")
		span.RecordError(err)
		return nil, err
	}

	return &message.DeleteResponse{
		Message: response.Message,
	}, nil
}
