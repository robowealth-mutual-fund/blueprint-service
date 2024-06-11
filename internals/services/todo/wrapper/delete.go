package wrapper

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	traceOtel "go.opentelemetry.io/otel/trace"
)

func (w Wrapper) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, w.Config.Trace.OtelServiceName, "Service.Todo.Delete")
	defer span.End()

	span.AddEvent("delete-todo-request", traceOtel.WithAttributes(
		attribute.String("id", request.Id),
	))

	res, err := w.Service.Delete(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Create Todo")
		span.RecordError(err)
		return nil, err
	}
	return res, nil
}
