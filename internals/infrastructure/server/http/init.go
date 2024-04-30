package http

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	healthzClient "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/healthz"
)

type Server struct {
	config  config.Config
	rMux    *runtime.ServeMux
	httpMux *http.ServeMux
}

func New(config config.Config) *Server {
	client := &healthzClient.Controller{}
	server := &Server{
		config:  config,
		rMux:    runtime.NewServeMux(runtime.WithHealthzEndpoint(client)),
		httpMux: http.NewServeMux(),
	}

	return server
}

func (s *Server) RMux() *runtime.ServeMux {
	return s.rMux
}
