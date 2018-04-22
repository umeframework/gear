package orm

import (
	"database/sql"
	"sync"
)

// Orm Context
type OrmContext struct {
	conn *sql.DB
}
// 'OrmContext'指针变量
var ormContext *OrmContext
// 同步控制变量
var ormContextLock sync.Once
// 单例'OrmContext'
func singleOrmContext(driver string, dataSource string) OrmContext {
	ormContextLock.Do(func() {
		ormContext = new(OrmContext)
		db, err := sql.Open(driver, dataSource)
		if err != nil {
			panic(err)
		}
		ormContext.conn = db

	})
	return *ormContext
}

// 获取上下文
func GetOrmContext(driver string, dataSource string) OrmContext {
	return singleOrmContext(driver, dataSource)
}

// 释放上下文
func (owner *OrmContext) Close() {
	owner.conn.Close()
}

// 获取数据库访问实例
func (owner *OrmContext) DB()  *sql.DB {
	return owner.conn
}




