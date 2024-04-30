package exporter

import (
	"context"

	log "github.com/robowealth-mutual-fund/stdlog"
	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Exporter struct {
	SpanExporter trace.SpanExporter
}

func New() *Exporter {
	spanExporter, err := autoexport.NewSpanExporter(context.Background())

	if err != nil {
		log.Error("New Exporter: failed to create exporter", err)
		panic(err)
	}

	return &Exporter{
		SpanExporter: spanExporter,
	}
}
