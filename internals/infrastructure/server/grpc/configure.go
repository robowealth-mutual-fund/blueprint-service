package grpc

import (
	todoV1 "github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo"
	healthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

// Configure ...
func (s *Server) Configure() {
	healthV1.RegisterHealthServer(s.server, s.controller.healthController)
	todoV1.RegisterTodoServiceServer(s.server, s.controller.todoController)
}
