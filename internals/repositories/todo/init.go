package todo

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/oracle"
)

type Repository struct {
	Config config.Config
	oracle oracle.Interface
}

func New(config config.Config, oracle oracle.Interface) Interface {
	return &Repository{
		Config: config,
		oracle: oracle,
	}
}
