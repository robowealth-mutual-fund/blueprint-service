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
	return suite.todoRepository.On("Create", suite.ctx, suite.mockCreateRequest()).Once()
}

func (suite *PackageTestSuite) mockDeleteRequest() *model.DeleteRequest {
	return &model.DeleteRequest{Id: mock.Anything}
}

func (suite *PackageTestSuite) mockDelete() *mock.Call {
	return suite.todoRepository.On("Delete", suite.ctx, suite.mockDeleteRequest()).Once()
}

func (suite *PackageTestSuite) mockGetRequest() *model.GetRequest {
	return &model.GetRequest{Id: mock.Anything}
}

func (suite *PackageTestSuite) mockGet() *mock.Call {
	return suite.todoRepository.On("Get", suite.ctx, suite.mockGetRequest()).Once()
}

func (suite *PackageTestSuite) mockList() *mock.Call {
	return suite.todoRepository.On("List", suite.ctx).Once()
}

func (suite *PackageTestSuite) mockUpdateRequest() *model.UpdateRequest {
	return &model.UpdateRequest{
		Id:       mock.Anything,
		TaskName: mock.Anything,
		Status:   mock.Anything,
	}
}

func (suite *PackageTestSuite) mockUpdate() *mock.Call {
	return suite.todoRepository.On("Update", suite.ctx, suite.mockUpdateRequest()).Once()
}
