package orm

import (
	"reflect"
	"bytes"
	"strings"
	"strconv"
	"sync"
)

// 实体Metadata保存数据结构
type EntityConfig struct {
	Entities map[string]EntityMetadata
}

// 实体(表)Metadata数据结构
type EntityMetadata struct {
	Table                 string
	TableComment          string
	Columns               map[string]ColumnMetadata
	SQLInsertDefault      string
	SQLUpdateDefault      string
	SQLDeleteDefault      string
	SQLSelectDefault      string
	SQLSelectOneDefault   string
	SQLSelectCountDefault string
}

// 实体字段(列)Metadata数据结构
type ColumnMetadata struct {
	FieldId         string
	FieldIndex      int
	FieldType       reflect.Type
	Column          string
	ColumnType      string
	ColumnComment   string
	Key             bool
	NotNull         bool
	VersionCheck    bool
}

// 'EntityConfig'指针变量
var entityConfig *EntityConfig
// 同步控制变量
var entityConfigLock sync.Once
// 单例'EntityConfig'
func singleEntityConfig() *EntityConfig {
	entityConfigLock.Do(func() {
		entityConfig = new(EntityConfig)                        // Init EntityConfig
		entityConfig.Entities = make(map[string]EntityMetadata) // Init map
	})
	return entityConfig
}

//+ 根据实体实例获取Metadata
func GetEntityMetadata(entity Entity) EntityMetadata {
	instance := singleEntityConfig()
	tableName := entity.TableName()
	entMetadata, exist := instance.Entities[tableName]
	if exist {
		return entMetadata
	}
	nem := parseEntity(entity)
	instance.Entities[tableName] = nem
	return instance.Entities[tableName]
}

// 解析实体实例获得Metadata
func parseEntity(entity Entity) EntityMetadata {
	var sqlSelect bytes.Buffer
	var sqlInsertItem bytes.Buffer
	var sqlInsertValue bytes.Buffer
	var sqlUpdateItem bytes.Buffer
	var sqlUpdateValue bytes.Buffer
	var sqlDelete bytes.Buffer
	sqlSelect.WriteString("SELECT ")
	sqlInsertItem.WriteString("INSERT INTO ")
	sqlInsertItem.WriteString(entity.TableName())
	sqlInsertItem.WriteString("(")
	sqlInsertValue.WriteString(" VALUES(")
	sqlUpdateItem.WriteString("UPDATE ")
	sqlUpdateItem.WriteString(entity.TableName())
	sqlUpdateItem.WriteString(" SET ")
	sqlDelete.WriteString("DELETE FROM ")
	sqlDelete.WriteString(entity.TableName())

	rftType := reflect.TypeOf(entity).Elem()

	var entMetadata EntityMetadata
	colMetadataMap := make(map[string]ColumnMetadata)

	for i := 0; i < rftType.NumField(); i++ {
		field := rftType.Field(i)
		fieldName := field.Name
		fieldTag := string(field.Tag)
		if fieldTag == "" {
			continue
		}
		var colMetadata ColumnMetadata
		colMetadata.FieldId = field.Name
		colMetadata.FieldType = field.Type
		colMetadata.FieldIndex = i
		tagElements := strings.Split(fieldTag, ",")
		for _, e := range tagElements {
			e = strings.TrimSpace(e)
			if strings.HasPrefix(e, "name:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "name:", "", -1)), "\"")
				colMetadata.Column = val
				sqlSelect.WriteString(val)
				sqlSelect.WriteString(" AS `")
				sqlSelect.WriteString(field.Name)
				sqlSelect.WriteString("`,")
				sqlInsertItem.WriteString(val)
				sqlInsertItem.WriteString(",")
				sqlInsertValue.WriteString("?,")
				sqlUpdateItem.WriteString(val)
				sqlUpdateItem.WriteString("=?,")
			} else if strings.HasPrefix(e, "type:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "type:", "", -1)), "\"")
				colMetadata.ColumnType = val
			} else if strings.HasPrefix(e, "comment:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "comment:", "", -1)), "\"")
				colMetadata.ColumnComment = val
			} else if strings.HasPrefix(e, "key:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "key:", "", -1)), "\"")
				colMetadata.Key,_ = strconv.ParseBool(val)
				if colMetadata.Key {
					sqlUpdateValue.WriteString(colMetadata.Column)
					sqlUpdateValue.WriteString("=? AND ")
				}
			} else if strings.HasPrefix(e, "notnull:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "notnull:", "", -1)), "\"")
				colMetadata.NotNull,_ = strconv.ParseBool(val)
			} else if strings.HasPrefix(e, "version:") {
				val := strings.Trim(strings.TrimSpace(strings.Replace(e, "version:", "", -1)), "\"")
				colMetadata.VersionCheck,_ = strconv.ParseBool(val)
			}
		}
		colMetadataMap[fieldName] = colMetadata
	}

	entMetadata.Table = entity.TableName()
	entMetadata.Columns = colMetadataMap
	entMetadata.SQLInsertDefault = strings.TrimRight(sqlInsertItem.String(),",") + ")" +  strings.TrimRight(sqlInsertValue.String(),",") + ")"
	entMetadata.SQLUpdateDefault = strings.TrimRight(sqlUpdateItem.String(),",")
	entMetadata.SQLDeleteDefault = sqlDelete.String()
	entMetadata.SQLSelectDefault = strings.TrimRight(sqlSelect.String(),",") + " FROM " + entity.TableName()
	entMetadata.SQLSelectOneDefault = strings.TrimRight(sqlSelect.String(),",") + " FROM " + entity.TableName()
	entMetadata.SQLSelectCountDefault = "SELECT COUNT(*) AS `count` FROM " + entity.TableName()

	if sqlUpdateValue.Len() > 0 {
		entMetadata.SQLSelectOneDefault += " WHERE " + strings.TrimRight(sqlUpdateValue.String()," AND ")
		entMetadata.SQLUpdateDefault += " WHERE " + strings.TrimRight(sqlUpdateValue.String()," AND ")
		entMetadata.SQLDeleteDefault += " WHERE " + strings.TrimRight(sqlUpdateValue.String()," AND ")
	}
	return entMetadata
}



