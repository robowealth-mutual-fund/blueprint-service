package redis_test

import (
	"testing"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/redis/tests"
	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(tests.PackageTestSuite))
}
