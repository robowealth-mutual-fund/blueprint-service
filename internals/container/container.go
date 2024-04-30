package container

import (
	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func New() (*Container, error) {
	container := &Container{
		container: dig.New(),
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	appConfig := config.New()
	container.NewLogger(appConfig)

	return container, nil
}
