package main

import "fmt"

// add 是一个简单的加法函数，用于演示单元测试
// 参数:
//   - a: 第一个加数
//   - b: 第二个加数
//
// 返回值:
//   - int: a 和 b 的和
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// traditionalTest 是传统测试方法的示例函数
// 传统方法的缺点：
// 1. 不方便：必须停止正在运行的程序才能执行测试
// 2. 不易管理：测试多个函数或模块时，都需要写在 main 函数中，导致代码混乱
// 3. 缺乏标准化：没有统一的测试输出和错误处理方式
// 4. 难以自动化：不适合集成到持续集成/持续部署 (CI/CD) 流程中
func traditionalTest() {
	// 使用 panic 来报告测试失败，但这会导致整个程序崩溃
	if add(1, 2) != 3 {
		panic("add error")
	}
	fmt.Println("add test passed")
}

// golang 中自带一个轻量级的测试框架 testing 和自带的 go test 命令实现单元测试和性能测试
// 此命令会运行本包下所有测试函数，加上-v会输出详细测试信息
// testing 框架的优势：
// 1、确保每个函数时可运行的，并且运行结果正确
// 2、确保写出来代码性能是好的
// 3、及时发现代码的逻辑错误
func main() {
	// 调用传统测试方法
	traditionalTest()

}
