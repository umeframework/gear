package orm

import (
    "context"
)

// O/R Mapping实例
type Orm struct {
}

// SQL retrieve order by condition
type OrderByCondition struct {
    Name string
    DESC   bool
}

// 执行数据库查询
func (this *Orm) Retrieve(ctx OrmContext, sql string, sqlParams ...interface{}) (*OrmRows) {
    rows, err := ctx.DB().Query(sql, sqlParams[:]...)
    ormRows := newOrmRows(rows, err)

    if err != nil {
        panic(err)
    }
    return ormRows
}

// 执行数据库'Count'查询
func (this *Orm) Count(ctx OrmContext, sqlText string, sqlParams ...interface{}) int64 {
    rows, err  := ctx.DB().QueryContext(context.Background(), sqlText, sqlParams[0:]...)
    if err != nil {
        panic(err)
    }
    rows.Next()
    var count int64
    error := rows.Scan(&count)
    if error != nil {
        panic(error)
    }
    return count
}

// 执行数据库更新
func (this *Orm) Exec(ctx OrmContext, sqlText string, sqlParams ...interface{}) (*OrmResult, error) {
    execResult, err := ctx.DB().Exec(sqlText, sqlParams[:]...)
    result := newOrmResult(execResult)
    return result, err
}

// 执行数据库插入
func (this *Orm) Insert(ctx OrmContext, sqlText string, sqlParams ...interface{}) int64 {
    ormResult, execErr := this.Exec(ctx, sqlText, sqlParams[:]...)
    if execErr != nil {
        panic(execErr)
    }
    insertId, readErr := ormResult.LastInsertId()
    if readErr != nil {
        panic(readErr)
    }
    return insertId
}

// 执行数据库更新
func (this *Orm) Update(ctx OrmContext, sqlText string, sqlParams ...interface{}) int64 {
    ormResult, execErr := this.Exec(ctx, sqlText, sqlParams[:]...)
    if execErr != nil {
        panic(execErr)
    }
    affected, readErr := ormResult.RowsAffected()
    if readErr != nil {
        panic(readErr)
    }
    return affected
}

// 执行数据库更新
func (this *Orm) Delete(ctx OrmContext, sqlText string, sqlParams ...interface{}) int64 {
    ormResult, execErr := this.Exec(ctx, sqlText, sqlParams[:]...)
    if execErr != nil {
        panic(execErr)
    }
    affected, readErr := ormResult.RowsAffected()
    if readErr != nil {
        panic(readErr)
    }
    return affected
}

