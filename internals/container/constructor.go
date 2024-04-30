package container

import (
	"github.com/go-resty/resty/v2"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	healthController "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/health"
	healthzController "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/healthz"
	todoController "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/client/http"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/database"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/gateway"
	grpcGateway "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/gateway/grpc"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/redis"
	grpcServer "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/server/grpc"
	httpGateway "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/server/http"
	otelTrace "github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/trace/otel"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/trace/otel/exporter"
	oracleRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/oracle"
	redisRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/redis"
	restRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/rest"
	todoRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/todo"
	todoSvc "github.com/robowealth-mutual-fund/blueprint-service/internals/services/todo"
	"github.com/robowealth-mutual-fund/shared-utility/validator"
)

var constructors = []any{
	config.New,
	database.New,
	gateway.New,
	validator.NewCustomValidator,
	grpcGateway.New,
	httpGateway.New,
	grpcServer.New,
	resty.New,
	redis.New,
	http.NewHttpClient,
	exporter.New,
	otelTrace.New,

	// Controllers
	grpcServer.NewController,
	healthController.New,
	healthzController.New,
	todoController.New,

	// Services
	todoSvc.New,

	// Repositories
	restRepo.New,
	oracleRepo.New,
	redisRepo.New,
	todoRepo.New,
}
