package tests

import (
	"time"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
)

func (suite *PackageTestSuite) TestListShouldSuccess() {
	now := time.Now().Unix()
	expectedResponse := []model.ListResponse{
		{Id: "1", TaskName: "TODO 1", Status: "active", CreatedAt: now, UpdatedAt: now},
		{Id: "2", TaskName: "TODO 2", Status: "inactive", CreatedAt: now, UpdatedAt: now},
	}

	suite.mockList().Return(&expectedResponse, nil)

	res, err := suite.service.List(suite.ctx)

	suite.NotNil(res)
	suite.NoError(err)
}
