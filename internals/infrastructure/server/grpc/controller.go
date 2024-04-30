package grpc

import (
	healthCtrl "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/health"
	todoCtrl "github.com/robowealth-mutual-fund/blueprint-service/internals/controllers/todo"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/database"
)

type Controller struct {
	healthController *healthCtrl.Controller
	database         *database.DB
	todoController   *todoCtrl.Controller
}

func NewController(
	healthController *healthCtrl.Controller,
	database *database.DB,
	todoController *todoCtrl.Controller,
) *Controller {
	return &Controller{
		healthController: healthController,
		database:         database,
		todoController:   todoController,
	}
}
