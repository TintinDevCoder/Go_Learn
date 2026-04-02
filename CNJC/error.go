package main

import (
	"errors"
	"fmt"
)

/*
error 接口
Go 标准库定义了一个 error 接口，表示一个错误的抽象。
*/
/*
实现 error 接口：任何实现了 Error() 方法的类型都可以作为错误。
Error() 方法返回一个描述错误的字符串。
type error interface {
	Error() string
}
*/
// 使用 errors 包创建错误
// 我们可以在编码中通过实现 error 接口类型来生成错误信息。

// 函数通常在最后的返回值中返回错误信息，使用 errors.New 可返回一个错误信息：
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return f, nil
}
func error_test1() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

// 显式返回错误
// Go 中，错误通常作为函数的返回值返回，开发者需要显式检查并处理。
func divide1(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func error_test2() {
	result, err := divide1(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// 自定义错误
// 通过定义自定义类型，可以扩展 error 接口。
type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}
func error_test3() {
	_, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err) // 输出：cannot divide 10 by 0
	}
}

// fmt 包与错误格式化
/*
fmt 包提供了对错误的格式化输出支持：

%v：默认格式。
%+v：如果支持，显示详细的错误信息。
%s：作为字符串输出。
*/
// 定义一个 DivideError 结构
type DivideError2 struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError2) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError2{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func error_test4() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if r2, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg, r2)
	}
}

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}
func error_test5() {
	err := findItem(1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error:", err)
	}
}

// errors.As
// 将错误转换为特定类型以便进一步处理。
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func getError() error {
	return &MyError{Code: 404, Msg: "Not Found"}
}
func error_test6() {
	err := getError()
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Custom error - Code: %d, Msg: %s\n", myErr.Code, myErr.Msg)
	}
}

/*
panic 和 recover
Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。

panic:导致程序崩溃并输出堆栈信息。常用于程序无法继续运行的情况。
recover:捕获 panic，避免程序崩溃。
*/
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}
func error_test7() {
	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")
}
func main() {
	error_test7()
}
