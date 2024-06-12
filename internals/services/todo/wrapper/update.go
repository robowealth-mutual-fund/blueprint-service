package wrapper

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	traceOtel "go.opentelemetry.io/otel/trace"
)

func (w Wrapper) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, w.Config.Trace.OtelServiceName, "Service.Todo.Update")
	defer span.End()

	span.AddEvent("delete-todo-request", traceOtel.WithAttributes(
		attribute.String("id", request.Id),
		attribute.String("taskName", request.TaskName),
		attribute.String("status", request.Status),
	))

	res, err := w.Service.Update(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Update Todo")
		span.RecordError(err)
		return nil, err
	}
	return res, nil
}
