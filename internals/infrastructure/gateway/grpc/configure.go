package grpc

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	todoV1 "github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) Configure(rMux *runtime.ServeMux) {
	const ip = "0.0.0.0:%s"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	_ = todoV1.RegisterTodoServiceHandlerFromEndpoint(
		context.Background(),
		rMux,
		fmt.Sprintf(ip, s.config.Server.GRPCPort),
		opts,
	)
}
