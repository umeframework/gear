package orm

import (
	"reflect"
	"database/sql"
	"strings"
	"errors"
)

//
type MapperCallback func(row *sql.Rows, result interface{}) error
//
type SimpleMapperCallback func(result interface{}) []interface{}

type defaultSimpleMapper struct {
	t reflect.Type
}
//
type callbackMapper struct {
	callback MapperCallback
}
type ormTypeInfo struct {
	t reflect.Type
	fields []reflect.StructField
	fieldMap map[string]*reflect.StructField
}
type defaultStructMapper struct {
	typeInfo *ormTypeInfo
}


// Cached  OrmMapper instance
var (
	ormMapper OrmMapper
)
var (
	typeInfoMap = make(map[reflect.Type]*ormTypeInfo)
)

func (this callbackMapper) Mapping(row *sql.Rows, result interface{}) error {
	return this.callback(row, result)
}

//
func NewCallbackMapper(callback MapperCallback) OrmMapper {
	return callbackMapper{callback}
}

func NewSimpleCallbackMapper(callback SimpleMapperCallback) OrmMapper {
	return NewCallbackMapper(func(row *sql.Rows, result interface{}) error {
		scanMapping := callback(result)
		return row.Scan(scanMapping...)
	})
}

func NewSimpleValueMapper() OrmMapper {
	if ormMapper == nil {
		ormMapper = NewCallbackMapper(func(row *sql.Rows, result interface{}) error {
			return row.Scan(result)
		})
	}
	return ormMapper
}

func newDefaultOrmMapper(t reflect.Type) OrmMapper {
	var mapper OrmMapper = nil

	if t.Kind() == reflect.Struct {
		// Retrieve type info
		// (Use a caching map to reduce reflect operations)
		typeInfoPtr, found := typeInfoMap[t]
		if !found {
			typeInfo := newOrmTypeInfo(t)
			typeInfoPtr = &typeInfo
			typeInfoMap[t] = typeInfoPtr
		}

		// Create Mapper
		mapper = defaultStructMapper{typeInfoPtr}
	} else {
		mapper = defaultSimpleMapper{t}
	}
	return mapper
}

func (this defaultSimpleMapper) Mapping(row *sql.Rows, result interface{}) error {
	var err error = nil
	err = row.Scan(result)
	return err
}

func normalizeFieldName(fieldName string) string {
	return strings.ToUpper(fieldName)
}

func newOrmTypeInfo(t reflect.Type) ormTypeInfo {
	typeInfo := ormTypeInfo{}
	typeInfo.t = t

	// Get field Info one by one
	if t.Kind() == reflect.Struct {
		fieldCount := t.NumField()
		typeInfo.fields = make([]reflect.StructField, 0, fieldCount)
		typeInfo.fieldMap = make(map[string]*reflect.StructField)
		for i := 0; i < fieldCount; i++ {
			fieldInfo := t.Field(i)
			typeInfo.fields = append(typeInfo.fields, fieldInfo)
			typeInfo.fieldMap[normalizeFieldName(fieldInfo.Name)] = &fieldInfo
		}
	}

	// Exit
	return typeInfo
}

// A simple mapper only matching with column names
func (self defaultStructMapper) Mapping(row *sql.Rows, result interface{}) error {
	var err error = nil
	var columnNames []string

	if columnNames, err = row.Columns(); err != nil {
		return err
	}
	columnCount := len(columnNames)
	fieldInfos := make([]*reflect.StructField, 0, columnCount)
	columnMappings := make([]interface{}, 0, columnCount)
	for i := 0; i < columnCount; i++ {
		columnName := columnNames[i]
		fieldInfo := self.findFieldInfo(columnName)
		if fieldInfo == nil {
			return errors.New("field info not found for column: " + columnName)
		}
		fieldInfos = append(fieldInfos, fieldInfo)

		// Create object to store column value
		fieldValue := reflect.New(fieldInfo.Type)
		columnMappings = append(columnMappings, fieldValue.Interface())
	}

	// Scan from row
	if err = row.Scan(columnMappings...); err != nil {
		return err
	}

	// DefaultMapping row value to column fields
	resultElem := reflect.ValueOf(result).Elem()
	for i := 0; i < columnCount; i++ {
		fieldInfo := fieldInfos[i]
		columnMapping := columnMappings[i]
		fieldValue := reflect.ValueOf(columnMapping).Elem()
		//fmt.Println(fieldValue)
		resultElem.Field(fieldInfo.Index[0]).Set(fieldValue)
	}

	// Exit
	return err
}

func (self *defaultStructMapper) findFieldInfo(columnName string) *reflect.StructField {
	var fieldInfo *reflect.StructField = nil
	var found bool
	fieldInfo, found = self.typeInfo.fieldMap[normalizeFieldName(columnName)]
	if !found {
		fieldInfo = nil
	}
	return fieldInfo
}