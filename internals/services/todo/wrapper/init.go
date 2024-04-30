package wrapper

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	service "github.com/robowealth-mutual-fund/blueprint-service/internals/services/todo"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperTodo"`
	Service service.Interface
	Config  config.Config
}
