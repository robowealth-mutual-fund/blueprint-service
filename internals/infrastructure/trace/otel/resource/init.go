package resource

import (
	"fmt"

	log "github.com/robowealth-mutual-fund/stdlog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// New returns a resource describing this application.
func New(config config.Config) *resource.Resource {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			config.Trace.OtelResourceSemconv,
			semconv.ServiceNameKey.String(fmt.Sprintf("%s-%s", config.Trace.OtelServiceName, config.Environment)),
			semconv.ServiceVersionKey.String(config.Trace.OtelTracesExporter),
			attribute.String("environment", config.Environment),
		),
	)

	if err != nil {
		log.Error("New Resource: failed to create OTEL resource", err)
		panic(err)
	}

	return r
}
