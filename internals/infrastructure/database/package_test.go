//go:build integration
// +build integration

package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

const (
	testDBName = "oa_master_data_test"
)

type (
	PackageTestSuite struct {
		suite.Suite
		dbConn     *pgx.Conn
		connString string
	}

	testTables struct {
		SchemaName  string
		TableName   string
		TableOwner  string
		TableSpace  *string
		HasIndexes  bool
		HasRules    bool
		HasTriggers bool
		RowSecurity bool
	}
)

var (
	expectedTables = map[string]string{
		"occupations":           "occupations",
		"business_types":        "business_types",
		"monthly_income_levels": "monthly_income_levels",
		"countries":             "countries",
		"provinces":             "provinces",
		"districts":             "districts",
		"sub_districts":         "sub_districts",
	}
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.connString = "postgres://postgres:postgres@localhost:5433"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5433")
	if err != nil {
		suite.T().Fatal(err)
	}

	rows, err := conn.Query(context.Background(),
		fmt.Sprintf(`select datname from pg_database where datname = '%s' LIMIT 1`, testDBName))
	if err != nil {
		suite.T().Fatal(err)
	}

	if !rows.Next() {
		_, err = conn.Exec(context.Background(),
			fmt.Sprintf(`CREATE DATABASE %s`, testDBName))
		if err != nil {
			suite.T().Fatal(err)
		}
	}

	conn.Close(context.Background())

	suite.connString = fmt.Sprintf("postgres://postgres:postgres@localhost:5433/%s", testDBName)
	suite.dbConn, err = pgx.Connect(context.Background(), suite.connString)
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *PackageTestSuite) TearDownSuite() {
	suite.dbConn.Close(context.Background())

	conn, err := pgx.Connect(context.Background(), suite.connString)
	if err != nil {
		suite.T().Fatal(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `DROP SCHEMA public CASCADE`)
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `CREATE SCHEMA public`)
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `GRANT ALL ON SCHEMA public TO postgres`)
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `GRANT ALL ON SCHEMA public TO public`)
	if err != nil {
		suite.T().Fatal(err)
	}
}
