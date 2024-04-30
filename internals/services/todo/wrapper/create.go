package wrapper

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	traceOtel "go.opentelemetry.io/otel/trace"
)

func (w Wrapper) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, w.Config.Trace.OtelServiceName, "Service.Todo.Create")
	defer span.End()

	span.AddEvent("create-todo-request", traceOtel.WithAttributes(
		attribute.String("taskName", request.TaskName),
		attribute.String("status", request.Status),
	))

	res, err := w.Service.Create(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Create Todo")
		span.RecordError(err)
		return nil, err
	}
	return res, nil
}
