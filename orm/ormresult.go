package orm

import (
	"database/sql"
	"errors"
	"reflect"
)

// Errors定义
var (
	errorRowsNotSpecified = errors.New("no sql.Rows specified.")
	errorRecordNotFound = errors.New("No records found.")
)

// O/R Mapping 结果集
type OrmResult struct {
	result sql.Result
	err    error
}

// O/R Mapping 结果集
type OrmRows struct {
	rows       *sql.Rows
	err        error
	closed bool
}

// 创建'OrmRows'实例
func newOrmRows(rows *sql.Rows, err error) *OrmRows {
	return &OrmRows{rows, err, false}
}

// 创建'OrmResult'实例
func newOrmResult(execResult sql.Result) *OrmResult {
	return &OrmResult{execResult, errorRowsNotSpecified}
}

// 获取'*sql.Rows'
func (this *OrmRows) Rows() (*sql.Rows, error) {
	return this.rows, this.err
}

// 获取最近插入记录的序号
func (this *OrmResult) LastInsertId() (int64, error) {
	return this.result.LastInsertId()
}

// 获取最近更新的记录数
func (this *OrmResult) RowsAffected() (int64, error) {
	return this.result.RowsAffected()
}

// 关闭'*sql.Rows'
func (this *OrmRows) Close() error {
	var err error
	if !this.closed {
		err = this.rows.Close()
		this.closed = true
	}
	return err
}

// 查询结果映射处理
func (this *OrmRows) DefaultMapping(dest interface{}) error {
	return this.Mapping(dest, nil)
}

// 带有特定mapper的查询结果映射处理
//func (this *OrmRows) Mapping(tar interface{}, mapper OrmMapper) error {
func (this *OrmRows) Mapping(tar interface{}, mapper func(entity interface{}) []interface{}) error {
	ormMapper := NewSimpleCallbackMapper(mapper)
	var err error = nil
	// Check tar type: Must be pointer
	t := reflect.TypeOf(tar)
	if t.Kind() != reflect.Ptr {
		return errors.New("[tar] parameter must be a pointer to object or a slice")
	}

	elemType := t.Elem()
	if elemType.Kind() == reflect.Slice {
		err = this.mapToSlice(tar, elemType, ormMapper)
	} else {
		err = this.mapToObject(tar, elemType, ormMapper)
	}

	// Exit
	return err
}

// 映射sql.Rows数据至目标实例
func (this *OrmRows) mapToSlice(tar interface{}, sliceType reflect.Type, mapper OrmMapper) error {
	var err error = nil

	elemType := sliceType.Elem()

	if mapper == nil {
		mapper = newDefaultOrmMapper(elemType)
	}

	// Create slice object
	sliceObj := reflect.New(sliceType)
	slice := sliceObj.Elem()

	// Read each row and convert to object
	for this.rows.Next() {
		// Create new element object
		elem := reflect.New(elemType)

		// DefaultMapping row to object
		this.mapRowToObject(this.rows, elem.Interface(), mapper)

		// Add to slice
		slice = reflect.Append(slice, elem.Elem())
	}

	// Write slice object back to tar interface{}
	reflect.ValueOf(tar).Elem().Set(slice)

	// Exit
	return err
}

// 映射sql.Rows数据至目标实例
func (this *OrmRows) mapToObject(tar interface{}, elemType reflect.Type, mapper OrmMapper) error {
	if !this.rows.Next() {
		return errorRecordNotFound
	}

	if mapper == nil {
		mapper = newDefaultOrmMapper(elemType)
	}

	return this.mapRowToObject(this.rows, tar, mapper)
}

// 映射sql.Rows数据至目标实例
func (this *OrmRows) mapRowToObject(row *sql.Rows, tar interface{}, mapper OrmMapper) error {
	var err error = nil
	if mapper != nil {
		err = mapper.Mapping(row, tar)
	}
	return err
}
