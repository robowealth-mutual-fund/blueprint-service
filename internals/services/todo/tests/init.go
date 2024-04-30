package tests

import (
	"context"

	todoRepository "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/todo/mocks"
	service "github.com/robowealth-mutual-fund/blueprint-service/internals/services/todo"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx            context.Context
	todoRepository *todoRepository.Interface
	service        service.Interface
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.todoRepository = &todoRepository.Interface{}
	suite.service = service.New(suite.todoRepository)
}
