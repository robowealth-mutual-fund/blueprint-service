package tests

import "errors"

func (suite *PackageTestSuite) TestUpdateShouldSuccess() {
	suite.mockUpdate().Return(nil, nil)
	_, err := suite.repository.Update(suite.ctx, suite.mockUpdateRequest())
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestUpdateShouldError() {
	suite.mockUpdate().Return(nil, errors.New("todo not found"))
	res, err := suite.repository.Update(suite.ctx, suite.mockUpdateRequest())
	suite.Nil(res)
	suite.Error(err)
	suite.EqualError(err, "todo not found")
}
