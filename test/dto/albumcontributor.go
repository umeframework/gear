package dto
import (
    "database/sql"
    ."github.com/umeframework/gear"
    ."github.com/umeframework/gear/orm"
)

// '参加该唱片录音的艺术家信息管理表'表实体结构(基础类型描述)
type AlbumContributorDto struct {
    // 唱片编号
    AlbumId int64
    // 参与曲目
    TrackNo int64
    // 艺术家
    Contributor string
    // 备注
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

// '参加该唱片录音的艺术家信息管理表'表实体结构(SQL类型描述)
type AlbumContributorEntity struct {
    // 唱片编号 
    AlbumId sql.NullInt64 `name:"ALBUM_ID", type:"INT", comment:"唱片编号", key:true, notnull:true`
    // 参与曲目 
    TrackNo sql.NullInt64 `name:"TRACK_NO", type:"INT", comment:"参与曲目", key:true, notnull:true`
    // 艺术家 
    Contributor sql.NullString `name:"CONTRIBUTOR", type:"VARCHAR", comment:"艺术家", key:true, notnull:true`
    // 备注 
    Comment sql.NullString `name:"COMMENT", type:"VARCHAR", comment:"备注", key:false, notnull:false`
    // 创建者
    CreateAuthor sql.NullString `name:"CREATE_AUTHOR",  type:"VARCHAR", comment:"创建者", key:false, notnull:false`
    // 创建时间
    CreateDatetime sql.NullString `name:"CREATE_DATETIME",  type:"TIMESTAMP", comment:"创建时间", key:false, notnull:false`
    // 更新者
    UpdateAuthor sql.NullString `name:"UPDATE_AUTHOR",  type:"VARCHAR", comment:"更新者", key:false, notnull:false`
    // 更新时间
    UpdateDatetime sql.NullString `name:"UPDATE_DATETIME",  type:"TIMESTAMP", comment:"更新时间", key:false, notnull:false`
}

// 返回'参加该唱片录音的艺术家信息管理表'表名
func (owner *AlbumContributorEntity) TableName() string {
    return "ALBUM_CONTRIBUTOR"
}

// 从'map'创建
func (owner *AlbumContributorEntity) FromMap(src map[string]interface{}) *AlbumContributorEntity {
    var value interface{}
    var exist bool
    value,exist = src["AlbumId"]
    if exist {
        owner.AlbumId = sql.NullInt64{Int64:value.(int64), Valid:true}
    }
    value,exist = src["TrackNo"]
    if exist {
        owner.TrackNo = sql.NullInt64{Int64:value.(int64), Valid:true}
    }
    value,exist = src["Contributor"]
    if exist {
        owner.Contributor = sql.NullString{String:value.(string), Valid:true}
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
func (owner *AlbumContributorEntity) ToMap() map[string]interface{} {
    tar := make(map[string]interface{})
    if owner.AlbumId.Valid {
        tar["AlbumId"] = owner.AlbumId.Int64
    }
    if owner.TrackNo.Valid {
        tar["TrackNo"] = owner.TrackNo.Int64
    }
    if owner.Contributor.Valid {
        tar["Contributor"] = owner.Contributor.String
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

// 从'AlbumContributorDto'创建
func (owner *AlbumContributorEntity) FromDto(src AlbumContributorDto) *AlbumContributorEntity {
    owner.AlbumId = sql.NullInt64{Int64:src.AlbumId, Valid:true}
    owner.TrackNo = sql.NullInt64{Int64:src.TrackNo, Valid:true}
    owner.Contributor = sql.NullString{String:src.Contributor, Valid:true}
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

// 变换为'AlbumContributorDto'
func (owner *AlbumContributorEntity) ToDto() AlbumContributorDto {
    var tar AlbumContributorDto
    tar.AlbumId = owner.AlbumId.Int64
    tar.TrackNo = owner.TrackNo.Int64
    tar.Contributor = owner.Contributor.String
    tar.Comment = owner.Comment.String
    tar.CreateAuthor = owner.CreateAuthor.String
    tar.CreateDatetime = owner.CreateDatetime.String
    tar.UpdateAuthor = owner.UpdateAuthor.String
    tar.UpdateDatetime = owner.UpdateDatetime.String
    return tar
}

// 创建行数据映射列表
func (owner *AlbumContributorEntity) Mapper(entity interface{}) []interface{} {
    e := entity.(*AlbumContributorEntity)
    return []interface{}{ &e.AlbumId,&e.TrackNo,&e.Contributor,&e.Comment,&e.CreateAuthor,&e.CreateDatetime,&e.UpdateAuthor,&e.UpdateDatetime}
}

// 查询
func (owner *AlbumContributorEntity) Retrieve(ctx OrmContext, orderBy ...OrderByCondition) []AlbumContributorEntity {
    sqlText, sqlParams := GetDao().BuildSqlSelect(owner, orderBy)
    var nl []AlbumContributorEntity
    GetDao().Retrieve(ctx, sqlText, sqlParams[:]...).Mapping(&nl, owner.Mapper)
    return nl
}

// 查询(使用map类型参数查询)
func (owner *AlbumContributorEntity) RetrieveByMap(ctx OrmContext, param map[string]interface{}, orderBy ...OrderByCondition) []AlbumContributorEntity {
    return owner.FromMap(param).Retrieve(ctx,orderBy[:]...)
}

// 统计
func (owner *AlbumContributorEntity) Count(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlCount(owner)
    return GetDao().Count(ctx, sqlText, sqlParams[:]...)
}

// 登录
func (owner *AlbumContributorEntity) Insert(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlInsert(owner)
    return GetDao().Insert(ctx, sqlText, sqlParams[:]...)
}

// 更新
func (owner *AlbumContributorEntity) Update(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlUpdate(owner)
    return GetDao().Update(ctx, sqlText, sqlParams[:]...)
}

// 删除
func (owner *AlbumContributorEntity) Delete(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlDelete(owner)
    return GetDao().Delete(ctx, sqlText, sqlParams[:]...)
}







