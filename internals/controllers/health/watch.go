package health

import "google.golang.org/grpc/health/grpc_health_v1"

func (c Controller) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return nil
}
