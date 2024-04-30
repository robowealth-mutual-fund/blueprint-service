package container

import "github.com/stretchr/testify/assert"

func (suite *PackageTestSuite) TestNewContainerSuccess() {
	container, err := New()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), container)
}
