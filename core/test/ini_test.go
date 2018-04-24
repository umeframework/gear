package test

import (
    "fmt"
)
import (_ "github.com/go-sql-driver/mysql"
    "github.com/umeframework/gear/core"
    "testing"
)

func TestIni(t *testing.T) {
    cfg := core.NewIniConfig("sample.ini")
    fmt.Println(cfg.ParagraphSet())
    fmt.Println(cfg.Paragraph("JDBC basic"))
}



//func main() {
//    defer func() {
//        if r := recover(); r != nil {
//            fmt.Printf("panic recover! p: %v", r)
//            debug.PrintStack()
//        }
//    }()
//    // 测试
//    fmt.Println("Test start")
//
//    testIni()
//
//    fmt.Println("Test finished")
//
//}

