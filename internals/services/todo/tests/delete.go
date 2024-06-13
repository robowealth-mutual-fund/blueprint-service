package tests

import (
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDeleteShouldSuccess() {
	suite.mockDelete().Return(&model.DeleteResponse{
		Message: mock.Anything,
	}, nil)
	res, err := suite.service.Delete(suite.ctx, suite.mockDeleteRequest())
	suite.NotNil(res)
	suite.NoError(err)
}
