package container

import (
	log "github.com/robowealth-mutual-fund/stdlog"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/gateway"
	grpcServer "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/server/grpc"
	httpGateway "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/server/http"
	otelTrace "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/trace/otel"
)

func (c *Container) Start() error {
	log.Info("Start Container")

	if err := c.container.Invoke(func(
		tracer *otelTrace.Tracer,
		grpcServer *grpcServer.Server,
		httpGateway *httpGateway.Server,
		gateway *gateway.Server,
	) {
		go func() {
			grpcServer.Start()
		}()

		tracerShutdown := tracer.Configure()
		defer tracerShutdown()
		gateway.Start()

	}); err != nil {
		log.Error("Error Container Invoke", err)
		return err
	}

	return nil
}
