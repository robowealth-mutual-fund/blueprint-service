package health

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Controller struct {
	status grpc_health_v1.HealthCheckResponse_ServingStatus
	code   codes.Code
}

func New(
	status grpc_health_v1.HealthCheckResponse_ServingStatus,
	code codes.Code,
) *Controller {
	return &Controller{
		status: status,
		code:   code,
	}
}
