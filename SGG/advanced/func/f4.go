package main

import (
	"fmt"
	"strings"
)

// 闭包是一个函数和与其相关的引用环境组合的一个整体

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var s string = "hello"
	return func(x int) int {
		n += x
		s += fmt.Sprintf("%d", x)
		fmt.Println(s)
		return n
	}
}

// AddUpper是一个函数，返回的数据类型是fun(int) int
// 返回的是一个匿名函数，这个匿名函数被引用到函数外的n，因此这个匿名函数就和n形成了一个整体，构成闭包
// 可以理解闭包为一个类，函数是操作，n是一个字段。
// 函数只有在使用某个变量，才会和这个变量形成一个整体，构成闭包
func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	// 当反复调用f函数时，n的值会不断增加
	// 这就是闭包的特性，函数和引用环境形成一个整体，n的值只初始化一次
	// 因此每次调用就进行一次累加
	suff := makeSuffix()
	result := suff("str.go")
	fmt.Println(result)
}

func makeSuffix() func(str string) string {
	var hz string = "go"
	return func(str string) string {
		index := strings.LastIndex(str, ".")
		if index == -1 {
			return str + "." + hz
		}
		s1 := str[:index+1]
		s2 := str[index+1:]
		if s2 != hz {
			return s1 + hz
		} else {
			return str
		}
	}
}
