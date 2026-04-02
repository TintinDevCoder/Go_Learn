package main

import "fmt"

/*
运算符是一种特殊的符号
Go 语言中的运算符可以分为以下几类：
算术运算符：用于执行基本的数学运算，如加法（+）、减法（-）、乘法（*）、除法（/）和取模（%）。
关系运算符：用于比较两个值，并返回一个布尔结果，如等于（==）、不等于（!=）、大于（>）、小于（<）、大于或等于（>=）和小于或等于（<=）。
逻辑运算符：用于组合多个条件，并返回一个布尔结果，如逻辑与（&&）、逻辑或（||）和逻辑非（!）。
位运算符：用于对整数类型的二进制位进行操作，如按位与（&）、按位或（|）、按位异或（^）。
赋值运算符：用于将值赋给变量，如简单赋值（=）和复合赋值（+=、-=、*=、/=、%=、&=、|=、^= 和 &^=）。
其他运算符：如地址运算符（&）和指针解引用运算符（*）。
*/

// 算术运算符
func arithmeticOperators() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	// 如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	var n1 float32 = 10 / 4
	fmt.Println(n1) //结果是2.0
	var n2 float32 = 10.0 / 4
	fmt.Println(n2) //结果是2.5

	// go语言的自增自减只能当作一个独立语言使用，不能放在表达式中使用
	// 下面的代码会报错
	// var n3 int = 10
	// var n4 int = n3++ // 错误：cannot use n3++ as value
	// ++和--只能写在变量的后面
	n5 := 1
	// ++n5 报错
	n5++
}

// 关系运算符
func relationalOperators() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}
}

// 逻辑运算
func t5_test3() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

// 位运算
func t5_test4() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

// 赋值运算
func t5_test5() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6 行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)
}

// 其他运算符
func t5_test6() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("ptr 为 %v, *ptr 为 %d\n", ptr, *ptr)
	fmt.Printf("ptr 为 %x, *ptr 为 %d\n", ptr, *ptr)

}
func main() {
	arithmeticOperators()
}
