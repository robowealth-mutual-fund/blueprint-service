package oracle

import (
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/copier"
)

func GetTableName(entity interface{}) string {
	valueOfEntity := reflect.ValueOf(entity)

	if reflect.TypeOf(entity).Elem().Kind() == reflect.Slice {
		elemType := reflect.TypeOf(entity).Elem().Elem()
		if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
			valueOfEntity = reflect.New(elemType.Elem())
		}
	}

	if method := valueOfEntity.MethodByName("TableName"); method.IsValid() {
		valList := method.Call([]reflect.Value{})
		// has return and the first val is string
		if len(valList) > 0 && valList[0].Kind() == reflect.String {
			return valList[0].String()
		}
	}

	return ""
}

func GetInsertBuilder(entity interface{}) (string, []interface{}, error) {
	valMap := make(map[string]interface{}, 0)
	e := reflect.ValueOf(entity).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i).Name
		if val, ok := e.Type().Field(i).Tag.Lookup("sql"); ok {
			field = val
		}
		valMap[strcase.ToScreamingSnake(field)] = e.Field(i).Interface()
	}

	return sq.Insert(GetTableName(entity)).SetMap(valMap).PlaceholderFormat(sq.Colon).ToSql()
}

func GetDeleteBuilder(filter, entity interface{}) (string, []interface{}, error) {
	return sq.Delete(GetTableName(entity)).Where(filter).PlaceholderFormat(sq.Colon).ToSql()
}

func GetListBuilder(offset, limit int64, orderBy string, selects []string, filters, entity interface{}) (string, []interface{}, error) {
	return sq.Select(GetColumn(selects)).
		From(GetTableName(entity)).
		Where(filters).
		OrderBy(orderBy).
		PlaceholderFormat(sq.Colon).
		Offset(uint64(offset)).
		Suffix("ROWS FETCH NEXT ? ROWS ONLY", limit).
		ToSql()
}

func Clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func GetUpdateBuilder(filter, entity interface{}) (string, []interface{}, error) {
	valMap := make(map[string]interface{}, 0)
	e := reflect.ValueOf(entity).Elem()
	for i := 0; i < e.NumField(); i++ {
		v := reflect.ValueOf(e.Field(i).Interface())
		if !v.IsZero() {
			valMap[strcase.ToScreamingSnake(e.Type().Field(i).Name)] = e.Field(i).Interface()
		}
	}

	return sq.Update(GetTableName(entity)).Where(filter).SetMap(valMap).PlaceholderFormat(sq.Colon).ToSql()
}

func GetCountBuilder(filters, entity interface{}) (string, []interface{}, error) {
	return sq.Select("COUNT(*)").
		From(GetTableName(entity)).
		Where(filters).
		PlaceholderFormat(sq.Colon).
		ToSql()
}

func GetFirstBuilder(selects []string, orderBy string, filters map[string]any, entity interface{}) (string, []interface{}, error) {
	return sq.Select(GetColumn(selects)).
		From(GetTableName(entity)).
		Where(filters).
		OrderBy(orderBy + " ASC").
		PlaceholderFormat(sq.Colon).
		Suffix("FETCH FIRST 1 ROW ONLY").
		ToSql()
}

func GetLastBuilder(selects []string, orderBy string, filters map[string]any, entity interface{}) (string, []interface{}, error) {
	return sq.Select(GetColumn(selects)).
		From(GetTableName(entity)).
		Where(filters).
		OrderBy(orderBy + " DESC").
		PlaceholderFormat(sq.Colon).
		Suffix("FETCH FIRST 1 ROW ONLY").
		ToSql()
}

func GetFindBuilder(orderBy string, selects []string, filters, entity interface{}) (string, []interface{}, error) {
	return sq.Select(GetColumn(selects)).
		From(GetTableName(entity)).
		Where(filters).
		OrderBy(orderBy).
		PlaceholderFormat(sq.Colon).
		ToSql()
}

func GetColumn(selects []string) string {
	column := "*"

	if len(selects) > 1 {
		column = strings.Join(selects, ", ")
		if strings.Contains(column, "*") {
			column = "*"
		}
	}

	return column
}
