package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func (s *Server) Start() {
	s.Configure()
	go func() {
		port := s.config.Server.GRPCPort
		listen, err := net.Listen("tcp", ":"+port)
		if err != nil {
			panic(err)
		}

		log.Println("Listening and serving GRPC server on", port)
		if err := s.server.Serve(listen); err != nil {
			panic(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
	log.Println("GRPC server gracefully stopped")
	s.Stop()
	s.controller.database.Stop()
}
