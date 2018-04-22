package core

import (
    "io"
    "strings"
    "strconv"
    "os"
    "bufio"
    "sync"
)

// INI文件内容存贮结构
type IniConfig struct {
    // 文件名
    file       string
    // 段落内容
    paragraphs map[string]map[string]string
    sync       *sync.Mutex
}

// 创建配置对象
func NewIniConfig(file string) *IniConfig {
    cfg := new(IniConfig)
    cfg.paragraphs = make(map[string]map[string]string)
    cfg.file = file
    cfg.sync = &sync.Mutex {}
    cfg.Load()
    return cfg
}

// 内容加载
func (this *IniConfig) Load() {
    this.sync.Lock()
    defer this.sync.Unlock()

    file, err := os.Open(this.file)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    //var paragraph map[string]map[string]string
    var paragraph string
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
        switch {
        case len(lineNS) == 0:
        case lineNS[0] == '[' && lineNS[len(lineNS)-1] == ']':
            paragraph = strings.TrimSpace(lineNS[1 : len(lineNS)-1])

            _,exist := this.paragraphs[paragraph]
            if !exist {
                this.paragraphs[paragraph] = make(map[string]string)
            }
        default:
            i := strings.IndexAny(lineNS, "=")
            key := strings.TrimSpace(lineNS[0:i])
            value := strings.TrimSpace(lineNS[i+1 :])
            this.paragraphs[paragraph][key] = value
        }
    }
}

// 返回段落集合
func (this *IniConfig) ParagraphSet() []string {
    var keys []string
    for key := range this.paragraphs {
        keys = append(keys, key)
    }
    return keys
}

// 返回段落内容
func (this *IniConfig) Paragraph(paragraph string) map[string]string {
    _,exist := this.paragraphs[paragraph]
    if exist {
        return this.paragraphs[paragraph]
    }
    return nil
}

// 读取文本值
func (this *IniConfig) Get(paragraph string, key string) string {
    return this.paragraphs[paragraph][key]
}

// 设置文本值（暂不支持文件保存）
func (this *IniConfig) Set(paragraph string, key string, value string) {
    this.sync.Lock()
    defer this.sync.Unlock()
    _,exist := this.paragraphs[paragraph]
    if !exist {
        this.paragraphs[paragraph] = make(map[string]string)
    }
    this.paragraphs[paragraph][key] = value
}

// 移除（暂不支持文件保存）
func (this *IniConfig) Remove(paragraph string, key string) {
    this.sync.Lock()
    defer this.sync.Unlock()
    _,exist := this.paragraphs[paragraph]
    if exist {
        delete(this.paragraphs[paragraph], key)
    }
}

// 读取文本值列表
func (this *IniConfig) GetList(paragraph string, key string) []string {
    value :=  this.Get(paragraph, key)
    if value == "" {
        return nil
    }
    values := strings.Split(value, ",")
    return values
}

// 读取整数值
func (this *IniConfig) GetInt(paragraph string, key string) (int, error) {
    s := this.Get(paragraph, key)
    return strconv.Atoi(s)
}

// 读取整数值列表
func (this *IniConfig) GetIntList(paragraph string, key string) []int {
    strValues := this.GetList(paragraph, key)
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


