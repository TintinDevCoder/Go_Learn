package main

import (
	"fmt"
	"unsafe"
)

/*
*
变量和常量
*/
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	g1 int
	g2 bool
)

func main() {
	var s1 uint64 = 12345678901234567890
	var s2 int64 = -1234567890123456789

	fmt.Printf("s1: %d\n", s1)
	fmt.Printf("s2: %d\n", s2)
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	v1, v2, v3 := 1, 2, "three"
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	g1 = 1
	g2 = true
	fmt.Println(g1)
	fmt.Println(g2)

	//常量
	const LENGTH int = 10
	const length1, length2 = 20, 30

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	const (
		a1 = "abc"
		b1 = len(a1)
		c1 = unsafe.Sizeof(a1)
	)
	// 多个常量赋值，未给值的常量会与上面一行的常量值相同
	const (
		i3 = 1
		j3
		k3 = 2
		l3
	)
	fmt.Println("i=", i3)
	fmt.Println("j=", j3)
	fmt.Println("k=", k3)
	fmt.Println("l=", l3)
}
