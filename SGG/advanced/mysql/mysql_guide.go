package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // 只执行驱动的初始化代码，不直接使用驱动包
)

var db *sql.DB

// 数据库的基本连接
func connectToDb() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local"
	// sql.Open不会立即建立连接，而是返回一个*sql.DB对象，只有在第一次使用时才会尝试连接数据库。
	// 因此，sql.Open的错误通常是由于参数错误引起的，而不是连接失败。
	// 第一个参数是数据库驱动的名称，这里使用的是MySQL，所以是"mysql"，第二个参数是数据源名称（DSN），包含了连接数据库所需的信息，如用户名、密码、地址、数据库名称等。
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 验证连接
	err = db.Ping() // 此时第一次使用db对象，才会尝试连接数据库，如果连接失败，Ping会返回错误。
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("成功连接到数据库")
	return
}

// 连接池的配置
// sql.DB对象内部维护了一个连接池，默认情况下，连接池的最大打开连接数为0（表示无限制），最大空闲连接数为2，连接的最大生命周期为0（表示无限制）。可以通过以下方法配置连接池：
// db.SetMaxOpenConns(25) // 设置最大打开连接数为25
// db.SetMaxIdleConns(10) // 设置最大空闲连接数为10
// db.SetConnMaxLifetime(5 * time.Minute) // 设置连接的最大生命周期为5分钟
func dbPool() {
	db.SetMaxOpenConns(25)                 // 最大打开连接数
	db.SetMaxIdleConns(10)                 // 最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute) // 连接最大生存时间
}

type user struct {
	ID       int
	Username string
	Email    string
	Age      int
}

// 创建表
func goCreateTable() (err error) {
	createTableSQL := `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
	// db.Exec方法返回两个参数，第一个参数是一个Result,它包含了执行结果的信息，如受影响的行数和最后插入的ID等；
	// 第二个参数是一个error类型的错误对象，如果执行过程中发生了错误，error对象将包含错误信息，否则为nil。
	e, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
		return
	}
	rowsAffected, _ := e.RowsAffected()
	fmt.Printf("创建表成功，受影响的行数: %d\n", rowsAffected)
	lastInsertID, _ := e.LastInsertId()
	fmt.Printf("最后插入的ID: %d\n", lastInsertID)
	return
}

// 插入数据
func goInsertData(u user) (id int, err error) {
	fmt.Println("插入数据")

	// 简单插入
	e, err := db.Exec("INSERT INTO users (username, email, age) VALUES (?, ?, ?)",
		u.Username, u.Email, u.Age)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("插入数据成功\n")

	lastInsertID, _ := e.LastInsertId()
	fmt.Printf("插入的ID: %d\n", lastInsertID)
	id = int(lastInsertID)

	rowsAffected, _ := e.RowsAffected()
	fmt.Printf("受影响的行数: %d\n", rowsAffected)
	return
}

// 查询数据
// db.QueryRow查询数据库，使用Scan方法将查询结果扫描到变量中。
// 此时是按照sql的顺序来扫描的，所以Scan方法中的参数顺序必须与SQL查询语句中的列顺序一致。
func goQueryData(id int) (u user, err error) {
	fmt.Println("查询数据")
	// 查询单行
	err = db.QueryRow("SELECT id, username, email, age FROM users WHERE id = ?", id).Scan(
		&u.ID, &u.Username, &u.Email, &u.Age)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

// db.Query方法返回一个*sql.Rows对象，可以使用Next方法迭代结果集，并使用Scan方法将每行数据扫描到变量中。
// db.Query方法执行sql语句，查询出一个结果集，返回一个*sql.Rows对象。可以使用Next方法迭代结果集，并使用Scan方法将每行数据扫描到变量中。
// Next方法使得go程序从mysql的缓存中逐行读取数据，直到没有更多的行可供读取为止。
// 每次调用Next方法时，都会将当前行的数据加载到内存中，并将指针移动到下一行。
// 当Next方法返回false时，表示已经没有更多的行可供读取了，此时需要检查是否发生了错误。
// 在迭代完成后，还需要调用Rows.Err方法来检查迭代过程中是否发生了错误。
func goQueryMoreData(age int) (users []user, err error) {
	// 查询多行
	rows, err := db.Query("SELECT id, username, email, age FROM users WHERE age >= ?", age)
	defer rows.Close()
	var u user
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Age)
		if err != nil {
			log.Fatal(err)
			return
		}
		users = append(users, u)
	}
	err = rows.Err() // 检查迭代过程中的错误
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

// 更新数据
func goUpdateData(id int, age int) (err error) {
	result, err := db.Exec("UPDATE users SET age = ? WHERE id = ?", age, id)
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("更新成功，受影响的行数: %d\n", rows)
	return
}

// 删除数据
func goDeleteData(id int) (err error) {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := result.RowsAffected()
	fmt.Printf("删除成功，受影响的行数: %d\n", rows)
	return
}

// 预处理语句
// 预处理语句是一种在数据库中预先编译的SQL语句，可以提高执行效率和安全性。
// 通过使用预处理语句，可以避免SQL注入攻击，并且在需要多次执行相似的SQL语句时，可以减少编译的开销。
// db.Prepare返回一个*sql.Stmt对象，表示预处理语句，可以使用Exec方法执行预处理语句，并传入参数。预处理语句在执行完成后需要调用Close方法关闭，以释放资源。
// Stmt.Exec将动态生成的参数“注入”到预先定义的 SQL 模板中并执行。
func prepareSql() {
	// 准备预处理语句
	stmt, err := db.Prepare("INSERT INTO users (username, email, age) VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 重复使用
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(fmt.Sprintf("user%d", i),
			fmt.Sprintf("user%d@example.com", i), 20+i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

// 事务
// 事务是一组操作的集合，这些操作要么全部成功，要么全部失败。事务可以确保数据的一致性和完整性，特别是在涉及多个相关操作时。
// db.Begin方法开始一个新的事务，返回一个*sql.Tx对象，表示事务。可以使用Tx.Exec方法执行事务中的SQL语句，并传入参数。
// 如果在执行过程中发生错误，可以调用Tx.Rollback方法回滚事务，撤销已经执行的操作。
// 如果所有操作都成功，可以调用Tx.Commit方法提交事务，使所有操作生效。
func transactionSql() {
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
}

// 批量操作
func batchSql() {
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
	if err != nil {
		log.Fatal(err)
		return
	}
}
func main() {
	err := connectToDb()
	if err != nil {
		return
	}
	dbPool()

	fmt.Println("创建user表")
	goCreateTable()

	id, err := goInsertData(user{Username: "dd", Email: "182", Age: 18})
	if err != nil {
		return
	}
	fmt.Printf("插入的用户ID: %d\n", id)

	u, err := goQueryData(id)
	if err != nil {
		return
	}
	fmt.Printf("查询到的用户: %+v\n", u)

	// 插入多条测试数据
	goInsertData(user{Username: "ss", Email: "1821", Age: 20})
	goInsertData(user{Username: "yy", Email: "18211", Age: 22})
	users, err := goQueryMoreData(18)
	if err != nil {
		return
	}
	fmt.Printf("查询到的用户: %+v\n", users)

	// 更新用户年龄
	err = goUpdateData(id, 20)
	if err != nil {
		return
	}

	// 删除用户
	err = goDeleteData(id)
	if err != nil {
		return
	}

	defer db.Close()
}
