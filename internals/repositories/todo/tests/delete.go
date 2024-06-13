package tests

import "errors"

func (suite *PackageTestSuite) TestDeleteShouldSuccess() {
	suite.mockDelete().Return(nil, nil)
	_, err := suite.repository.Delete(suite.ctx, suite.mockDeleteRequest())
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteShouldError() {
	suite.mockDelete().Return(nil, errors.New("delete todo error"))
	res, err := suite.repository.Delete(suite.ctx, suite.mockDeleteRequest())
	suite.Nil(res)
	suite.Error(err)
}
