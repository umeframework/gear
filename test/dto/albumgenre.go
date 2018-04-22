package dto
import (
    "database/sql"
    ."github.com/umeframework/gear"
    ."github.com/umeframework/gear/orm"
)

// '唱片风格分类描述表'表实体结构(基础类型描述)
type AlbumGenreDto struct {
    // 风格编码
    GenreId string
    // 风格名称
    GenreName string
    // 风格描述
    Comment string
    // 创建者
    CreateAuthor string
    // 创建时间
    CreateDatetime string
    // 更新者
    UpdateAuthor string
    // 更新时间
    UpdateDatetime string
}

// '唱片风格分类描述表'表实体结构(SQL类型描述)
type AlbumGenreEntity struct {
    // 风格编码 
    GenreId sql.NullString `name:"GENRE_ID", type:"CHAR", comment:"风格编码", key:true, notnull:true`
    // 风格名称 
    GenreName sql.NullString `name:"GENRE_NAME", type:"VARCHAR", comment:"风格名称", key:false, notnull:true`
    // 风格描述 
    Comment sql.NullString `name:"COMMENT", type:"VARCHAR", comment:"风格描述", key:false, notnull:false`
    // 创建者
    CreateAuthor sql.NullString `name:"CREATE_AUTHOR",  type:"VARCHAR", comment:"创建者", key:false, notnull:false`
    // 创建时间
    CreateDatetime sql.NullString `name:"CREATE_DATETIME",  type:"TIMESTAMP", comment:"创建时间", key:false, notnull:false`
    // 更新者
    UpdateAuthor sql.NullString `name:"UPDATE_AUTHOR",  type:"VARCHAR", comment:"更新者", key:false, notnull:false`
    // 更新时间
    UpdateDatetime sql.NullString `name:"UPDATE_DATETIME",  type:"TIMESTAMP", comment:"更新时间", key:false, notnull:false`
}

// 返回'唱片风格分类描述表'表名
func (owner *AlbumGenreEntity) TableName() string {
    return "ALBUM_GENRE"
}

// 从'map'创建
func (owner *AlbumGenreEntity) FromMap(src map[string]interface{}) *AlbumGenreEntity {
    var value interface{}
    var exist bool
    value,exist = src["GenreId"]
    if exist {
        owner.GenreId = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["GenreName"]
    if exist {
        owner.GenreName = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["Comment"]
    if exist {
        owner.Comment = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["CreateAuthor"]
    if exist {
        owner.CreateAuthor = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["CreateDatetime"]
    if exist {
        owner.CreateDatetime = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["UpdateAuthor"]
    if exist {
        owner.UpdateAuthor = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["UpdateDatetime"]
    if exist {
        owner.UpdateDatetime = sql.NullString{String:value.(string), Valid:true}
    }
    return owner
}

// 变换为'map'
func (owner *AlbumGenreEntity) ToMap() map[string]interface{} {
    tar := make(map[string]interface{})
    if owner.GenreId.Valid {
        tar["GenreId"] = owner.GenreId.String
    }
    if owner.GenreName.Valid {
        tar["GenreName"] = owner.GenreName.String
    }
    if owner.Comment.Valid {
        tar["Comment"] = owner.Comment.String
    }
    if owner.CreateAuthor.Valid {
        tar["CreateAuthor"] = owner.CreateAuthor.String
    }
    if owner.CreateDatetime.Valid {
        tar["CreateDatetime"] = owner.CreateDatetime.String
    }
    if owner.UpdateAuthor.Valid {
        tar["UpdateAuthor"] = owner.UpdateAuthor.String
    }
    if owner.UpdateDatetime.Valid {
        tar["UpdateDatetime"] = owner.UpdateDatetime.String
    }
    return tar
}

// 从'AlbumGenreDto'创建
func (owner *AlbumGenreEntity) FromDto(src AlbumGenreDto) *AlbumGenreEntity {
    owner.GenreId = sql.NullString{String:src.GenreId, Valid:true}
    owner.GenreName = sql.NullString{String:src.GenreName, Valid:true}
    owner.Comment = sql.NullString{String:src.Comment, Valid:true}
    owner.CreateAuthor = sql.NullString{String:src.CreateAuthor, Valid:true}
    if src.CreateDatetime != "" {
        owner.CreateDatetime = sql.NullString{String:src.CreateDatetime, Valid:true}
    }
    owner.UpdateAuthor = sql.NullString{String:src.UpdateAuthor, Valid:true}
    if src.CreateDatetime != "" {
        owner.UpdateDatetime = sql.NullString{String:src.UpdateDatetime, Valid:true}
    }
    return owner
}

// 变换为'AlbumGenreDto'
func (owner *AlbumGenreEntity) ToDto() AlbumGenreDto {
    var tar AlbumGenreDto
    tar.GenreId = owner.GenreId.String
    tar.GenreName = owner.GenreName.String
    tar.Comment = owner.Comment.String
    tar.CreateAuthor = owner.CreateAuthor.String
    tar.CreateDatetime = owner.CreateDatetime.String
    tar.UpdateAuthor = owner.UpdateAuthor.String
    tar.UpdateDatetime = owner.UpdateDatetime.String
    return tar
}

// 创建行数据映射列表
func (owner *AlbumGenreEntity) Mapper(entity interface{}) []interface{} {
    e := entity.(*AlbumGenreEntity)
    return []interface{}{ &e.GenreId,&e.GenreName,&e.Comment,&e.CreateAuthor,&e.CreateDatetime,&e.UpdateAuthor,&e.UpdateDatetime}
}

// 查询
func (owner *AlbumGenreEntity) Retrieve(ctx OrmContext, orderBy ...OrderByCondition) []AlbumGenreEntity {
    sqlText, sqlParams := GetDao().BuildSqlSelect(owner, orderBy)
    var nl []AlbumGenreEntity
    GetDao().Retrieve(ctx, sqlText, sqlParams[:]...).Mapping(&nl, owner.Mapper)
    return nl
}

// 查询(使用map类型参数查询)
func (owner *AlbumGenreEntity) RetrieveByMap(ctx OrmContext, param map[string]interface{}, orderBy ...OrderByCondition) []AlbumGenreEntity {
    return owner.FromMap(param).Retrieve(ctx,orderBy[:]...)
}

// 统计
func (owner *AlbumGenreEntity) Count(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlCount(owner)
    return GetDao().Count(ctx, sqlText, sqlParams[:]...)
}

// 登录
func (owner *AlbumGenreEntity) Insert(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlInsert(owner)
    return GetDao().Insert(ctx, sqlText, sqlParams[:]...)
}

// 更新
func (owner *AlbumGenreEntity) Update(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlUpdate(owner)
    return GetDao().Update(ctx, sqlText, sqlParams[:]...)
}

// 删除
func (owner *AlbumGenreEntity) Delete(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlDelete(owner)
    return GetDao().Delete(ctx, sqlText, sqlParams[:]...)
}







