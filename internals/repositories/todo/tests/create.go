package tests

import (
	"errors"
)

func (suite *PackageTestSuite) TestCreateShouldSuccess() {
	suite.mockCreate().Return(nil, nil)
	_, err := suite.repository.Create(suite.ctx, suite.mockCreateRequest())
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCreateShouldError() {
	suite.mockCreate().Return(nil, errors.New("create todo error"))
	res, err := suite.repository.Create(suite.ctx, suite.mockCreateRequest())
	suite.Nil(res)
	suite.Error(err)
}
