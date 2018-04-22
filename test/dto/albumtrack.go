package dto
import (
    "database/sql"
    ."github.com/umeframework/gear"
    ."github.com/umeframework/gear/orm"
)

// '唱片曲目信息表'表实体结构(基础类型描述)
type AlbumTrackDto struct {
    // 所属唱片
    AlbumId int64
    // 曲目编号
    TrackNo int64
    // 曲目名称
    TrackName string
    // 播放时间
    PlayTime float64
    // 创建者
    CreateAuthor string
    // 创建时间
    CreateDatetime string
    // 更新者
    UpdateAuthor string
    // 更新时间
    UpdateDatetime string
}

// '唱片曲目信息表'表实体结构(SQL类型描述)
type AlbumTrackEntity struct {
    // 所属唱片 
    AlbumId sql.NullInt64 `name:"ALBUM_ID", type:"INT", comment:"所属唱片", key:true, notnull:true`
    // 曲目编号 
    TrackNo sql.NullInt64 `name:"TRACK_NO", type:"INT", comment:"曲目编号", key:true, notnull:true`
    // 曲目名称 
    TrackName sql.NullString `name:"TRACK_NAME", type:"VARCHAR", comment:"曲目名称", key:false, notnull:true`
    // 播放时间 
    PlayTime sql.NullFloat64 `name:"PLAY_TIME", type:"DECIMAL", comment:"播放时间", key:false, notnull:false`
    // 创建者
    CreateAuthor sql.NullString `name:"CREATE_AUTHOR",  type:"VARCHAR", comment:"创建者", key:false, notnull:false`
    // 创建时间
    CreateDatetime sql.NullString `name:"CREATE_DATETIME",  type:"TIMESTAMP", comment:"创建时间", key:false, notnull:false`
    // 更新者
    UpdateAuthor sql.NullString `name:"UPDATE_AUTHOR",  type:"VARCHAR", comment:"更新者", key:false, notnull:false`
    // 更新时间
    UpdateDatetime sql.NullString `name:"UPDATE_DATETIME",  type:"TIMESTAMP", comment:"更新时间", key:false, notnull:false`
}

// 返回'唱片曲目信息表'表名
func (owner *AlbumTrackEntity) TableName() string {
    return "ALBUM_TRACK"
}

// 从'map'创建
func (owner *AlbumTrackEntity) FromMap(src map[string]interface{}) *AlbumTrackEntity {
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
    value,exist = src["TrackName"]
    if exist {
        owner.TrackName = sql.NullString{String:value.(string), Valid:true}
    }
    value,exist = src["PlayTime"]
    if exist {
        owner.PlayTime = sql.NullFloat64{Float64:value.(float64), Valid:true}
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
func (owner *AlbumTrackEntity) ToMap() map[string]interface{} {
    tar := make(map[string]interface{})
    if owner.AlbumId.Valid {
        tar["AlbumId"] = owner.AlbumId.Int64
    }
    if owner.TrackNo.Valid {
        tar["TrackNo"] = owner.TrackNo.Int64
    }
    if owner.TrackName.Valid {
        tar["TrackName"] = owner.TrackName.String
    }
    if owner.PlayTime.Valid {
        tar["PlayTime"] = owner.PlayTime.Float64
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

// 从'AlbumTrackDto'创建
func (owner *AlbumTrackEntity) FromDto(src AlbumTrackDto) *AlbumTrackEntity {
    owner.AlbumId = sql.NullInt64{Int64:src.AlbumId, Valid:true}
    owner.TrackNo = sql.NullInt64{Int64:src.TrackNo, Valid:true}
    owner.TrackName = sql.NullString{String:src.TrackName, Valid:true}
    owner.PlayTime = sql.NullFloat64{Float64:src.PlayTime, Valid:true}
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

// 变换为'AlbumTrackDto'
func (owner *AlbumTrackEntity) ToDto() AlbumTrackDto {
    var tar AlbumTrackDto
    tar.AlbumId = owner.AlbumId.Int64
    tar.TrackNo = owner.TrackNo.Int64
    tar.TrackName = owner.TrackName.String
    tar.PlayTime = owner.PlayTime.Float64
    tar.CreateAuthor = owner.CreateAuthor.String
    tar.CreateDatetime = owner.CreateDatetime.String
    tar.UpdateAuthor = owner.UpdateAuthor.String
    tar.UpdateDatetime = owner.UpdateDatetime.String
    return tar
}

// 创建行数据映射列表
func (owner *AlbumTrackEntity) Mapper(entity interface{}) []interface{} {
    e := entity.(*AlbumTrackEntity)
    return []interface{}{ &e.AlbumId,&e.TrackNo,&e.TrackName,&e.PlayTime,&e.CreateAuthor,&e.CreateDatetime,&e.UpdateAuthor,&e.UpdateDatetime}
}

// 查询
func (owner *AlbumTrackEntity) Retrieve(ctx OrmContext, orderBy ...OrderByCondition) []AlbumTrackEntity {
    sqlText, sqlParams := GetDao().BuildSqlSelect(owner, orderBy)
    var nl []AlbumTrackEntity
    GetDao().Retrieve(ctx, sqlText, sqlParams[:]...).Mapping(&nl, owner.Mapper)
    return nl
}

// 查询(使用map类型参数查询)
func (owner *AlbumTrackEntity) RetrieveByMap(ctx OrmContext, param map[string]interface{}, orderBy ...OrderByCondition) []AlbumTrackEntity {
    return owner.FromMap(param).Retrieve(ctx,orderBy[:]...)
}

// 统计
func (owner *AlbumTrackEntity) Count(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlCount(owner)
    return GetDao().Count(ctx, sqlText, sqlParams[:]...)
}

// 登录
func (owner *AlbumTrackEntity) Insert(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlInsert(owner)
    return GetDao().Insert(ctx, sqlText, sqlParams[:]...)
}

// 更新
func (owner *AlbumTrackEntity) Update(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlUpdate(owner)
    return GetDao().Update(ctx, sqlText, sqlParams[:]...)
}

// 删除
func (owner *AlbumTrackEntity) Delete(ctx OrmContext) int64 {
    sqlText, sqlParams := GetDao().BuildSqlDelete(owner)
    return GetDao().Delete(ctx, sqlText, sqlParams[:]...)
}







