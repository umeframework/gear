package core

import (
    "io"
    "strings"
    "strconv"
    "os"
    "bufio"
    "sync"
)

// 属性文件内容存贮结构
type PropertyConfig struct {
    // 文件名
    file     string
    // 内容
    contents map[string]string
    sync     *sync.Mutex
}

// 创建配置对象
func NewPropertyConfig(file string) *PropertyConfig {
    cfg := new(PropertyConfig)
    cfg.contents = make(map[string]string)
    cfg.file = file
    cfg.sync = &sync.Mutex {}
    cfg.Load()
    return cfg
}

// 内容加载
func (this *PropertyConfig) Load() {
    this.sync.Lock()
    defer this.sync.Unlock()

    file, err := os.Open(this.file)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    buf := bufio.NewReader(file)
    for {
        line, err := buf.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                panic(err)
            }
            if len(line) == 0 {
                break
            }
        }
        lineNS := strings.TrimSpace(line)
        if strings.HasPrefix(lineNS, "#") || lineNS == "" {
            continue
        }

        i := strings.IndexAny(lineNS, "=")
        key := strings.TrimSpace(lineNS[0:i])
        value := strings.TrimSpace(lineNS[i+1 :])
        this.contents[key] = value
    }
}

// 返回Key集合
func (this *PropertyConfig) KeySet() []string {
    var keys []string
    for key := range this.contents {
        keys = append(keys, key)
    }
    return keys
}

// 读取文本值
func (this *PropertyConfig) Get(name string) string {
    return this.contents[name]
}

// 设置文本值（暂不支持文件保存）
func (this *PropertyConfig) Set(key string, value string) {
    this.contents[key] = value
}

// 移除（暂不支持文件保存）
func (this *PropertyConfig) Remove(key string) {
    delete(this.contents, key)
}

// 读取文本值列表
func (this *PropertyConfig) GetList(key string) []string {
    value :=  this.Get(key)
    if value == "" {
        return nil
    }
    values := strings.Split(value, ",")
    return values
}

// 读取整数值
func (this *PropertyConfig) GetInt(key string) (int,error ) {
   s := this.Get(key)
   return strconv.Atoi(s)
}

// 读取整数值列表
func (this *PropertyConfig) GetIntList(key string) ([]int) {
    strValues := this.GetList(key)
    var intValues []int
    for _, s := range strValues {
        v,err:=strconv.Atoi(s)
        if err != nil {
            return nil
        }
        intValues=append(intValues,v)
    }
    return intValues
}


