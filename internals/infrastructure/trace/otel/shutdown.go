package otel

import (
	"context"
	"time"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func (t *Tracer) shutdown() func() {
	return func() {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		// Cleanly shutdown and flush telemetry when the application exits.
		defer func(ctx context.Context) {
			// Do not make the application hang when it is shutdown.
			timeout := time.Duration(t.config.Trace.ShutdownTimeout) * time.Second
			ctx, cancel := context.WithTimeout(ctx, timeout)

			defer cancel()

			if err := t.traceProvider.Shutdown(ctx); err != nil {
				log.Error("TraceProvider: Failed to shutdown", err)
				panic(err)
			}

			log.Info("Shutdown shutdown tracer successfully")
		}(ctx)
	}
}
