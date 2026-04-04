package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 定义用户结构体
type User struct {
	ID        int64
	Username  string
	Email     string
	Age       int
	CreatedAt time.Time
}

func main() {
	// 1. 连接数据库
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("数据库连接成功")

	// 2. 创建表
	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 插入数据
	userID, err := insertUser(db, "john_doe", "john@example.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("插入用户成功，ID: %d\n", userID)

	// 4. 查询单条数据
	user, err := getUserByID(db, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查询用户: %+v\n", user)

	// 5. 批量插入
	err = batchInsertUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	// 6. 查询多条数据
	users, err := getUsersByAge(db, 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("年龄大于25的用户数量: %d\n", len(users))

	// 7. 更新数据
	rowsAffected, err := updateUserAge(db, userID, 31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("更新影响行数: %d\n", rowsAffected)

	// 8. 删除数据
	rowsAffected, err = deleteUser(db, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("删除影响行数: %d\n", rowsAffected)

	// 9. 事务示例
	err = transactionExample(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("所有操作完成")
}

// 连接数据库
func connectDB() (*sql.DB, error) {
	// DSN格式: username:password@protocol(address)/dbname?param=value
	dsn := "root:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 验证连接
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// 创建表
func createTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL,
		age INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_username (username),
		INDEX idx_age (age)
	)`
	_, err := db.Exec(createTableSQL)
	return err
}

// 插入用户
func insertUser(db *sql.DB, username, email string, age int) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO users (username, email, age) VALUES (?, ?, ?)",
		username, email, age)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// 根据ID查询用户
func getUserByID(db *sql.DB, id int64) (*User, error) {
	user := &User{}
	err := db.QueryRow(
		"SELECT id, username, email, age, created_at FROM users WHERE id = ?",
		id).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 批量插入用户
func batchInsertUsers(db *sql.DB) error {
	// 准备预处理语句
	stmt, err := db.Prepare("INSERT INTO users (username, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 批量插入
	for i := 1; i <= 5; i++ {
		username := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		age := 20 + i

		_, err := stmt.Exec(username, email, age)
		if err != nil {
			return err
		}
	}

	return nil
}

// 根据年龄查询用户
func getUsersByAge(db *sql.DB, minAge int) ([]*User, error) {
	rows, err := db.Query(
		"SELECT id, username, email, age, created_at FROM users WHERE age > ? ORDER BY created_at DESC",
		minAge)
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

	// 检查迭代过程中的错误
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// 更新用户年龄
func updateUserAge(db *sql.DB, id int64, newAge int) (int64, error) {
	result, err := db.Exec(
		"UPDATE users SET age = ? WHERE id = ?",
		newAge, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// 删除用户
func deleteUser(db *sql.DB, id int64) (int64, error) {
	result, err := db.Exec(
		"DELETE FROM users WHERE id = ?",
		id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// 事务示例：转账操作
func transactionExample(db *sql.DB) error {
	// 创建测试表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
		)`)
	if err != nil {
		return err
	}

	// 插入测试数据
	_, err = db.Exec("INSERT IGNORE INTO accounts (name, balance) VALUES ('Alice', 1000), ('Bob', 500)")
	if err != nil {
		return err
	}

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 如果事务中出现错误，回滚
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 从Alice账户扣除100
	_, err = tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE name = 'Alice'")
	if err != nil {
		return err
	}

	// 检查Alice余额是否足够（实际应用中应该先检查）
	var aliceBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE name = 'Alice'").Scan(&aliceBalance)
	if err != nil {
		return err
	}
	if aliceBalance < 0 {
		return fmt.Errorf("余额不足")
	}

	// 向Bob账户增加100
	_, err = tx.Exec("UPDATE accounts SET balance = balance + 100 WHERE name = 'Bob'")
	if err != nil {
		return err
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("转账事务执行成功")
	return nil
}

// 处理NULL值的示例
func nullExample(db *sql.DB) error {
	// 创建包含NULL值的表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS nullable_demo (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50),
			age INT,
			email VARCHAR(100)
		)`)
	if err != nil {
		return err
	}

	// 插入包含NULL值的数据
	_, err = db.Exec("INSERT INTO nullable_demo (name, age, email) VALUES (?, NULL, ?)", "John", "john@example.com")
	if err != nil {
		return err
	}

	// 查询NULL值
	var name sql.NullString
	var age sql.NullInt64
	var email string

	err = db.QueryRow("SELECT name, age, email FROM nullable_demo WHERE id = 1").Scan(&name, &age, &email)
	if err != nil {
		return err
	}

	if name.Valid {
		fmt.Printf("Name: %s\n", name.String)
	} else {
		fmt.Println("Name is NULL")
	}

	if age.Valid {
		fmt.Printf("Age: %d\n", age.Int64)
	} else {
		fmt.Println("Age is NULL")
	}

	fmt.Printf("Email: %s\n", email)

	return nil
}
