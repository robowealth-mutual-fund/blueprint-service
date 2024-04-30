package otel

import (
	traceSdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/trace/otel/exporter"
)

type Tracer struct {
	config        config.Config
	traceProvider *traceSdk.TracerProvider
	exporter      *exporter.Exporter
}

func New(config config.Config, exporter *exporter.Exporter) *Tracer {
	return &Tracer{
		config:   config,
		exporter: exporter,
	}
}
