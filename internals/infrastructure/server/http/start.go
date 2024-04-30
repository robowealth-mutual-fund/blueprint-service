package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func (s *Server) Start() {
	port := s.config.Server.HTTPPort
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.httpMux,
	}

	s.httpMux.Handle("/", s.rMux)
	s.swagger()
	signal.Notify(make(chan os.Signal, 1), os.Interrupt)
	log.Info(fmt.Sprintf("Listening and serving HTTP server on %s", port))

	gracefulStop := make(chan bool)

	go func() {
		sigint := make(chan os.Signal, 3)

		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint

		timeout := time.Duration(s.config.Server.HTTPServerTimeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Error("HTTP server shutdown: %v", err)
		}

		<-ctx.Done()

		log.Info(fmt.Sprintf("HTTP Server Timeout of %s", timeout.String()))
		log.Info("HTTP Server exiting")
		close(gracefulStop)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Error("Failed to close HTTP server: %v\n", err)
	}

	<-gracefulStop
	log.Info("HTTP server gracefully stopped")
}
