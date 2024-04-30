package tests

import (
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/infrastructure/database"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/oracle"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx        context.Context
	db         *database.DB
	DB         *sql.DB
	Mock       sqlmock.Sqlmock
	repository oracle.Interface
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		suite.FailNowf("Failed to create sqlmock", "%s", err)
	}

	suite.db = &database.DB{
		Sql:        db,
		Connection: "",
	}

	suite.DB = db
	suite.Mock = mock

	suite.repository = oracle.New(suite.db)
}

func (suite *PackageTestSuite) TearDownTest() {
	suite.DB.Close()
}
