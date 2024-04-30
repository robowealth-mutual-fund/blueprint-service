package tests

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/config"
	todoRepository "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/todo/mocks"
	repository "github.com/robowealth-mutual-fund/blueprint-service/internals/services/todo"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx            context.Context
	config         config.Config
	todoRepository *todoRepository.Interface
	repository     repository.Interface
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.config = config.Config{}
	suite.todoRepository = &todoRepository.Interface{}
	suite.repository = repository.New(suite.todoRepository)
}
