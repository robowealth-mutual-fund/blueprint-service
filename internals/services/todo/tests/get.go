package tests

import (
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestGetShouldSuccess() {
	suite.mockGet().Return(&model.GetResponse{
		Id: mock.Anything,
	}, nil)
	res, err := suite.service.Get(suite.ctx, suite.mockGetRequest())
	suite.NotNil(res)
	suite.NoError(err)
}
