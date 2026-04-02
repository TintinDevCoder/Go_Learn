package main

import "fmt"

// 全局变量
var m1 int
var m2 = 2

// 多变量声明
var (
	m3 = 3
	m4 = "hello"
)

func main() {
	// 三种变量声明方式
	var n1 int
	var n2 = 1
	n3 := 2
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)
	// 多变量声明方法
	var s1, s2, s3 int
	var s4, s5, s6 = 0, "tom", 10.1
	s7, s8, s9 := 0, "tom", 10.1
	fmt.Println("s1 = ", s1, "s2 = ", s2, "s3 = ", s3)
	fmt.Println("s4 = ", s4, "s5 = ", s5, "s6 = ", s6)
	fmt.Println("s7 = ", s7, "s8 = ", s8, "s9 = ", s9)

	fmt.Println("m1 = ", m1, "m2 = ", m2, "m3 = ", m3, "m4 = ", m4)

	// 加号两边是数字时，表示加法运算；加号两边是字符串时，表示字符串连接
	num1 := 10
	num2 := 20
	num3 := num1 + num2
	fmt.Println("num1 = ", num1, "num2 = ", num2, "num3 = ", num3)
	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	fmt.Println("str1 = ", str1, "str2 = ", str2, "str3 = ", str3)
}
