package grpc

import (
	"context"
	"time"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/keepalive"

	validatorUtils "github.com/robowealth-mutual-fund/shared-utility/validator"
	switchLanguageUtil "github.com/robowealth-mutual-fund/switch-language"
	"google.golang.org/grpc"
)

type Server struct {
	config     config.Config
	server     *grpc.Server
	controller *Controller
}

func New(config config.Config, controller *Controller, validator *validatorUtils.CustomValidator) *Server {

	option := grpc.ChainUnaryInterceptor(
		unaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
		switchLanguageUtil.UnaryServerInterceptor(),
	)

	server := &Server{
		server: grpc.NewServer(option,
			grpc.WriteBufferSize(config.Server.GRPCWriteBufferSize),
			grpc.ReadBufferSize(config.Server.GRPCReadBufferSize),
			grpc.MaxConcurrentStreams(config.Server.GRPCMaxConcurrentStreams),
			grpc.MaxRecvMsgSize(10*10e6),
			grpc.MaxSendMsgSize(10*10e6),
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionAge:      time.Second * 30,
				MaxConnectionAgeGrace: time.Second * 10,
			}),
		),
		config:     config,
		controller: controller,
	}
	return server
}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
