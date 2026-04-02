package main

import (
	"fmt"
	"strconv"
)

/*
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

Go 语言类型转换基本格式如下：
type_name(expression)
type_name 为类型，expression 为表达式
*/
func intToFloat() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

// 字符串类型转换
func stringAndInteger() {
	// 字符串转化为浮点数
	str := "123"
	// 参数一：要转换的字符串，
	// 参数二：进制基数（取值范围是 0, 2 到 36）。如果 base 为 0，Go 将自动推断数字格式。
	// 参数三：指定要解析的整数的位数，可以是 0 到 64。此参数用于确定返回值的大小。例如，使用 64 会解析为 int64 类型，32 会解析为 int32。
	num, err := strconv.ParseInt(str, 10, 0)
	// 或者使用 Atoi 函数
	// num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}
	// 浮点数转化为字符串
	// 参数一：要转换为字符串的整数。此参数是 int64 类型。
	// 参数二：指定进制基数（取值范围为 2 到 36）
	num2 := 123
	str2 := strconv.FormatInt(int64(num2), 10)
	// 或者使用 Itoa 函数
	// str2 := strconv.Itoa(num2)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str2)
}
func stringAndFloat() {
	// 字符串转化为浮点数
	str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
	}
	/*
		f	float64	num2	待转换的浮点数值。
		fmt	byte	'f'	格式标记。决定了输出的展现形式（如是否带指数）。
		prec	int	2	精度。对于 'f' 格式，指的是小数点后的位数。
		bitSize	int	64	位大小。指定转换时参考的原始类型是 float32 (32) 还是 float64 (64)。
	*/
	// 浮点数转化为字符串
	num2 := 3.14
	str2 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num2, str2)
}

// 接口类型转换
/**
类型断言
类型断言用于将接口类型转换为指定类型，其语法为：
value.(type) 或者 value.(T)
其中 value 是接口类型的变量，type 或 T 是要转换成的类型。
如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功
*/
func interfaceChange1() {
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

/*
*
类型转换
类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
T(value)
T 是目标接口类型，value 是要转换的值。
在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。
*/
// 定义一个接口 Writer
type Writer interface {
	Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}
func interfaceChange2() {
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, World"

	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)
}
func main() {
	interfaceChange2()
}
