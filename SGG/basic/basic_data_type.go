package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 整数数据类型
	var a int = 10
	var b int8 = 20
	// var b2 int8 = 128 //超出int8的范围，编译错误，int8的范围是-128~127
	var c int16 = 30
	var d int32 = 40
	var e int64 = 50

	// 无符号整数数据类型
	var f uint = 60
	var g uint8 = 70
	var h uint16 = 80
	var i uint32 = 90
	var j uint64 = 100

	// 浮点数数据类型
	var k float32 = 3.14
	var l float64 = 3.141592653589793

	// 布尔数据类型
	var m bool = true

	// 字符数据类型
	// 使用byte存储单个字母字符
	var n byte = 'A' // byte是uint8的别名，表示一个字节
	// 使用rune存储单个Unicode字符
	var o rune = '中' // rune是int32的别名，表示一个Unicode码点

	// 字符串数据类型
	var p string = "Hello, Go!"

	// 输出变量的值
	println(a, b, c, d, e)
	println(f, g, h, i, j)
	println(k, l)
	println(m)
	println(n, o)
	println(p)

	// 查看某个变量所占用的字节大小和数据类型
	// unsafe.Sizeof可以返回变量的字节数
	// %T可以输出变量的类型，%v可以输出变量的值，%#v可以输出变量的详细信息，包括类型和值
	// 存放的大小是根据数据类型来决定的，int类型在32位系统上占4字节，在64位系统上占8字节，float32占4字节，float64占8字节，bool占1字节，
	// byte占1字节，rune占4字节，string占16字节（包含指向字符串数据的指针和字符串长度）
	// int默认使用平台相关的大小，通常是32位或64位，具体取决于编译器和操作系统
	fmt.Printf("Size of a:%d bytes, Type of a:%T, 值为:%v \n", unsafe.Sizeof(a), a, a)
	fmt.Printf("Size of b:%d bytes, Type of b:%T, 值为:%v \n", unsafe.Sizeof(b), b, b)

	// 浮点数可能造成精度损失
	// float默认使用float64类型，float32的精度较低，可能会导致数值不准确
	var fln1 float32 = -123.0000901
	var fln2 float64 = -123.0000901
	fmt.Printf("Size of fln1:%d bytes, Type of fln1:%T, 值为:%v \n", unsafe.Sizeof(fln1), fln1, fln1)
	fmt.Printf("Size of fln2:%d bytes, Type of fln2:%T, 值为:%v \n", unsafe.Sizeof(fln2), fln2, fln2)
	var sum float64
	for i := 0; i < 10; i++ {
		sum += 0.1
	}
	fmt.Println(sum) // 输出可能是 0.9999999999999999 而不是 1.0
	// 科学计数法
	fln3 := 1.23e4 // 等价于 1.23 * 10^4
	fmt.Printf("Size of fln3:%d bytes, Type of fln3:%T, 值为:%v \n", unsafe.Sizeof(fln3), fln3, fln3)

	// 直接输出字符类型的变量时，会输出对应的字符码值
	fmt.Printf("Size of n:%d bytes, Type of n:%T, 值为:%v \n", unsafe.Sizeof(n), n, n)
	// 如果想要输出字符本身，可以使用%c格式化符
	fmt.Printf("Size of n:%d bytes, Type of n:%T, 值为:%c \n", unsafe.Sizeof(n), n, n)

	var c1 int = 'A'
	fmt.Printf("c1=%c 对应的整数值是:%d \n", c1, c1)
	// 字符类型是可以进行运算的，相当于一个整数
	var c2 int = 'B'
	fmt.Printf("c2=%c 对应的整数值是:%d \n", c2, c2)
	fmt.Printf("c1 + c2 = %d \n", c1+c2) // 输出 131，因为 'A' 的整数值是 65，'B' 的整数值是 66
	/*
		存储：字符-》对应码值（整数）-》二进制-》存储在内存中
		读取：内存中的二进制-》对应的整数-》对应的字符
		go语言使用的全部都是utf8
	*/

	// 字符串类型的变量占用16字节，因为字符串在Go语言中是一个结构体，包含一个指向字符串数据的指针和字符串的长度
	fmt.Printf("Size of p:%d bytes, Type of p:%T, 值为:%v \n", unsafe.Sizeof(p), p, p)
	// 字符串是不可变的，修改字符串会创建一个新的字符串
	// 字符串的两种表现形式
	// 双引号，可以包含转义字符，支持多行字符串
	str1 := "Hello\nWorld"
	fmt.Printf("str1:%s \n", str1)
	// 反引号，原样输出字符串，不支持转义字符，适合表示多行字符串
	str2 := `Hello\nWorld`
	fmt.Printf("str2:%s \n", str2)

	// 各种数据类型的默认值
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string
	fmt.Printf("Default value of int: %d \n", defaultInt)         // 输出 0
	fmt.Printf("Default value of float64: %f \n", defaultFloat)   // 输出 0.000000
	fmt.Printf("Default value of bool: %t \n", defaultBool)       // 输出 false
	fmt.Printf("Default value of string: '%s' \n", defaultString) // 输出 ''

	// 基本数据类型的转换
	// Go语言中的类型转换需要显式进行，不能隐式转换
	// Type(value)的形式可以将value转换为T类型，如果转换不合法会导致编译错误
	// Go中数据类型的转换可以是表示范围小-》大，也可以是表示范围大-》小，但需要注意可能会发生数据丢失或溢出

	var x int = 10
	var y float64 = float64(x) //x本身的数据类型没有变化，只是将x的值转换为float64类型来赋值给y
	fmt.Printf("x: %d, y: %f \n", x, y)
	var z int = int(y)
	fmt.Printf("z: %d \n", z)

	// 可能会发生数据丢失或溢出
	var nu1 int64 = 999999
	var nu2 int8 = int8(nu1)
	// 输出 nu1: 999999, nu2: 63，因为 int8 的范围是 -128 到 127，999999 超出了这个范围，导致溢出，最终结果是 999999 % 256 = 63
	fmt.Printf("nu1: %d, nu2: %d \n", nu1, nu2)

	// 字符类型和整数类型之间的转换
	var char byte = 'A'
	var charToInt int = int(char)
	fmt.Printf("char: %c, charToInt: %d \n", char, charToInt)
	var intToChar byte = byte(charToInt)
	fmt.Printf("intToChar: %c \n", intToChar)

}
