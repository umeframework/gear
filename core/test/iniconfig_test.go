package test

import (
    "fmt"
    "testing"
)
import (. "github.com/umeframework/gear/core")

func TestIni(t *testing.T) {
    cfg := NewIniConfig("sample.ini")
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

