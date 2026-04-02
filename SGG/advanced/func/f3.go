package main

import "fmt"

// 匿名函数
// 匿名函数没有函数名，直接通过func关键字定义，匿名函数可以立即调用

// 全局匿名函数，如果匿名函数赋给一个全局变量，那么这个匿名函数就可以在其他函数中调用
var a2 = func(n1 int, n2 int) {
	println("全局匿名函数被调用了", n1*n2)
}

func main() {
	// 定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	n1 := func(n1 int, n2 int) int {
		println("匿名函数被调用了", n1, n2)
		return n1 + n2
	}(1, 2)
	fmt.Println(n1)
	// 定义匿名函数并赋值给一个变量，这种方式匿名函数可以调用多次
	a1 := func(n1 int, n2 int) int {
		println("匿名函数被调用了", n1, n2)
		return n1 + n2
	}
	a1(3, 4)
	a1(5, 6)

	a2(7, 8)

	// 定义可以自调用的匿名函数
	// 1. 先声明函数变量及其签名
	var a3 func(int) int

	// 2. 将匿名函数赋值给已声明的变量
	a3 = func(n1 int) int {
		if n1 <= 1 {
			return n1
		}
		// 现在 a3 在当前作用域已定义，可以递归调用
		return a3(n1-1) + n1
	}

	// 测试结果
	result := a3(5) // 1+2+3+4+5 = 15
	fmt.Println(result)
}
