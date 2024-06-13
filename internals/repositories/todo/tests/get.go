package tests

import (
	"errors"
)

func (suite *PackageTestSuite) TestGetTodoShouldSuccess() {
	suite.mockGet().Return(nil, nil)
	_, err := suite.repository.Get(suite.ctx, suite.mockGetRequest())
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestGetTodoShouldError() {
	suite.mockGet().Return(nil, errors.New("todo not found"))

	res, err := suite.repository.Get(suite.ctx, suite.mockGetRequest())

	suite.Nil(res)
	suite.Error(err)
	suite.EqualError(err, "todo not found")
}
