package tests

import (
	"errors"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/robowealth-mutual-fund/blueprint-service/internals/entities"
	_ "github.com/sijms/go-ora/v2"
	"github.com/stretchr/testify/mock"
)

var (
	id          = uuid.New().String()
	now         = time.Now().Unix()
	createdAt   = now
	updatedAt   = now
	offset      = int64(0)
	limit       = int64(10)
	code        = "001"
	mockRow     = []string{"ID", "CODE", "NAME_TH", "NAME_EN", "VERSION", "CREATED_AT", "UPDATED_AT"}
	createdDate = time.Now().Unix()
	version     = "v1.0.0"
	orderBy     = "CREATED_AT"
	country     = entities.Mock{}
)

func (suite *PackageTestSuite) TestCreateSuccess() {

	suite.Mock.ExpectPrepare("^INSERT INTO MOCK \\(CODE,CREATED_AT,ID,NAME_EN,NAME_TH,UPDATED_AT,VERSION\\) VALUES \\(:1,:2,:3,:4,:5,:6,:7\\)$").
		ExpectExec().WithArgs("001", createdAt, id, "test EN 1", "test TH 1", updatedAt, "v1.0.0"). // refund amount, user id
		WillReturnResult(sqlmock.NewResult(1, 1))
	create, err := suite.repository.Create(suite.ctx, &entities.Mock{
		ID:        id,
		Code:      "001",
		NameTH:    "test TH 1",
		NameEN:    "test EN 1",
		Version:   "v1.0.0",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})

	suite.NotNil(create)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestListSuccess() {

	suite.Mock.ExpectQuery("^SELECT COUNT(.+) FROM MOCK WHERE CODE = :1").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows([]string{" COUNT(*)"}).
			AddRow(1))

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY ID DESC OFFSET 0 ROWS FETCH NEXT :2 ROWS ONLY$").
		WithArgs(code, limit).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(id, code, mock.Anything, mock.Anything, mock.Anything, time.Now().Unix(), time.Now().Unix()))

	list, err := suite.repository.List(suite.ctx, offset, limit, "ID DESC", []string{"*"}, map[string]any{"CODE": code}, &country)
	suite.NotNil(list)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestListFailed() {

	suite.Mock.ExpectQuery("^SELECT COUNT(.+) FROM MOCK WHERE CODE = :1").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows([]string{" COUNT(*)"}).
			AddRow(1))
	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY ID OFFSET 0 ROWS FETCH NEXT :2 ROWS ONLY$").
		WithArgs(code, limit).
		WillReturnError(errors.New("list failed"))

	list, err := suite.repository.List(suite.ctx, offset, limit, "ID DESC", []string{"*"}, map[string]any{"CODE": code}, &country)
	suite.Nil(list)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestListCountFailed() {

	suite.Mock.ExpectQuery("^SELECT COUNT(.+) FROM MOCK WHERE CODE = :1").
		WithArgs(code).
		WillReturnError(errors.New("count list failed"))

	list, err := suite.repository.List(suite.ctx, offset, limit, "ID DESC", []string{"*"}, map[string]any{"CODE": code}, &country)
	suite.Nil(list)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestListCountBuilderFailed() {
	list, err := suite.repository.List(suite.ctx, offset, limit, "ID DESC", []string{"*"}, errors.New("error"), &country)
	suite.Nil(list)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestUpdateSuccess() {
	suite.Mock.ExpectPrepare("^UPDATE MOCK SET CODE = :1 WHERE ID = :2$").
		ExpectExec().WithArgs(code, id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	update, err := suite.repository.Update(suite.ctx, map[string]any{"ID": id}, &entities.Mock{Code: code})

	suite.NotNil(update)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestRawSuccess() {
	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 AND VERSION = :2$").
		WithArgs(code, version).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(uuid.New().String(), code, "test1", "test1", version, createdDate, time.Now().Unix()).
			AddRow(uuid.New().String(), code, "test2", "test2", version, createdDate, time.Now().Unix()).
			AddRow(uuid.New().String(), code, "test3", "test3", version, createdDate, time.Now().Unix()))

	sql := fmt.Sprintf("%s * %s MOCK WHERE CODE = :1 AND VERSION = :2", "SELECT", "FROM")

	raw, err := suite.repository.Raw(suite.ctx, &country, sql, code, version)

	suite.NotNil(raw)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteSuccess() {
	suite.Mock.ExpectExec("^DELETE FROM MOCK WHERE ID = :1$").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repository.Delete(suite.ctx, map[string]any{"ID": id}, &country)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteFailed() {
	suite.Mock.ExpectExec("^DELETE FROM MOCK WHERE ID = :1$").
		WithArgs(id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repository.Delete(suite.ctx, map[string]any{"ID": id}, &country)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestDeleteBuilderFailed() {
	err := suite.repository.Delete(suite.ctx, errors.New("error builder"), &country)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestCountSuccess() {

	suite.Mock.ExpectQuery("^SELECT COUNT(.+) FROM MOCK WHERE CODE = :1$").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows([]string{
			"COUNT"}).
			AddRow(2))

	count, err := suite.repository.Count(suite.ctx, map[string]any{"CODE": code}, &country)

	suite.NotNil(count)
	suite.NoError(err)

}

func (suite *PackageTestSuite) TestCountFailed() {
	suite.Mock.ExpectQuery("^SELECT COUNT(.+) FROM MOCK WHERE CODE = :1$").
		WithArgs(code).
		WillReturnError(errors.New("error count on list"))

	count, err := suite.repository.Count(suite.ctx, map[string]any{"CODE": code}, &country)

	suite.NotNil(count)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestCountBuilderFailed() {
	count, err := suite.repository.Count(suite.ctx, errors.New("error builder"), &country)

	suite.NotNil(count)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestFirstSuccess() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT ASC FETCH FIRST 1 ROW ONLY$").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(uuid.New().String(), code, "test", "test", "v3.0.0", createdDate, time.Now().Unix()))

	err := suite.repository.First(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestFirstFailed() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT ASC FETCH FIRST 1 ROW ONLY$").
		WithArgs(code).
		WillReturnError(errors.New("error first"))

	err := suite.repository.First(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestLastSuccess() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT DESC FETCH FIRST 1 ROW ONLY$").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(uuid.New().String(), code, "test", "test", "v3.0.0", createdDate, time.Now().Unix()))

	err := suite.repository.Last(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestLastFailed() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT DESC FETCH FIRST 1 ROW ONLY$").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(uuid.New().String(), code, "test", "test", "v3.0.0", createdDate, time.Now().Unix()))

	err := suite.repository.Last(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestFindSuccess() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT$").
		WithArgs(code).
		WillReturnRows(sqlmock.NewRows(mockRow).
			AddRow(uuid.New().String(), code, "test1", "test1", "v1.0.0", createdDate, time.Now().Unix()).
			AddRow(uuid.New().String(), code, "test2", "test2", "v2.0.0", createdDate, time.Now().Unix()).
			AddRow(uuid.New().String(), code, "test3", "test3", "v3.0.0", createdDate, time.Now().Unix()))

	find, err := suite.repository.Find(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.NotNil(find)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestFindFailed() {

	suite.Mock.ExpectQuery("^SELECT (.+) FROM MOCK WHERE CODE = :1 ORDER BY CREATED_AT$").
		WithArgs(code).
		WillReturnError(errors.New("error find"))

	find, err := suite.repository.Find(suite.ctx, orderBy, []string{"*"}, map[string]any{"CODE": code}, &country)

	suite.NotNil(find)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestFindBuilderFailed() {

	find, err := suite.repository.Find(suite.ctx, orderBy, []string{"*"}, errors.New("error builder"), &country)

	suite.NotNil(find)
	suite.Error(err)
}
