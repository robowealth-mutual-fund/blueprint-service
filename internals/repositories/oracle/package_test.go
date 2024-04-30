package oracle_test

import (
	"testing"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/oracle/tests"
	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(tests.PackageTestSuite))
}
