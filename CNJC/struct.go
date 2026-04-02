package main

import "fmt"

/*
*
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

	type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/
type Books struct {
	Title   string
	Author  string
	Subject string
	Book_id int
}

func struct_test1() {
	book1 := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	// 创建一个新的结构体
	fmt.Println(book1)

	// 也可以使用 key => value 格式
	fmt.Println(Books{Title: "Go 语言", Author: "www.runoob.com", Subject: "Go 语言教程", Book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{Title: "Go 语言", Author: "www.runoob.com"})
}

// 结构体作为函数参数
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.Book_id)
}
func struct_test2() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.Title = "Go 语言"
	Book1.Author = "www.runoob.com"
	Book1.Subject = "Go 语言教程"
	Book1.Book_id = 6495407

	/* book 2 描述 */
	Book2.Title = "Python 教程"
	Book2.Author = "www.runoob.com"
	Book2.Subject = "Python 语言教程"
	Book2.Book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	// 结构体指针
	/*
		可以定义指向结构体的指针类似于其他指针变量，格式如下：

		var struct_pointer *Books
		以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

		struct_pointer = &Book1
		使用结构体指针访问结构体成员，使用 "." 操作符：

		struct_pointer.title
	*/
	var pbook1 *Books
	pbook1 = &Book1
	author := pbook1.Author
	fmt.Println(author)
}

func main() {
	struct_test2()
}
