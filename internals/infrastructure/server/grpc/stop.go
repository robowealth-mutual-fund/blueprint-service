package grpc

import (
	log "github.com/robowealth-mutual-fund/stdlog"
)

func (s *Server) Stop() {
	s.server.GracefulStop()
	log.Info("Server gracefully stopped")
}
