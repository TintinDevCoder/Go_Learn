package main

import "fmt"

// 一个指针变量指向了一个值的内存地址。
// var var_name *var-type
// var ip *int        /* 指向整型*/
// var fp *float32    /* 指向浮点型 */

// 指针基础用法
func pointer_test1() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	//Go 空指针
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr != nil { /* ptr 不是空指针 */
		fmt.Printf("ptr 不是空指针\n")
	}
	if ptr == nil { /* ptr 是空指针 */
		fmt.Printf("ptr 是空指针\n")
	}
}

// 指针数组
/**
以下声明了整型指针数组：
var ptr [MAX]*int;
*/
const MAX int = 3

func pointer_test2() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

// 指向指针的指针
/**
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
*/
// var ptr **int;
func pointer_test3() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d，地址：%x\n", a, &a)
	fmt.Printf("指针变量 *ptr = %d，内容：%x，地址：%x\n", *ptr, ptr, &ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d，内容：%x，地址：%x\n", **pptr, pptr, &pptr)
}

// 指针作为函数参数
// 函数定义的参数上设置为指针类型即可
func pointer_test4() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d，地址 : %x\n", a, &a)
	fmt.Printf("交换前 b 的值 : %d，地址 : %x\n", b, &b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap3(&a, &b)

	fmt.Printf("交换前 a 的值 : %d，地址 : %x\n", a, &a)
	fmt.Printf("交换前 b 的值 : %d，地址 : %x\n", b, &b)
}

// 传入地址，提取该地址的值，交换
func swap3(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func main() {
	pointer_test4()
}
