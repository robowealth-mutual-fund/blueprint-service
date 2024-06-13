package tests

func (suite *PackageTestSuite) TestListShouldSuccess() {
	suite.mockList().Return(nil, nil)
	_, err := suite.repository.List(suite.ctx)
	suite.NoError(err)
}
