package main

import (
    "fmt"
    "database/sql"
    "github.com/umeframework/gear/core"
    "github.com/umeframework/gear/orm"
)
import (_ "github.com/go-sql-driver/mysql")
import (. "github.com/umeframework/gear/test/dto")

func main() {
    cfg := core.NewConfigProperties("src/github.com/umeframework/gear/config/gear.properties")
    driver := cfg.Get("ume.gdbc.driver")
    url := cfg.Get("ume.gdbc.url")
    username := cfg.Get("ume.gdbc.username")
    password := cfg.Get("ume.gdbc.password")

    ctx := orm.GetOrmContext(driver, username + ":" + password + "@" + url)
    defer ctx.Close()

    // 测试
    TestSql()
    //TestQuery(ctx)
    TestUpdate(ctx)
    fmt.Println("End test")

}



// 测试
func TestSql() {
    e := AlbumEntity{}
    em := orm.GetEntityMetadata(&e)
    fmt.Println(em.SQLSelectDefault)
    fmt.Println(em.SQLSelectOneDefault)
    fmt.Println(em.SQLSelectCountDefault)
    fmt.Println(em.SQLInsertDefault)
    fmt.Println(em.SQLUpdateDefault)
    fmt.Println(em.SQLDeleteDefault)

}

func TestQuery(ctx orm.OrmContext) {
    e := &AlbumEntity{}
    fmt.Println("@Entity查询")
    fmt.Println("Total:", e.Count(ctx))
    list := e.Retrieve(ctx)
    for _, e := range list {
        fmt.Println(e)
    }
}

// 测试
func TestUpdate(ctx orm.OrmContext) {
    fmt.Println("查询Id为999的记录集")
    // 定义查询参数，AlbumEntity的Id为999，其余字段为空
    e := new(AlbumEntity)
    e.Id = sql.NullInt64{999, true}

    // 执行查询
    list := e.Retrieve(ctx)
    fmt.Println(list)

    // 判断查询件数
    if len(list) > 0 {
        // 执行删除，删除Id为999的记录
        c := e.Delete(ctx)
        if c > 0 {
            fmt.Println("删除ID为999的记录:", c)
        }
    } else {
        fmt.Println("未查到ID为999的记录")
    }

    // 构建新参数，Id依然为999，Artist为'Sting'
    d := AlbumDto{Id: 999, Title: "Nothing Can Like The Sun", Artist: "Sting"}
    // 执行插入
    c := e.FromDto(d).Insert(ctx)
    if c != 0 {
        fmt.Println("插入记录:", c)
    }
    e.CreateAuthor=sql.NullString{"TestUpdater",true}
    fmt.Println(e.Update(ctx))
    fmt.Println("查询所有记录")
    e = &AlbumEntity{}
    list = e.Retrieve(ctx)
    for _, e := range list {
        fmt.Println(e)
    }
}


