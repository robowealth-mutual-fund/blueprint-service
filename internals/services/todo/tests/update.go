package tests

import (
	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdateShouldSuccess() {
	suite.mockUpdate().Return(&model.UpdateResponse{Id: mock.Anything}, nil)
	res, err := suite.service.Update(suite.ctx, suite.mockUpdateRequest())
	suite.NotNil(res)
	suite.NoError(err)
}
