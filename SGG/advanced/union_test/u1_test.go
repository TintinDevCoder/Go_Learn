package main

import (
	"testing"
)

// testing包提供对Go包的自动化测试，测试用例通过测试包中的测试函数进行编写。
// 测试函数的格式为func TestXxx(t *testing.T)，其中t为testing.T类型的对象。其中Xxx中第一个字母不能是小写字母，否则Go会认为该函数不是测试函数。
// 通过运行go test命令，Go会自动寻找并运行所有以Test开头的测试函数。
func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res != 3 {
		// Fatalf函数会停止测试，并打印错误信息
		// 该方法的第一个参数必须是常量字符串，不能是变量，而后面的参数可以是变量，填充前面的参数
		t.Fatalf("测试失败，结果为：%d", res)
	}
	// 正确，输出日志
	t.Logf("测试成功，结果为：%d", res)
}

// 运行程序，发现testing包被引入，就会调用testing框架
// testing框架将 xxx_test的文件引入，将以Test开头的函数作为测试函数，调用函数并打印结果
// test文件的Test方法引入待测试的方

// 注意事项：
// 1、文件命名规范：测试用例文件名必须以 _test.go 结尾。例如 cal_test.go，其中 cal 部分不是固定的。
// 2、函数命名规范：测试用例函数名必须以 Test 开头。一般来说就是 Test + 被测试的函数名，例如 TestAddUpper。
// 3、函数签名：TestAddUpper(t *testing.T) 的形参类型必须是 *testing.T。
// 4、多用例支持：一个测试用例文件中可以包含多个测试用例函数，例如 TestAddUpper、TestSub。
// 5、运行指令 (Command-line Arguments)：
//  (1) cmd > go test：如果运行正确，无日志；只有错误时才会输出日志。
//  (2) cmd > go test -v：无论运行正确还是错误，都会输出详细日志 (Verbose)。
// 6、错误处理：当出现错误时，可以使用 t.Fatalf 来格式化输出错误信息并退出程序。
// 7、日志输出：使用 t.Logf 方法可以输出相应的调试日志。
// 8、执行机制：测试用例函数并没有放在 main 函数中也执行了，这就是测试用例的方便之处。
// 9、结果标识：PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败。
// 10、测试单个文件：必须带上被测试的原文件。
//  指令：go test -v cal_test.go cal.go
// 11、测试单个方法 (Method)：
//  指令：go test -v -test.run TestAddUpper
