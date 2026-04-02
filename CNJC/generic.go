package main

import (
	"cmp"
	"fmt"
)

/*
泛型是 Go 语言在 1.18 版本中引入的重要特性，它让开发者能够编写更加灵活和可重用的代码。

泛型主要通过以下两个核心概念来实现：

类型参数（Type Parameters）：允许你在函数或类型定义中使用一个或多个类型作为参数。
类型约束（Type Constraints）：指定类型参数必须满足的条件，确保在函数内部可以安全地操作这些类型。

泛型（Generics）允许我们编写不依赖特定数据类型的代码。
在引入泛型之前，如果我们想要处理不同类型的数据，通常需要为每种类型编写重复的函数。
*/
// 传统方法：
// 处理 int 类型的函数
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 处理 float64 类型的函数
func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 使用泛型的解决方案：
// 一个函数处理多种类型
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

/*
// 基本语法结构
func 函数名[T 约束](参数 T) 返回值类型 {
    // 函数体
}

type 类型名[T 约束] struct {
    // 结构体字段
}
*/
/*
约束（Constraints）
约束定义了类型参数必须满足的条件，是泛型的核心概念。
*/
// 1. any 约束 any 是空接口 interface{} 的别名，表示任何类型都可以。
func PrintAny[T any](value T) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}
func generic_test1() {
	// 使用示例
	PrintAny(42)      // Value: 42, Type: int
	PrintAny("hello") // Value: hello, Type: string
	PrintAny(3.14)    // Value: 3.14, Type: float64
}

// 2. comparable 约束 comparable 表示类型支持 == 和 != 操作符。
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
func generic_test2() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(FindIndex(numbers, 3)) // 输出: 2

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(FindIndex(names, "Bob")) // 输出: 1
}

// 3. 联合约束（Union Constraints） 使用 | 运算符组合多个类型。
// 数字类型约束
// 使用波浪号支持自定义命名类型
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func generic_test3() {
	fmt.Println(Add(10, 20))     // 输出: 30
	fmt.Println(Add(3.14, 2.71)) // 输出: 5.85
}

// 自定义约束
// 1. 方法约束  定义需要特定方法的约束。
// 定义 Stringer 约束
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// 实现自定义类型
type Person2 struct {
	Name string
	Age  int
}

func (p Person2) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}
func generic_test4() {
	person := Person2{Name: "Alice", Age: 25}
	PrintString(person) // 输出: Alice (25 years old)
}

// 2. 复杂约束
// 结合类型和方法要求。
// 要求类型是数字且实现 String() 方法
type NumericStringer interface {
	Number
	String() string
}

// --- 实现该约束的具体类型 ---

// CustomFloat 自定义类型，底层是 float64
type CustomFloat float64

// String 实现方法集合要求
func (cf CustomFloat) String() string {
	return fmt.Sprintf("%.2f", float64(cf))
}

// --- 利用约束进行泛型编写 ---
// FormatAndDouble 这是一个泛型函数
// 它利用 NumericStringer 约束确保 T 既能做数学运算，又能调用 String()
func FormatAndDouble[T NumericStringer](value T) {
	// 这里的 T 被识别为支持 Number 的操作
	// 编译器知道 T 的底层是数字，所以允许 * 2
	result := value * T(2)
	fmt.Printf("Result: %v\n", result)
}
func generic_test5() {
	price := CustomFloat(19.99)
	FormatAndDouble(price)
}
func main() {
	generic_test5()
}
