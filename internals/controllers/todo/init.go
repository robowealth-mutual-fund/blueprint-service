package todo

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	todoService "github.com/robowealth-mutual-fund/blueprint-service/internals/services/todo/wrapper"
)

type Controller struct {
	config  config.Config
	service todoService.Wrapper
}

func New(config config.Config,
	service todoService.Wrapper,
) *Controller {
	return &Controller{
		config:  config,
		service: service,
	}
}
