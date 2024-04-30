package gateway

func (s *Server) Start() {
	s.grpcServer.Configure(s.httpServer.RMux())
	s.httpServer.Start()
}
