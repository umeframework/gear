package orm

import (
	"reflect"
	"bytes"
	"strings"
	"database/sql"
)

// 构建SELECT SQL文
func (this *Orm) BuildSqlSelect(entity Entity, orderByList []OrderByCondition) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}
	var sqlCondition bytes.Buffer
	var sqlParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		typeName := field.Type.String()
		column := entMetadata.Columns[name].Column
		value := rftValue.Field(i).Interface()

		if this.isNotNull(typeName, value)  {
			sqlCondition.WriteString(column)
			sqlCondition.WriteString("=? AND ")
			sqlParamList = append(sqlParamList, value)
		}
	}

	var sql bytes.Buffer
	sql.WriteString(entMetadata.SQLSelectDefault)
	if sqlCondition.Len() > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.TrimRight(sqlCondition.String(), " AND "))
	}
	if orderByList != nil && len(orderByList) > 0 {
		sql.WriteString(" ORDER BY ")
		for i,orderBy := range orderByList {
			if i > 0 {
				sql.WriteString(",")
			}
			sql.WriteString(orderBy.Name)
			if orderBy.DESC {
				sql.WriteString(" DESC")
			}
		}
	}
	return  sql.String(), sqlParamList
}

// 构建COUNT SQL文
func (this *Orm) BuildSqlCount(entity Entity) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}

	var sqlCondition bytes.Buffer
	var sqlParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		typeName := field.Type.String()
		column := entMetadata.Columns[name].Column
		value := rftValue.Field(i).Interface()

		if this.isNotNull(typeName, value)  {
			sqlCondition.WriteString(column)
			sqlCondition.WriteString("=? AND ")
			sqlParamList = append(sqlParamList, value)
		}
	}
	sql := entMetadata.SQLSelectCountDefault
	if sqlCondition.Len() > 0 {
		sql += " WHERE " + strings.TrimRight(sqlCondition.String(), " AND ")
	}
	return  sql, sqlParamList
}

// 构建主键SELECT SQL文
func (this *Orm) BuildSqlSelectOne(entity Entity) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}
	var sqlCondition bytes.Buffer
	var sqlParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		typeName := field.Type.String()
		column := entMetadata.Columns[name].Column
		key := entMetadata.Columns[name].Key
		value := rftValue.Field(i).Interface()

		if key {
			if this.isNotNull(typeName, value) {
				sqlCondition.WriteString(column)
				sqlCondition.WriteString("=? AND ")
				sqlParamList = append(sqlParamList, value)
			} else {
				panic("Primary key parameter can not be empty.")
			}
		}
	}
	sql := entMetadata.SQLSelectDefault
	if sqlCondition.Len() > 0 {
		sql += " WHERE " + strings.TrimRight(sqlCondition.String(), " AND ")
	}
	return  sql, sqlParamList
}

// 构建UPDATE SQL文
func (this *Orm) BuildSqlUpdate(entity Entity) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}

	var sqlItem bytes.Buffer
	var sqlValue bytes.Buffer
	sqlItem.WriteString("UPDATE ")
	sqlItem.WriteString(entMetadata.Table)
	sqlItem.WriteString(" SET ")
	var sqlItemParamList []interface{}
	var sqlValueParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		colMetadata := entMetadata.Columns[name]
		typeName := field.Type.String()
		column := colMetadata.Column
		key := colMetadata.Key
		version := colMetadata.VersionCheck
		value := rftValue.Field(i).Interface()

		if this.isNotNull(typeName, value) {
			sqlItem.WriteString(column)
			sqlItem.WriteString("=?,")
			sqlItemParamList = append(sqlItemParamList, value)
		}
		if key {
			if this.isNotNull(typeName, value) {
				sqlValue.WriteString(column)
				sqlValue.WriteString("=? AND ")
				sqlValueParamList = append(sqlValueParamList, value)
			} else {
				panic("Primary key parameter can not be empty.")
			}
		} else if version {
			if this.isNotNull(typeName, value) {
				sqlValue.WriteString(column)
				sqlValue.WriteString("=? AND ")
				sqlValueParamList = append(sqlValueParamList, value)
			}
		}
	}
	sql := strings.TrimRight(sqlItem.String(), ",")
	if sqlValue.Len() > 0 {
		sql += " WHERE " + strings.TrimRight(sqlValue.String(), " AND ")
	}

	for _,e := range sqlValueParamList {
		sqlItemParamList = append(sqlItemParamList, e)
	}
	return  sql, sqlItemParamList
}

// 构建DELETE SQL文
func (this *Orm) BuildSqlDelete(entity Entity) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}

	var sqlCondition bytes.Buffer
	sqlCondition.WriteString("DELETE FROM ")
	sqlCondition.WriteString(entMetadata.Table)
	sqlCondition.WriteString(" WHERE ")
	var sqlParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		typeName := field.Type.String()
		column := entMetadata.Columns[name].Column
		key := entMetadata.Columns[name].Key
		value := rftValue.Field(i).Interface()

		if key {
			if this.isNotNull(typeName, value) {
				sqlCondition.WriteString(column)
				sqlCondition.WriteString("=? AND ")
				sqlParamList = append(sqlParamList, value)
			} else {
				panic("Primary key parameter can not be empty.")
			}
		}
	}
	return  strings.TrimRight(sqlCondition.String(), " AND "), sqlParamList
}

// 构建INSERT SQL文
func (this *Orm) BuildSqlInsert(entity Entity) (string, []interface{}) {
	entMetadata := GetEntityMetadata(entity)
	var rftType reflect.Type
	var rftValue reflect.Value
	if reflect.Ptr == reflect.TypeOf(entity).Kind() {
		rftType = reflect.TypeOf(entity).Elem()
		rftValue = reflect.ValueOf(entity).Elem()
	} else {
		rftType = reflect.TypeOf(entity)
		rftValue = reflect.ValueOf(entity)
	}

	var sqlItem bytes.Buffer
	var sqlValue bytes.Buffer
	sqlItem.WriteString("INSERT INTO ")
	sqlItem.WriteString(entMetadata.Table)
	sqlItem.WriteString("(")
	sqlValue.WriteString(") VALUES(")
	var sqlParamList []interface{}
	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		name := field.Name
		typeName := field.Type.String()
		column := entMetadata.Columns[name].Column
		key := entMetadata.Columns[name].Key
		value := rftValue.Field(i).Interface()

		if this.isNotNull(typeName, value) {
			sqlItem.WriteString(column)
			sqlItem.WriteString(",")
			sqlValue.WriteString("?,")
			sqlParamList = append(sqlParamList, value)
		} else {
			if key {
				panic("Primary key value can not be empty.")
			}
		}
	}
	return  strings.TrimRight(sqlItem.String(),",") + strings.TrimRight(sqlValue.String(),",") + ")", append(sqlParamList)
}

// 检查空值
func (this *Orm) isNotNull(typeName string, value interface{}) (bool) {
	var notNull bool
	//var r interface{}
	switch typeName {
	case "sql.NullString":
		{
			v := value.(sql.NullString)
			notNull = v.Valid //&& v.String != ""
			break
		}
	case "sql.NullInt64":
		{
			notNull = value.(sql.NullInt64).Valid
			break
		}
	case "sql.NullFloat64":
		{
			notNull = value.(sql.NullFloat64).Valid
			break
		}
	case "sql.NullBool":
		{
			notNull = value.(sql.NullBool).Valid
			break
		}
	//case "string":
	//	{
	//		notNull = value.(string) != ""
	//		break
	//	}
	default:
		{
			notNull = true
			break
		}
	}
	return notNull
}




