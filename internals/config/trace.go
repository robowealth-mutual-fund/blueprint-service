package config

type Trace struct {
	OtelServiceName     string `env:"OTEL_SERVICE_NAME,expand" envDefault:"123-${APP_NAME}"`
	OtelTracesExporter  string `env:"OTEL_TRACES_EXPORTER" envDefault:"otlp"`
	OtelResourceSemconv string `env:"OTEL_RESOURCE_SEMCONV" envDefault:"https://opentelemetry.io/schemas/1.21.0"`
	ShutdownTimeout     int    `env:"TRACE_SHUTDOWN_TIMEOUT" envDefault:"10"`
}
