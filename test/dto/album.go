package dto
import (
    "database/sql"
    ."github.com/umeframework/gear"
    ."github.com/umeframework/gear/orm"
)

// '唱片基本信息表'表实体结构(基础类型描述)
type AlbumDto struct {
    // 编号
    Id int64
    // 标题
    Title string
    // 艺术家
    Artist string
    // 发行时间
    IssueDate string
    // 风格
    Genre string
    // 封面设计
    CoverPhoto string
    // 创建者
    CreateAuthor string
    // 创建时间
    CreateDatetime string
    // 更新者
    UpdateAuthor string
    // 更新时间
    UpdateDatetime string
}

// '唱片基本信息表'表实体结构(SQL类型描述)
type AlbumEntity struct {
    // 编号 
    Id sql.NullInt64 `name:"ID", type:"INT", comment:"编号", key:true, notnull:true`
    // 标题 
    Title sql.NullString `name:"TITLE", type:"VARCHAR", comment:"标题", key:false, notnull:true`
    // 艺术家 
    Artist sql.NullString `name:"ARTIST", type:"VARCHAR", comment:"艺术家", key:false, notnull:true`
    // 发行时间 
    IssueDate sql.NullString `name:"ISSUE_DATE", type:"DATE", comment:"发行时间", key:false, notnull:false`
    // 风格 
    Genre sql.NullString `name:"GENRE", type:"CHAR", comment:"风格", key:false, notnull:false`
    // 封面设计 
    CoverPhoto sql.NullString `name:"COVER_PHOTO", type:"BLOB", comment:"封面设计", key:false, notnull:false`
    // 创建者
    CreateAuthor sql.NullString `name:"CREATE_AUTHOR",  type:"VARCHAR", comment:"创建者", key:false, notnull:false`
    // 创建时间
    CreateDatetime sql.NullString `name:"CREATE_DATETIME",  type:"TIMESTAMP", comment:"创建时间", key:false, notnull:false`
    // 更新者
    UpdateAuthor sql.NullString `name:"UPDATE_AUTHOR",  type:"VARCHAR", comment:"更新者", key:false, notnull:false`
    // 更新时间
    UpdateDatetime sql.NullString `name:"UPDATE_DATETIME",  type:"TIMESTAMP", comment:"更新时间", key:false, notnull:false`
}

// 返回'唱片基本信息表'表名
func (owner *AlbumEntity) TableName() string {
    return "ALBUM"
}

// 从'map'创建
func (owner *AlbumEntity) FromMap(src map[string]interface{}) *AlbumEntity {
    var value interface{}
    var exist bool
    value,exist = src["Id"]
    if exist {
        owner.Id = sql.NullInt64{Int64:value.(int64), Valid:true}
    }
    value,exist = src["Title"]
    if exist {
        owner.Title = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["Artist"]
    if exist {
        owner.Artist = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["IssueDate"]
    if exist {
        if value != "" {
            owner.IssueDate = sql.NullString{String:value.(string), Valid:true}
        }
    }
    value,exist = src["Genre"]
    if exist {
        owner.Genre = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["CoverPhoto"]
    if exist {
        owner.CoverPhoto = sql.NullString{String:value.(string), Valid:true}
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
func (owner *AlbumEntity) ToMap() map[string]interface{} {
    tar := make(map[string]interface{})
    if owner.Id.Valid {
        tar["Id"] = owner.Id.Int64
    }
    if owner.Title.Valid {
        tar["Title"] = owner.Title.String
    }
    if owner.Artist.Valid {
        tar["Artist"] = owner.Artist.String
    }
    if owner.IssueDate.Valid {
        tar["IssueDate"] = owner.IssueDate.String
    }
    if owner.Genre.Valid {
        tar["Genre"] = owner.Genre.String
    }
    if owner.CoverPhoto.Valid {
        tar["CoverPhoto"] = owner.CoverPhoto.String
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

// 从'AlbumDto'创建
func (owner *AlbumEntity) FromDto(src AlbumDto) *AlbumEntity {
    owner.Id = sql.NullInt64{Int64:src.Id, Valid:true}
    owner.Title = sql.NullString{String:src.Title, Valid:true}
    owner.Artist = sql.NullString{String:src.Artist, Valid:true}
    if src.IssueDate != "" {
        owner.IssueDate = sql.NullString{String:src.IssueDate, Valid:true}
    }
    owner.Genre = sql.NullString{String:src.Genre, Valid:true}
    owner.CoverPhoto = sql.NullString{String:src.CoverPhoto, Valid:true}
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

// 变换为'AlbumDto'
func (owner *AlbumEntity) ToDto() AlbumDto {
    var tar AlbumDto
    tar.Id = owner.Id.Int64
    tar.Title = owner.Title.String
    tar.Artist = owner.Artist.String
    tar.IssueDate = owner.IssueDate.String
    tar.Genre = owner.Genre.String
    tar.CoverPhoto = owner.CoverPhoto.String
    tar.CreateAuthor = owner.CreateAuthor.String
    tar.CreateDatetime = owner.CreateDatetime.String
    tar.UpdateAuthor = owner.UpdateAuthor.String
    tar.UpdateDatetime = owner.UpdateDatetime.String
    return tar
}

// 创建行数据映射列表
func (owner *AlbumEntity) Mapper(entity interface{}) []interface{} {
    e := entity.(*AlbumEntity)
    return []interface{}{ &e.Id,&e.Title,&e.Artist,&e.IssueDate,&e.Genre,&e.CoverPhoto,&e.CreateAuthor,&e.CreateDatetime,&e.UpdateAuthor,&e.UpdateDatetime}
}

// 查询
func (owner *AlbumEntity) Retrieve(ctx OrmContext, orderBy ...OrderByCondition) []AlbumEntity {
    sqlText, sqlParams := GetDao().BuildSqlSelect(owner, orderBy)
    var nl []AlbumEntity
    GetDao().Retrieve(ctx, sqlText, sqlParams[:]...).Mapping(&nl, owner.Mapper)
    return nl
}

// 查询(使用map类型参数查询)
func (owner *AlbumEntity) RetrieveByMap(ctx OrmContext, param map[string]interface{}, orderBy ...OrderByCondition) []AlbumEntity {
    return owner.FromMap(param).Retrieve(ctx,orderBy[:]...)
}

// 统计
func (owner *AlbumEntity) Count(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlCount(owner)
    return GetDao().Count(ctx, sqlText, sqlParams[:]...)
}

// 登录
func (owner *AlbumEntity) Insert(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlInsert(owner)
    return GetDao().Insert(ctx, sqlText, sqlParams[:]...)
}

// 更新
func (owner *AlbumEntity) Update(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlUpdate(owner)
    return GetDao().Update(ctx, sqlText, sqlParams[:]...)
}

// 删除
func (owner *AlbumEntity) Delete(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlDelete(owner)
    return GetDao().Delete(ctx, sqlText, sqlParams[:]...)
}







