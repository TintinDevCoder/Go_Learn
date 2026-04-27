# Go语言快速复习指南

基于SGG模块的结构化学习内容，本指南涵盖Go语言从基础到高级的核心概念，帮助你快速回顾。

## 基础概念

### 1. 数据类型
- **基本类型**：`int`, `int8`-`int64`, `uint`, `uint8`-`uint64`, `float32`, `float64`, `bool`, `byte`, `rune`, `string`
- **类型转换**：Go要求显式转换，使用`T(value)`形式
- **默认值**：数值类型为0，布尔类型为false，字符串为""
- **查看类型和大小**：`fmt.Printf("%T", var)`，`unsafe.Sizeof(var)`

**参考文件**：`SGG/basic/basic_data_type.go`

### 2. 变量与常量
- **变量声明**：`var name type`，`name := value`（短声明）
- **作用域**：局部变量、全局变量
- **常量**：`const`关键字，`iota`生成枚举
- **命名规范**：驼峰式，首字母大写表示导出

**参考文件**：`SGG/basic/bian_liang.go`, `SGG/basic/const.go`

### 3. 运算符
- **算术运算符**：`+`, `-`, `*`, `/`, `%`, `++`, `--`
- **关系运算符**：`==`, `!=`, `>`, `<`, `>=`, `<=`
- **逻辑运算符**：`&&`, `||`, `!`
- **位运算符**：`&`, `|`, `^`, `<<`, `>>`
- **赋值运算符**：`=`, `+=`, `-=`, `*=`, `/=`, `%=`, `<<=`, `>>=`, `&=`, `|=`, `^=`

**参考文件**：`SGG/basic/operator.go`

### 4. 控制流
- **条件语句**：`if`, `else if`, `else`
- **循环语句**：`for`（三种形式：基本for、while式for、range循环）
- **switch语句**：基本switch、类型switch、表达式switch
- **跳转语句**：`break`（带标签）、`continue`、`goto`（不推荐）

**参考文件**：`SGG/advanced/for.go`, `SGG/advanced/switch.go`, `SGG/advanced/break.go`, `SGG/advanced/goto.go`

### 5. 函数
- **函数定义**：`func name(params) returnType { ... }`
- **多返回值**：Go函数支持返回多个值
- **匿名函数与闭包**：函数作为一等公民
- **递归函数**：函数调用自身
- **高阶函数**：函数作为参数或返回值

**参考文件**：`SGG/advanced/func/` 目录下的文件

## 高级主题

### 1. 指针
- **指针概念**：存储内存地址的变量，`*T`表示指向T类型的指针
- **取地址与解引用**：`&`获取地址，`*`访问指向的值
- **指针与函数**：通过指针修改函数外部的变量

**参考文件**：`SGG/advanced/pointer.go`

### 2. 数组、切片与映射
- **数组**：固定长度的同类型元素序列，值类型
- **切片**：动态数组，引用类型，基于数组或`make()`创建
- **切片操作**：`len()`, `cap()`, `append()`, `copy()`, 切片表达式
- **映射**：键值对集合，`map[K]V`类型
- **映射操作**：增删改查，遍历，并发安全考虑

**参考文件**：`SGG/advanced/arrayAndslice/` 目录，`SGG/advanced/map.go`

### 3. 字符串处理
- **字符串特性**：不可变，UTF-8编码
- **常用操作**：拼接、分割、查找、替换、转换
- **字符串与切片**：`string`与`[]byte`/`[]rune`相互转换

**参考文件**：`SGG/advanced/string.go`, `SGG/advanced/arrayAndslice/stringslice.go`

### 4. 结构体与方法
- **结构体定义**：`type Name struct { fields }`
- **方法定义**：`func (receiver Type) methodName(params) returnType`
- **值接收者 vs 指针接收者**：决定方法是否能修改接收者
- **工厂模式**：使用函数创建结构体实例

**参考文件**：`SGG/advanced/struct/` 目录

### 5. 接口与多态
- **接口定义**：`type InterfaceName interface { methodSignatures }`
- **接口实现**：隐式实现，只要类型拥有接口所有方法
- **空接口**：`interface{}`可存储任何类型
- **类型断言**：`value, ok := interfaceVar.(ConcreteType)`
- **接口嵌套**：接口可以组合其他接口

**参考文件**：`SGG/advanced/struct/三大特性/interface.go`

### 6. 错误处理
- **error接口**：`type error interface { Error() string }`
- **错误创建**：`errors.New()`, `fmt.Errorf()`
- **错误检查**：`if err != nil { ... }`
- **自定义错误**：实现`Error()`方法
- **panic与recover**：`panic()`触发异常，`recover()`捕获异常

**参考文件**：`SGG/advanced/err_handle.go`

### 7. 并发编程
- **goroutine**：轻量级线程，`go function()`启动
- **channel**：`chan T`类型，用于goroutine间通信
- **channel操作**：发送`ch <- value`，接收`value := <-ch`
- **select语句**：多路channel操作
- **同步原语**：`sync.Mutex`, `sync.WaitGroup`, `sync.Once`
- **MPG模型**：M（内核线程）、P（处理器）、G（goroutine）

**参考文件**：`SGG/advanced/goroutine_and_channel/` 目录

### 8. 标准库常用功能
- **时间处理**：`time`包，时间获取、格式化、计算
- **JSON处理**：`encoding/json`包，`Marshal()`和`Unmarshal()`
- **文件操作**：`os`和`io`包，文件读写、创建、删除
- **命令行参数**：`os.Args`，`flag`包
- **内置函数**：`len()`, `cap()`, `make()`, `new()`, `append()`, `copy()`, `close()`

**参考文件**：`SGG/advanced/timefunc/`, `SGG/advanced/json.go`, `SGG/advanced/file/`, `SGG/advanced/args.go`, `SGG/advanced/build-in_func.go`

## 测试与调试

### 1. 单元测试
- **测试文件**：`*_test.go`命名
- **测试函数**：`func TestXxx(t *testing.T)`
- **运行测试**：`go test -v ./...`

**参考文件**：`SGG/advanced/union_test/` 目录

### 2. 性能测试
- **基准测试**：`func BenchmarkXxx(b *testing.B)`
- **运行基准**：`go test -bench=.`

## 项目结构与模块

### 1. Go模块
- **模块定义**：`go.mod`文件，`module moduleName`
- **依赖管理**：`go get`, `go mod tidy`
- **模块版本**：语义化版本控制

### 2. 包管理
- **包定义**：目录名作为包名，`package name`
- **导入路径**：`import "module/path"`
- **可见性**：首字母大写表示导出（公开），小写表示私有

**参考文件**：`SGG/packageRules/` 目录

## 实践练习

### 1. 家庭记账软件
- **功能**：收支记录、查询、修改、删除
- **技术点**：结构体、切片、文件操作、用户交互

**参考文件**：`SGG/test/FamilyAccount.go`

### 2. 客户关系管理系统
- **功能**：客户信息增删改查、排序、统计
- **技术点**：结构体切片、排序接口、文件持久化

**参考文件**：`SGG/test/CustomerRelationshipManagement.go`

## 数据库与网络

### 1. Redis操作
- **连接Redis**：`github.com/go-redis/redis`客户端
- **基本操作**：字符串、哈希、列表、集合、有序集合
- **连接池**：配置和管理连接

**参考文件**：`SGG/advanced/redis/` 目录

### 2. MySQL操作
- **连接MySQL**：`database/sql`和MySQL驱动
- **CRUD操作**：查询、插入、更新、删除
- **事务处理**：`Begin()`, `Commit()`, `Rollback()`

**参考文件**：`SGG/advanced/mysql/` 目录

### 3. TCP网络编程
- **TCP服务器**：`net.Listen()`监听端口，`Accept()`接受连接
- **TCP客户端**：`net.Dial()`建立连接
- **并发服务器**：为每个连接启动goroutine

**参考文件**：`SGG/advanced/tcp_socket/` 目录

## 高级特性

### 1. 反射
- **反射基础**：`reflect.TypeOf()`, `reflect.ValueOf()`
- **类型检查**：`Kind()`方法判断类型种类
- **值操作**：获取/设置字段值，调用方法

**参考文件**：`SGG/advanced/reflect/` 目录

### 2. 上下文
- **上下文作用**：传递请求范围的值、取消信号、超时
- **创建上下文**：`context.Background()`, `context.WithCancel()`
- **使用上下文**：在goroutine间传递，处理超时和取消

## 最佳实践

### 1. 代码风格
- **格式化**：`go fmt`自动格式化
- **命名约定**：驼峰命名，简洁明确
- **错误处理**：尽早处理错误，避免嵌套过深

### 2. 性能优化
- **内存分配**：减少不必要的内存分配，重用对象
- **并发模式**：合理使用goroutine和channel，避免数据竞争
- ** profiling**：使用`pprof`进行性能分析

### 3. 调试技巧
- **打印调试**：`fmt.Printf()`, `log.Println()`
- **调试器**：使用Delve进行调试
- **单元测试**：编写可测试的代码，使用表驱动测试

## 快速复习清单

1. **基础语法**：变量、常量、数据类型、控制流
2. **复合类型**：数组、切片、映射、结构体
3. **函数与方法**：函数定义、多返回值、方法接收者
4. **接口与多态**：接口定义、隐式实现、类型断言
5. **并发编程**：goroutine、channel、select、同步
6. **错误处理**：error接口、panic/recover、自定义错误
7. **标准库**：时间、JSON、文件、网络
8. **测试**：单元测试、基准测试
9. **工具**：go mod、go test、go fmt

## 下一步建议

1. **运行示例**：进入SGG目录，运行各个示例文件
2. **修改实验**：尝试修改代码，观察不同行为
3. **解决问题**：尝试LeetCode模块中的算法问题
4. **构建项目**：使用所学知识构建小型应用

---

*本指南基于SGG模块内容生成，最后更新：2026-04-17*  
*SGG模块路径：`D:\dsy\Go_Learn\SGG\`*