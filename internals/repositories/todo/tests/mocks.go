package tests

import (
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

// Create todo
func (suite *PackageTestSuite) mockCreateRequest() *model.CreateRequest {
	return &model.CreateRequest{
		TaskName: mock.Anything,
		Status:   mock.Anything,
	}
}

func (suite *PackageTestSuite) mockCreate() *mock.Call {
	return suite.todoRepository.On("Create", mock.Anything, mock.Anything).Once()
}

// Delete todo
func (suite *PackageTestSuite) mockDeleteRequest() *model.DeleteRequest {
	return &model.DeleteRequest{
		Id: mock.Anything,
	}
}

func (suite *PackageTestSuite) mockDelete() *mock.Call {
	return suite.todoRepository.On("Delete", mock.Anything, mock.Anything).Once()
}

// Get todo
func (suite *PackageTestSuite) mockGetRequest() *model.GetRequest {
	return &model.GetRequest{
		Id: mock.Anything,
	}
}

func (suite *PackageTestSuite) mockGet() *mock.Call {
	return suite.todoRepository.On("Get", mock.Anything, mock.Anything).Once()
}

// List todo
func (suite *PackageTestSuite) mockList() *mock.Call {
	return suite.todoRepository.On("List", mock.Anything, mock.Anything).Once()
}

// Update todo
func (suite *PackageTestSuite) mockUpdateRequest() *model.UpdateRequest {
	return &model.UpdateRequest{
		Id:       mock.Anything,
		TaskName: mock.Anything,
		Status:   mock.Anything,
	}
}

func (suite *PackageTestSuite) mockUpdate() *mock.Call {
	return suite.todoRepository.On("Update", mock.Anything, mock.Anything).Once()
}
