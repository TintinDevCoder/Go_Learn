package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型和字符串的转换
func main() {
	// 基本数据类型转换为字符串
	// 方式一
	// 使用Springf函数将基本数据类型转换为字符串，%v表示以默认格式输出变量的值，%T表示输出变量的类型，
	// %d表示输出整数，%.1f表示输出浮点数，%t表示输出布尔值，%c表示输出字符
	var a int = 10
	var b float64 = 3.14
	var c bool = true
	var d byte = 'A'
	var e rune = '中'
	var str string
	str = fmt.Sprintf("%d", a) // 将整数转换为字符串
	fmt.Printf("str:%s\n", str)
	str = fmt.Sprintf("%.1f", b)
	fmt.Printf("str:%s\n", str)
	str = fmt.Sprintf("%t", c)
	fmt.Printf("str:%s\n", str)
	str = fmt.Sprintf("%c", d)
	fmt.Printf("str:%s\n", str)
	str = fmt.Sprintf("%c", e)
	fmt.Printf("str:%s\n", str)

	// 方式二
	// 使用strconv包中的函数将基本数据类型转换为字符串，
	// strconv.FormatInt可以将整数转换为字符串
	var num1 = 123
	str = strconv.FormatInt(int64(num1), 10) // 将整数转换为字符串，10表示十进制
	fmt.Printf("str:%s\n", str)
	// strconv.Itoa可以将整数转换为字符串，
	var num2 int = 123
	str = strconv.Itoa(num2)
	fmt.Printf("str:%s\n", str)

	// strconv.FormatFloat可以将浮点数转换为字符串
	var num3 float64 = 3.141592653589793
	// 将浮点数转换为字符串，'f'表示以小数点形式输出，10表示保留的小数位，64表示float64类型
	str = strconv.FormatFloat(num3, 'f', 10, 64)
	fmt.Printf("str:%s\n", str)

	// strconv.FormatBool可以将布尔值转换为字符串
	var boolVal bool = true
	str = strconv.FormatBool(boolVal)
	fmt.Printf("str:%s\n", str)

	// strconv.FormatUint可以将无符号整数转换为字符串
	var num4 uint = 123
	str = strconv.FormatUint(uint64(num4), 10)
	fmt.Printf("str:%s\n", str)

	// 字符串转换为基本数据类型
	// 返回值和错误，如果转换成功，错误为nil，如果转换失败，返回值为0，错误不为nil
	// strconv.ParseInt可以将字符串转换为整数
	str = "123"
	// str是一个字符串，10表示十进制，64表示int64类型
	num8, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("转换错误:%v\n", err)
	} else {
		fmt.Printf("num8:%d\n", num8)
	}

	// strconv.Atoi可以将字符串转换为整数
	str = "123"
	num5, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("转换错误:%v\n", err)
	} else {
		fmt.Printf("num5:%d\n", num5)
	}

	// strconv.ParseFloat可以将字符串转换为浮点数
	str = "3.141592653589793"
	// str是一个字符串，64表示float64类型
	num6, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("转换错误:%v\n", err)
	} else {
		fmt.Printf("num6:%f\n", num6)
	}

	// strconv.ParseBool可以将字符串转换为布尔值
	str = "true"
	boolVal2, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Printf("转换错误:%v\n", err)
	} else {
		fmt.Printf("boolVal2:%t\n", boolVal2)
	}

	str = "hello"
	num9, err := strconv.Atoi(str)
	// 转换错误会赋予默认值
	if err != nil {
		fmt.Printf("num9:%d\n", num9) // 输出num9:0
		fmt.Printf("转换错误:%v\n", err)
	} else {
		fmt.Printf("num9:%d\n", num9)
	}
}
