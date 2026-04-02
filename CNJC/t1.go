package main

import "fmt"

/*
*
格式化输出
*/
func init() {
	//fmt.Println("init")
}
func test1() {
	fmt.Println("Hello, World!")
}
func test2() {
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	target_url := fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

}
func test3() {
	const num1, num2, num3 = 5, 10, 15
	num4 := fmt.Sprintf("%d %d %d", num1, num2, num3)
	fmt.Println("num4 =", num4)
}

type Person struct {
	Name string
	Age  int
}

func test4() {
	person := Person{Name: "Alice", Age: 30}

	// 使用 %+v 打印结构体
	fmt.Printf("Person details: %+v\n", person)
	fmt.Printf("%T\n", person)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%U\n", 65)

}
func main() {
	test4()
	fmt.Println(book{title: "Go 语言", author: "www.runoob.com"})
}
