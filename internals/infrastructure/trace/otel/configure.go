package otel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	otelResource "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/trace/otel/resource"
)

func (t *Tracer) Configure() func() {
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(t.exporter.SpanExporter),
		sdktrace.WithResource(otelResource.New(t.config)),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	t.traceProvider = tracerProvider

	return t.shutdown()
}
