package orm

import "database/sql"

// 'Entity'操作通用接口定义
type Entity interface {
    // 获取表名
    TableName() string
    //Mapper(dto interface{}) []interface{}
}

// Declare Mapping interface to handing *sql.Rows
type OrmMapper interface {
    Mapping(row *sql.Rows, result interface{}) error
}

