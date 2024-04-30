package gateway

import (
	grpcGateway "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/gateway/grpc"
	httpGateway "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/server/http"
)

type Server struct {
	grpcServer *grpcGateway.Server
	httpServer *httpGateway.Server
}

func New(grpcServer *grpcGateway.Server, httpServer *httpGateway.Server) *Server {
	return &Server{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}
