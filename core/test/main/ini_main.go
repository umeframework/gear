package main

import (
    "fmt"
)
import (_ "github.com/go-sql-driver/mysql"
    "github.com/umeframework/gear/core"
    "runtime/debug"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("panic recover! p: %v", r)
            debug.PrintStack()
        }
    }()
    // 测试
    fmt.Println("Test start")

    testIni()

    fmt.Println("Test finished")

}

func testIni() {
    cfg := core.NewIniConfig("src/github.com/umeframework/gear/core/test/main/sample.ini")
    fmt.Println(cfg.ParagraphSet())
    fmt.Println(cfg.Paragraph("JDBC basic"))
}



