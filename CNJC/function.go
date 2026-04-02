package main

import "fmt"

// 函数
/**
func function_name( [parameter list] ) [return_types] {
   函数体
}
*/
/* 函数返回两个数的最大值 */
// 逃逸分析（Escape Analysis）是在编译期完成的
func max(num1 float32, num2 int) {
	/* 声明局部变量 */
	var result interface{}
	fmt.Printf("result 接口变量本身的栈地址: %p\n", &result)
	if num1 > float32(num2) {
		result = num1
	} else {
		result = num2
	}
	fmt.Println(result)
	fmt.Printf("num1: %p\n", &num1)
	fmt.Printf("result 接口变量本身的栈地址: %p\n", &result)
}

func swap(x, y string) (string, string) {
	return y, x
}

// 变量
/* 声明全局变量 */
var g int

// 全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
func bl_test1() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b
	g = a + b + c
	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)
	fmt.Printf("全局变量 g = %d\n", g)
}
func main() {
	//max(20.5, 10)
	bl_test1()
}
