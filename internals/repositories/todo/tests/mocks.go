package tests

import (
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) mockCreateRequest() *model.CreateRequest {
	return &model.CreateRequest{
		TaskName: mock.Anything,
		Status:   mock.Anything,
	}
}

func (suite *PackageTestSuite) mockCreate() *mock.Call {
	return suite.todoRepository.On("Create", mock.Anything, mock.Anything).Once()
}
