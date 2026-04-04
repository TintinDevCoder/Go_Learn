# Go 操作 MySQL 数据库学习指南

## 概述

Go 语言通过 `database/sql` 标准库提供了一套统一的数据库操作接口，配合特定数据库驱动（如 `github.com/go-sql-driver/mysql`）可以实现对 MySQL 数据库的操作。

### 核心概念
- **database/sql**: Go 标准库中的数据库抽象层
- **驱动**: 特定数据库的实现，如 `mysql`, `postgres`, `sqlite3` 等
- **连接池**: `database/sql` 自动管理数据库连接池
- **预处理语句**: 提高性能和安全性的 SQL 执行方式

## 安装和配置

### 1. 安装 MySQL 驱动
```bash
go get github.com/go-sql-driver/mysql
```

### 2. 导入包
```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```
注意：使用 `_` 导入驱动，只执行驱动的初始化代码，不直接使用驱动包。

## 数据库连接

### 基本连接
```go
// DSN格式: username:password@protocol(address)/dbname?param=value
dsn := "root:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
db, err := sql.Open("mysql", dsn)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 验证连接
err = db.Ping()
if err != nil {
    log.Fatal(err)
}
```

### 连接参数说明
- `charset=utf8mb4`: 字符集，支持完整 Unicode
- `parseTime=True`: 将 DATE/DATETIME 解析为 time.Time
- `loc=Local`: 时区设置
- `timeout`: 连接超时时间
- `readTimeout`: 读超时时间
- `writeTimeout`: 写超时时间

### 连接池配置
```go
db.SetMaxOpenConns(25)      // 最大打开连接数
db.SetMaxIdleConns(10)      // 最大空闲连接数
db.SetConnMaxLifetime(5 * time.Minute)  // 连接最大生存时间
```

## 基本操作

### 创建表
```go
createTableSQL := `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
_, err := db.Exec(createTableSQL)
```

### 插入数据
```go
// 简单插入
result, err := db.Exec("INSERT INTO users (username, email, age) VALUES (?, ?, ?)", 
    "john_doe", "john@example.com", 30)

// 获取插入ID
id, err := result.LastInsertId()

// 获取影响行数
rows, err := result.RowsAffected()
```

### 查询数据
```go
// 查询单行
var user struct {
    ID       int
    Username string
    Email    string
    Age      int
}
err := db.QueryRow("SELECT id, username, email, age FROM users WHERE id = ?", 1).Scan(
    &user.ID, &user.Username, &user.Email, &user.Age)

// 查询多行
rows, err := db.Query("SELECT id, username, email, age FROM users WHERE age > ?", 25)
defer rows.Close()

for rows.Next() {
    err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Age)
    // 处理每行数据
}
err = rows.Err() // 检查迭代过程中的错误
```

### 更新数据
```go
result, err := db.Exec("UPDATE users SET age = ? WHERE id = ?", 31, 1)
rows, err := result.RowsAffected()
```

### 删除数据
```go
result, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
```

## 高级特性

### 预处理语句
```go
// 准备预处理语句
stmt, err := db.Prepare("INSERT INTO users (username, email, age) VALUES (?, ?, ?)")
defer stmt.Close()

// 重复使用
for i := 0; i < 10; i++ {
    _, err := stmt.Exec(fmt.Sprintf("user%d", i), 
        fmt.Sprintf("user%d@example.com", i), 20+i)
}
```

### 事务处理
```go
// 开始事务
tx, err := db.Begin()

// 执行事务操作
_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", 100, 1)
if err != nil {
    tx.Rollback()
    return
}

_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", 100, 2)
if err != nil {
    tx.Rollback()
    return
}

// 提交事务
err = tx.Commit()
```

### 批量操作
```go
// 使用 VALUES 批量插入
stmt := `INSERT INTO users (username, email, age) VALUES `
vals := []interface{}{}

for i := 0; i < 10; i++ {
    stmt += "(?, ?, ?),"
    vals = append(vals, fmt.Sprintf("user%d", i), 
        fmt.Sprintf("user%d@example.com", i), 20+i)
}

stmt = stmt[:len(stmt)-1] // 去掉最后一个逗号
_, err := db.Exec(stmt, vals...)
```

## 数据类型映射

### Go 与 MySQL 类型对应
- `INT` → `int`, `int64`
- `VARCHAR` → `string`
- `TEXT` → `string`
- `DECIMAL` → `string` 或自定义类型
- `DATE`/`DATETIME` → `time.Time` (需要 parseTime=True)
- `BOOLEAN` → `bool`
- `BLOB` → `[]byte`

### NULL 值处理
```go
var name sql.NullString
var age sql.NullInt64
var email *string // 使用指针处理 NULL

err := db.QueryRow("SELECT username, age FROM users WHERE id = ?", 1).Scan(&name, &age)
if name.Valid {
    fmt.Println("Name:", name.String)
}
if age.Valid {
    fmt.Println("Age:", age.Int64)
}
```

## 错误处理

### 常见错误类型
```go
import (
    "github.com/go-sql-driver/mysql"
)

// 检查 MySQL 特定错误
if mysqlErr, ok := err.(*mysql.MySQLError); ok {
    switch mysqlErr.Number {
    case 1062: // 重复键错误
        fmt.Println("Duplicate entry")
    case 1045: // 访问被拒绝
        fmt.Println("Access denied")
    }
}

// 检查没有结果
if err == sql.ErrNoRows {
    fmt.Println("No rows found")
}
```

## 性能优化

### 1. 使用连接池
合理配置连接池参数，避免连接泄漏。

### 2. 预处理语句
对重复执行的 SQL 使用预处理语句。

### 3. 批量操作
批量插入/更新减少网络往返。

### 4. 合理使用事务
事务范围不要过大，尽快提交。

### 5. 索引优化
确保查询使用合适的索引。

### 6. 查询优化
只查询需要的字段，避免 SELECT *。

## 安全注意事项

### 1. SQL 注入防护
总是使用参数化查询（? 占位符），不要拼接 SQL 字符串。

### 2. 密码安全
不要在代码中硬编码密码，使用环境变量或配置文件。

### 3. 连接安全
使用 TLS 加密数据库连接。

### 4. 权限控制
遵循最小权限原则，应用程序账户只拥有必要权限。

### 5. 输入验证
对用户输入进行验证和清理。

## 实践示例

### 用户管理系统
```go
// 用户结构体
type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
}

// 用户数据访问对象
type UserDAO struct {
    db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
    return &UserDAO{db: db}
}

func (dao *UserDAO) Create(user *User) error {
    result, err := dao.db.Exec(
        "INSERT INTO users (username, email, age) VALUES (?, ?, ?)",
        user.Username, user.Email, user.Age)
    if err != nil {
        return err
    }
    user.ID, err = result.LastInsertId()
    return err
}

func (dao *UserDAO) GetByID(id int64) (*User, error) {
    user := &User{}
    err := dao.db.QueryRow(
        "SELECT id, username, email, age, created_at FROM users WHERE id = ?",
        id).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (dao *UserDAO) Update(user *User) error {
    _, err := dao.db.Exec(
        "UPDATE users SET username = ?, email = ?, age = ? WHERE id = ?",
        user.Username, user.Email, user.Age, user.ID)
    return err
}

func (dao *UserDAO) Delete(id int64) error {
    _, err := dao.db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}

func (dao *UserDAO) List(limit, offset int) ([]*User, error) {
    rows, err := dao.db.Query(
        "SELECT id, username, email, age, created_at FROM users LIMIT ? OFFSET ?",
        limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []*User
    for rows.Next() {
        user := &User{}
        err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.CreatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, rows.Err()
}
```

## 测试

### 单元测试
```go
func TestUserDAO(t *testing.T) {
    // 使用测试数据库或内存数据库
    db := setupTestDB(t)
    defer db.Close()
    
    dao := NewUserDAO(db)
    
    // 测试创建
    user := &User{
        Username: "testuser",
        Email:    "test@example.com",
        Age:      25,
    }
    err := dao.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
    
    // 测试查询
    fetched, err := dao.GetByID(user.ID)
    assert.NoError(t, err)
    assert.Equal(t, user.Username, fetched.Username)
}
```

## 常见问题

### 1. 连接超时
检查网络、防火墙、MySQL 配置。

### 2. 字符集乱码
确保连接字符串和表字符集一致（建议 utf8mb4）。

### 3. 时区问题
使用 `parseTime=True&loc=Local` 正确处理时区。

### 4. 连接泄漏
确保及时关闭 `rows`、`stmt`、`tx`。

### 5. 内存泄漏
及时释放大字段数据（如 TEXT、BLOB）。

## 扩展学习

### 1. ORM 框架
- GORM: 功能丰富的 ORM
- XORM: 轻量级 ORM
- sqlx: database/sql 的扩展

### 2. 数据库迁移工具
- goose: 简单的数据库迁移工具
- migrate: 功能完整的迁移工具

### 3. 连接池监控
使用 `db.Stats()` 监控连接池状态。

### 4. 读写分离
使用中间件或自定义实现读写分离。

## 总结

Go 的 `database/sql` 包提供了简洁而强大的数据库操作接口。通过合理使用连接池、预处理语句、事务等特性，可以构建高性能、安全的数据库应用。

建议在实际开发中：
1. 封装数据访问层，隔离数据库操作
2. 合理配置连接池参数
3. 使用预处理语句防止 SQL 注入
4. 正确处理 NULL 值和错误
5. 编写单元测试确保数据访问逻辑正确

随着对 `database/sql` 的深入理解，可以进一步学习 ORM 框架、数据库迁移、性能优化等高级主题。