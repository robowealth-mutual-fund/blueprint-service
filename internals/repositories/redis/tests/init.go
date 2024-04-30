package tests

import (
	"context"

	redisMock "github.com/go-redis/redismock/v9"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/redis"
	redisRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/redis"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx        context.Context
	redis      *redis.Redis
	repository redisRepo.Interface
	mocks      redisMock.ClientMock
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	db, mock := redisMock.NewClientMock()

	suite.redis = &redis.Redis{RedisClient: db}
	suite.mocks = mock
	suite.repository = redisRepo.New(suite.redis)
}
