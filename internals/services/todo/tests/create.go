package tests

import (
	"time"

	model "github.com/robowealth-mutual-fund/blueprint-service/internals/models/todo"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreateShouldSuccess() {
	suite.mockCreate().Return(&model.CreateResponse{
		TaskName:  mock.Anything,
		Status:    mock.Anything,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}, nil)
	res, err := suite.service.Create(suite.ctx, suite.mockCreateRequest())
	suite.NotNil(res)
	suite.NoError(err)
}
