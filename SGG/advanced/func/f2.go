package main

import "fmt"

var age int = 1

// 全局变量定义 -》 init函数 -》 main函数
// init函数会在main函数之前执行，且init函数不能被调用
// 被引入的包的init函数会在当前包的init函数之前执行
func init() {
	fmt.Println("init函数被调用了")
}

// defer关键词用于在函数返回之前执行一些操作
// defer语句会将函数调用推迟到当前函数返回之后执行
// defer语句可以用来释放资源、关闭文件、解锁等操作
func defertest() {
	defer fmt.Println("defer语句被调用了")
	fmt.Println("defertest函数被调用了")
}
func main() {
	fmt.Println("main函数被调用了")
	fmt.Println(age)

	defertest()
}
