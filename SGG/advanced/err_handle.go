package main

import (
	"errors"
	"fmt"
)

// golang 的错误处理
// 默认情况下，当错误发生时，程序会退出并输出错误信息，但有时候我们希望在发生错误时能够继续执行程序，这时就需要使用错误处理机制。
// Go中可以抛出一个panic的异常，然后再defer中通过recocer捕获这个异常，而后正常处理
// 使用defer + recover 来捕获处理异常
func ceshi() {
	defer func() {
		err := recover() // recover()函数可以捕获到panic的异常，并返回异常信息，如果没有异常发生，则返回nil
		if err != nil {
			println("捕获到异常：", err)
			// 这里可以继续处理异常，例如记录日志、发送通知等
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	println(result)
}

// 自定义错误
// errors.New()函数可以创建一个新的错误对象，参数是错误信息字符串，返回一个实现了error接口的错误对象
// panic()函数可以触发一个运行时错误，导致程序崩溃，参数是错误信息字符串，可以是任意类型的值
func readConf(name string) (err error) {
	if name == "config.yaml" {
		return nil
	} else {
		return errors.New("配置文件错误")
	}
}

// Go语言中没有异常机制，但提供了一个内置的错误接口error，error接口是一个内置的接口类型，定义如下：
//
//	type error interface {
//	    Error() string
//	}
//
// 任何实现了Error() string方法的类型都可以被视为一个错误类型，Go语言中常见的错误类型有：
// 1. 内置错误类型：Go语言内置了一些常见的错误类型，如os.ErrNotExist、os.ErrPermission等，这些错误类型都实现了error接口，可以直接使用。
// 2. 自定义错误类型：我们也可以定义自己的错误类型，只需要实现Error() string方法即可，例如：
type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}
func testMyError(a, b int) (int, error) {
	if b == 0 {
		return -1, &MyError{Message: "除数不能为零"}
	}
	return a / b, nil
}

// errors.As(err, &target)函数可以判断一个错误是否是某个特定类型的错误，如果是，则返回true，否则返回false，例如：
// 链路搜索：在 err 的包装链（Wrapping Chain）中自上而下寻找。
// 类型匹配：寻找第一个与 target 所指向的类型一致的错误。
// 值拷贝（转换）：如果找到匹配项，它会将该错误实例的值赋值给 target 变量。
func main() {
	ceshi()
	fmt.Println("程序继续执行")

	err := readConf("config.yaml")
	if err != nil {
		// 输出错误，并终止程序
		panic(err)
	}

	num, err := testMyError(1, 0)
	if err != nil {
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Printf("发生错误：%s\n", myErr.Message)
		} else {
			fmt.Println("发生错误：", err.Error())
		}
	} else {
		fmt.Println(num)
	}
}
