package wrapper

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/utils"
	"go.opentelemetry.io/otel/codes"
)

func (w Wrapper) List(ctx context.Context) (*[]model.ListResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, w.Config.Trace.OtelServiceName, "Service.Todo.List")
	defer span.End()

	res, err := w.Service.List(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "Error List Todo")
		span.RecordError(err)
		return nil, err
	}
	return res, nil
}
