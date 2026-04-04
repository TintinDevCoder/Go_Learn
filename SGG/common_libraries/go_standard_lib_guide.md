# Go 标准库学习指南

## 概述

Go 标准库是 Go 语言的核心组成部分，提供了丰富的基础功能，包括 I/O 操作、网络编程、并发处理、数据编码等。掌握标准库是成为 Go 开发者的关键。

### 标准库特点
- **高质量**: 经过严格测试，性能优异
- **一致性**: 统一的 API 设计风格
- **文档齐全**: 每个包都有详细的文档和示例
- **无需第三方依赖**: 开箱即用

## 核心包分类

### 1. 输入输出 (I/O)
- `fmt`: 格式化 I/O
- `io`: 基础 I/O 接口
- `os`: 操作系统功能
- `bufio`: 缓冲 I/O
- `ioutil`: I/O 工具函数

### 2. 文本处理
- `strings`: 字符串操作
- `strconv`: 字符串转换
- `regexp`: 正则表达式
- `unicode`: Unicode 相关功能

### 3. 数据编码
- `encoding/json`: JSON 编码解码
- `encoding/xml`: XML 编码解码
- `encoding/base64`: Base64 编码
- `encoding/csv`: CSV 文件处理

### 4. 网络编程
- `net/http`: HTTP 客户端和服务器
- `net`: 网络 I/O
- `net/url`: URL 解析
- `net/mail`: 邮件处理

### 5. 并发编程
- `sync`: 同步原语
- `sync/atomic`: 原子操作
- `context`: 上下文管理
- `time`: 时间操作

### 6. 数据处理
- `sort`: 排序
- `container`: 容器数据结构
- `math`: 数学函数
- `crypto`: 加密算法

## 详细学习

### 1. fmt - 格式化 I/O

#### 主要功能
- 格式化输出到控制台
- 格式化字符串
- 扫描输入

#### 常用函数
```go
// 输出
fmt.Print("Hello")
fmt.Println("World")
fmt.Printf("Name: %s, Age: %d\n", name, age)

// 格式化字符串
s := fmt.Sprintf("Name: %s", name)

// 输入
fmt.Scan(&name)
fmt.Scanf("%s %d", &name, &age)
fmt.Scanln(&name)

// 错误输出
fmt.Fprint(os.Stderr, "Error")
fmt.Fprintf(os.Stderr, "Error: %v", err)
```

#### 格式化动词
- `%v`: 默认格式
- `%+v`: 输出结构体字段名
- `%#v`: Go 语法表示
- `%T`: 类型
- `%d`: 十进制整数
- `%f`: 浮点数
- `%s`: 字符串
- `%t`: 布尔值
- `%p`: 指针

### 2. strings - 字符串操作

#### 常用函数
```go
import "strings"

// 判断
strings.Contains("hello world", "world")  // true
strings.HasPrefix("hello", "he")         // true
strings.HasSuffix("hello", "lo")         // true

// 查找
strings.Index("hello", "l")              // 2
strings.LastIndex("hello", "l")          // 3

// 分割和连接
strings.Split("a,b,c", ",")              // ["a", "b", "c"]
strings.Join([]string{"a", "b"}, ",")    // "a,b"

// 替换
strings.Replace("hello", "l", "L", 1)    // "heLlo"
strings.ReplaceAll("hello", "l", "L")    // "heLLo"

// 大小写
strings.ToUpper("hello")                 // "HELLO"
strings.ToLower("HELLO")                 // "hello"

// 修剪
strings.Trim(" hello ", " ")             // "hello"
strings.TrimSpace(" hello ")             // "hello"
strings.TrimLeft("hello", "h")           // "ello"
strings.TrimRight("hello", "o")          // "hell"

// 重复
strings.Repeat("ha", 3)                  // "hahaha"

// 比较
strings.Compare("a", "b")                // -1
strings.EqualFold("Go", "go")            // true (不区分大小写)
```

#### 字符串构建器
```go
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteByte(' ')
builder.WriteString("World")
result := builder.String()  // "Hello World"
```

### 3. strconv - 字符串转换

#### 常用函数
```go
import "strconv"

// 字符串转整数
i, err := strconv.Atoi("123")          // 123
i64, err := strconv.ParseInt("123", 10, 64)

// 整数转字符串
s := strconv.Itoa(123)                 // "123"
s = strconv.FormatInt(123, 10)         // "123"

// 字符串转浮点数
f, err := strconv.ParseFloat("3.14", 64)

// 浮点数转字符串
s = strconv.FormatFloat(3.14, 'f', 2, 64)  // "3.14"

// 布尔值转换
b, err := strconv.ParseBool("true")
s = strconv.FormatBool(true)           // "true"

// 引用转换
s = strconv.Quote("hello")             // `"hello"`
s = strconv.QuoteToASCII("hello")      // `"hello"`
```

### 4. os - 操作系统功能

#### 文件和目录操作
```go
import "os"

// 文件操作
file, err := os.Open("file.txt")
defer file.Close()

file, err = os.Create("newfile.txt")
file.WriteString("content")
file.Close()

// 文件信息
info, err := os.Stat("file.txt")
info.Name()      // 文件名
info.Size()      // 文件大小
info.ModTime()   // 修改时间
info.IsDir()     // 是否是目录
info.Mode()      // 权限模式

// 目录操作
os.Mkdir("dir", 0755)
os.MkdirAll("path/to/dir", 0755)

files, err := os.ReadDir(".")
for _, file := range files {
    fmt.Println(file.Name(), file.IsDir())
}

// 路径操作
os.Remove("file.txt")
os.RemoveAll("dir")

// 重命名
os.Rename("old.txt", "new.txt")

// 环境变量
os.Setenv("KEY", "value")
value := os.Getenv("KEY")
envs := os.Environ()

// 工作目录
wd, _ := os.Getwd()
os.Chdir("/tmp")

// 进程信息
pid := os.Getpid()
hostname, _ := os.Hostname()
os.Exit(1)  // 退出程序
```

### 5. io - 基础 I/O 接口

#### 核心接口
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

#### 常用函数
```go
import "io"

// 复制数据
n, err := io.Copy(dstWriter, srcReader)

// 复制指定字节数
n, err = io.CopyN(dstWriter, srcReader, 1024)

// 读取全部数据
data, err := io.ReadAll(reader)

// 丢弃数据
n, err = io.Discard.Write(data)

// 管道
reader, writer := io.Pipe()
go func() {
    writer.Write([]byte("data"))
    writer.Close()
}()
data, _ := io.ReadAll(reader)

// 多写器
mw := io.MultiWriter(writer1, writer2)
mw.Write([]byte("hello"))

// 多读器
mr := io.MultiReader(reader1, reader2)
io.ReadAll(mr)

// 限制读取
lr := io.LimitReader(reader, 1024)  // 最多读取1024字节
```

### 6. bufio - 缓冲 I/O

#### 缓冲读写
```go
import "bufio"

// 缓冲读取器
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n')
line, err = reader.ReadBytes('\n')
data := make([]byte, 1024)
n, err := reader.Read(data)

// 按行读取
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // 处理每一行
}
if err := scanner.Err(); err != nil {
    // 处理错误
}

// 缓冲写入器
writer := bufio.NewWriter(file)
writer.WriteString("hello")
writer.WriteByte('\n')
writer.Write([]byte("world"))
writer.Flush()  // 重要：将缓冲数据写入底层写入器

// 读取器/写入器
rw := bufio.NewReadWriter(reader, writer)
```

### 7. ioutil - I/O 工具函数

#### 常用函数（注：Go 1.16后部分函数移到os包）
```go
import "io/ioutil"

// 读取文件全部内容
data, err := ioutil.ReadFile("file.txt")

// 写入文件
err = ioutil.WriteFile("file.txt", data, 0644)

// 读取目录
files, err := ioutil.ReadDir(".")

// 临时文件和目录
tempFile, err := ioutil.TempFile("", "prefix")
tempDir, err := ioutil.TempDir("", "prefix")

// 废弃但仍然可用（Go 1.16+）
import "os"
data, err = os.ReadFile("file.txt")          // 替代 ioutil.ReadFile
err = os.WriteFile("file.txt", data, 0644)   // 替代 ioutil.WriteFile
files, err = os.ReadDir(".")                 // 替代 ioutil.ReadDir
```

### 8. encoding/json - JSON 编码解码

#### 结构体标签
```go
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email,omitempty"`  // 空值忽略
    Secret  string `json:"-"`                // 忽略字段
    Created time.Time `json:"created_at"`
}
```

#### 编码（序列化）
```go
import "encoding/json"

person := Person{Name: "John", Age: 30}

// 编码为JSON
data, err := json.Marshal(person)
// data: {"name":"John","age":30,"created_at":"2023-..."}

// 格式化输出
data, err = json.MarshalIndent(person, "", "  ")
/*
{
  "name": "John",
  "age": 30,
  "created_at": "2023-..."
}
*/

// 编码到Writer
encoder := json.NewEncoder(writer)
encoder.SetIndent("", "  ")
err = encoder.Encode(person)
```

#### 解码（反序列化）
```go
jsonStr := `{"name":"John","age":30}`

// 解码JSON
var person Person
err := json.Unmarshal([]byte(jsonStr), &person)

// 从Reader解码
decoder := json.NewDecoder(reader)
err = decoder.Decode(&person)

// 解码未知结构
var data map[string]interface{}
err = json.Unmarshal([]byte(jsonStr), &data)
name := data["name"].(string)

// 流式解码
decoder := json.NewDecoder(reader)
for decoder.More() {
    var item Item
    err := decoder.Decode(&item)
    if err != nil {
        break
    }
    // 处理item
}
```

### 9. net/http - HTTP 客户端和服务器

#### HTTP 服务器
```go
import "net/http"

// 简单服务器
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
})

http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})

// 启动服务器
http.ListenAndServe(":8080", nil)

// 使用中间件
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// 文件服务器
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

// HTTPS服务器
http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
```

#### HTTP 客户端
```go
// GET请求
resp, err := http.Get("https://api.example.com/data")
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)

// 带参数的GET请求
req, err := http.NewRequest("GET", "https://api.example.com", nil)
q := req.URL.Query()
q.Add("key", "value")
req.URL.RawQuery = q.Encode()

// POST请求
data := strings.NewReader(`{"name":"John"}`)
resp, err := http.Post("https://api.example.com", "application/json", data)

// POST表单
resp, err := http.PostForm("https://api.example.com",
    url.Values{"key": {"value"}})

// 自定义请求
req, err := http.NewRequest("POST", "https://api.example.com", data)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token")

client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Do(req)

// 下载文件
resp, err := http.Get("https://example.com/file.zip")
defer resp.Body.Close()
file, err := os.Create("file.zip")
defer file.Close()
io.Copy(file, resp.Body)
```

### 10. sync - 同步原语

#### 互斥锁
```go
import "sync"

var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}

// 读写锁
var rwmu sync.RWMutex

func read() {
    rwmu.RLock()
    defer rwmu.RUnlock()
    // 读取操作
}

func write() {
    rwmu.Lock()
    defer rwmu.Unlock()
    // 写入操作
}
```

#### 等待组
```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        // 执行任务
    }(i)
}

wg.Wait()  // 等待所有goroutine完成
```

#### 一次性执行
```go
var once sync.Once

func initialize() {
    once.Do(func() {
        // 只执行一次
    })
}
```

#### 条件变量
```go
var cond = sync.NewCond(&sync.Mutex{})
var queue []string

// 生产者
cond.L.Lock()
queue = append(queue, "item")
cond.Signal()  // 唤醒一个等待者
cond.L.Unlock()

// 消费者
cond.L.Lock()
for len(queue) == 0 {
    cond.Wait()  // 等待信号
}
item := queue[0]
queue = queue[1:]
cond.L.Unlock()
```

#### 映射安全版本
```go
var m sync.Map

// 存储
m.Store("key", "value")

// 加载
value, ok := m.Load("key")

// 删除
m.Delete("key")

// 遍历
m.Range(func(key, value interface{}) bool {
    fmt.Println(key, value)
    return true  // 继续遍历
})

// 加载或存储
actual, loaded := m.LoadOrStore("key", "value")

// 加载并删除
value, loaded := m.LoadAndDelete("key")
```

### 11. context - 上下文管理

#### 创建上下文
```go
import "context"

// 空上下文
ctx := context.Background()

// 带取消的上下文
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // 释放资源

// 带超时的上下文
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 带截止时间的上下文
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// 带值的上下文
ctx = context.WithValue(context.Background(), "key", "value")
```

#### 使用上下文
```go
func worker(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("任务取消:", ctx.Err())
        return
    case <-time.After(10 * time.Second):
        fmt.Println("任务完成")
    }
}

// HTTP请求中使用
req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

// 数据库操作中使用
db.QueryContext(ctx, "SELECT ...")
```

### 12. time - 时间操作

#### 时间创建和格式化
```go
import "time"

// 当前时间
now := time.Now()

// 创建特定时间
t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

// 解析时间
t, err := time.Parse("2006-01-02", "2023-01-01")
t, err = time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")

// 格式化时间
fmt.Println(t.Format("2006-01-02 15:04:05"))
fmt.Println(t.Format(time.RFC3339))

// 时间戳
timestamp := t.Unix()      // 秒
nanos := t.UnixNano()      // 纳秒
t = time.Unix(timestamp, 0)
```

#### 时间计算
```go
// 加减时间
later := now.Add(2 * time.Hour)
earlier := now.Add(-30 * time.Minute)

// 时间差
duration := later.Sub(earlier)

// 比较时间
if now.After(earlier) {
    // now在earlier之后
}
if now.Before(later) {
    // now在later之前
}

// 睡眠
time.Sleep(2 * time.Second)

// 定时器
timer := time.NewTimer(2 * time.Second)
<-timer.C  // 等待定时器触发

// 打点器
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for t := range ticker.C {
    fmt.Println("Tick at", t)
}

// 超时控制
select {
case <-time.After(5 * time.Second):
    fmt.Println("超时")
case result := <-ch:
    fmt.Println("收到结果:", result)
}
```

### 13. sort - 排序

#### 基本排序
```go
import "sort"

// 切片排序
ints := []int{3, 1, 4, 1, 5}
sort.Ints(ints)

strings := []string{"c", "a", "b"}
sort.Strings(strings)

floats := []float64{3.14, 1.41, 2.71}
sort.Float64s(floats)

// 检查是否已排序
sort.IntsAreSorted(ints)

// 自定义排序
people := []struct {
    Name string
    Age  int
}{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 20},
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})

// 稳定排序
sort.SliceStable(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})

// 搜索
index := sort.SearchInts(ints, 4)  // 在已排序切片中搜索
```

### 14. container - 容器数据结构

#### 堆（优先队列）
```go
import "container/heap"

// 实现heap.Interface
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// 使用
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
min := heap.Pop(h).(int)
```

#### 链表
```go
import "container/list"

l := list.New()

// 添加元素
l.PushBack("back")
l.PushFront("front")

// 遍历
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}

// 删除元素
l.Remove(e)
```

#### 环
```go
import "container/ring"

r := ring.New(5)

// 初始化
for i := 0; i < r.Len(); i++ {
    r.Value = i
    r = r.Next()
}

// 遍历
r.Do(func(x interface{}) {
    fmt.Println(x)
})
```

### 15. math - 数学函数

#### 常用函数
```go
import "math"

// 基本运算
math.Abs(-3.14)      // 绝对值
math.Ceil(3.14)      // 向上取整
math.Floor(3.14)     // 向下取整
math.Round(3.14)     // 四舍五入
math.Trunc(3.14)     // 截断小数部分

// 指数和对数
math.Exp(1)          // e的x次方
math.Log(10)         // 自然对数
math.Log10(100)      // 以10为底的对数
math.Pow(2, 3)       // 2的3次方
math.Sqrt(16)        // 平方根
math.Cbrt(27)        // 立方根

// 三角函数
math.Sin(math.Pi/2)  // 正弦
math.Cos(0)          // 余弦
math.Tan(math.Pi/4)  // 正切

// 最大值/最小值
math.Max(1, 2)       // 2
math.Min(1, 2)       // 1

// 随机数
import "math/rand"
rand.Seed(time.Now().UnixNano())
rand.Int()           // 随机整数
rand.Intn(100)       // 0-99的随机整数
rand.Float64()       // 0.0-1.0的随机浮点数
```

### 16. crypto - 加密算法

#### 哈希函数
```go
import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

data := []byte("hello")

// MD5
hash := md5.Sum(data)
fmt.Printf("MD5: %x\n", hash)

// SHA1
hash = sha1.Sum(data)
fmt.Printf("SHA1: %x\n", hash)

// SHA256
hash = sha256.Sum256(data)
fmt.Printf("SHA256: %x\n", hash)

// 使用io.Writer
h := sha256.New()
h.Write([]byte("hello"))
h.Write([]byte("world"))
hash := h.Sum(nil)
hexHash := hex.EncodeToString(hash)
```

## 最佳实践

### 1. 错误处理
```go
// 使用errors.New创建错误
import "errors"
var ErrNotFound = errors.New("not found")

// 错误包装
import "fmt"
err := fmt.Errorf("operation failed: %w", originalErr)

// 错误检查
if errors.Is(err, ErrNotFound) {
    // 处理特定错误
}

// 错误类型断言
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("路径错误:", pathErr.Path)
}
```

### 2. 性能优化
- 使用`strings.Builder`构建字符串
- 使用缓冲I/O处理大文件
- 避免频繁的内存分配
- 使用连接池（数据库、HTTP客户端）

### 3. 并发安全
- 使用sync包提供的同步原语
- 避免数据竞争
- 合理使用goroutine
- 使用context管理goroutine生命周期

### 4. 代码组织
- 合理使用包组织代码
- 遵循Go命名约定
- 编写清晰的文档注释
- 保持函数简短专注

## 实战示例

### Web服务
```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    http.HandleFunc("/users", usersHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getUsers(w, r)
    case "POST":
        createUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()},
        {ID: 2, Name: "Bob", Email: "bob@example.com", CreatedAt: time.Now()},
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    user.ID = 3
    user.CreatedAt = time.Now()
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

### 并发数据处理
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)  // 模拟工作
        results <- job * 2
    }
}

func main() {
    const numJobs = 10
    const numWorkers = 3
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    var wg sync.WaitGroup
    
    // 启动worker
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // 发送jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 等待所有worker完成
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // 收集结果
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}
```

## 总结

Go标准库提供了丰富而强大的功能，涵盖了日常开发的大多数需求。掌握这些标准库是成为高效Go开发者的基础。

建议的学习路径：
1. 先从`fmt`、`strings`、`strconv`等基础包开始
2. 学习`io`、`os`、`bufio`等I/O相关包
3. 掌握`encoding/json`、`net/http`等常用包
4. 学习`sync`、`context`、`time`等并发相关包
5. 根据需要学习其他专业领域的包

通过实践和阅读官方文档，可以深入理解每个包的用法和最佳实践。标准库的设计体现了Go语言的哲学：简单、高效、实用。