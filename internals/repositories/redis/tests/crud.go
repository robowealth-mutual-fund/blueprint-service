package tests

import (
	"errors"
	"time"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestSetSuccess() {

	suite.mocks.ExpectSet("key", "value", 5*time.Minute).SetVal("OK")
	err := suite.repository.Set(suite.ctx, "key", "value", "5m")
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestSetParseDurationFailed() {

	suite.mocks.ExpectSet("key", "value", 1).SetErr(errors.New("set error"))
	err := suite.repository.Set(suite.ctx, "key", "value", mock.Anything)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestSetFailed() {

	suite.mocks.ExpectSet("key", "value", 5*time.Minute).SetErr(errors.New("set error"))
	err := suite.repository.Set(suite.ctx, "key", "value", "5m")
	suite.Error(err)
}

func (suite *PackageTestSuite) TestGetSuccess() {
	type MyStruct struct {
		Value string `json:"value"`
	}

	var myStruct MyStruct
	jsonData := `{"value":"OK"}`
	suite.mocks.ExpectGet("key").SetVal(jsonData)
	err := suite.repository.Get(suite.ctx, "key", &myStruct)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestGetUnmarshalFailed() {

	suite.mocks.ExpectGet("key").SetVal("OK")
	err := suite.repository.Get(suite.ctx, "key", mock.Anything)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestGetFailed() {

	suite.mocks.ExpectGet("key").SetErr(errors.New("get error"))
	err := suite.repository.Get(suite.ctx, "key", mock.Anything)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestDeleteSuccess() {

	suite.mocks.ExpectDel("key").SetVal(0)
	err := suite.repository.Delete(suite.ctx, "key")
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteFailed() {

	suite.mocks.ExpectDel("key").SetErr(errors.New("delete error"))
	err := suite.repository.Delete(suite.ctx, "key")
	suite.Error(err)
}
