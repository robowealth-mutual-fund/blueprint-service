package container

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
}

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}
