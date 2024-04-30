package grpc

import "github.com/robowealth-mutual-fund/blueprint-service/internals/config"

type Server struct {
	config config.Config
}

func New(config config.Config) *Server {
	server := &Server{
		config: config,
	}

	return server
}
