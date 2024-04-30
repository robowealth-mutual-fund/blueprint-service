package utils

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartSpanFromContext(
	ctx context.Context,
	serviceName, spanName string,
	opts ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	return otel.Tracer(serviceName).Start(ctx, spanName, opts...)
}
