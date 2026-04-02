package main

import "fmt"

// 数组可以存放多个同一类型数据，数组是一种数据类型
// 数组是值类型
func initAndVisitArr() {
	// 定义一个数组
	var hens1 [6]float64
	// 数组元素赋值
	hens1[0] = 1.2
	hens1[1] = 1.3
	hens1[2] = 1.4
	hens1[3] = 1.5
	hens1[4] = 1.6
	hens1[5] = 1.7

	// 定义一个数组并初始化
	var hens2 = [6]float64{1.2, 1.3, 1.4, 1.5, 1.6, 1.7}

	// 定义一个数组并初始化，编译器自动推断数组长度
	var hens3 = [...]float64{1.2, 1.3, 1.4, 1.5, 1.6, 1.7}

	// 遍历数组
	for i := 0; i < len(hens2); i++ {
		fmt.Print(hens2[i], " ")
	}
	fmt.Println()

	for i, v := range hens3 {
		fmt.Print("索引是", i, "值是", v)
	}
	fmt.Println()
	// 数组是值类型，数组赋值会复制整个数组
	hens4 := hens3
	hens4[0] = 2.0
	fmt.Println("hens3:", hens3)
	fmt.Println("hens4:", hens4)
}

// 数组声明后长度是固定的，不能动态变化
// 数组声明必须带长度，长度是数组类型的一部分，不带长度的数组是切片
// 数组中的元素可以是值类型也可以是引用类型，数组中的元素是连续存储的
// 数组长度固定后，会初始化值
// 数组的地址是数组第一个元素的地址，数组元素是连续存储的，所以数组元素的地址也是连续的
// 使用数组必须在其长度范围内访问元素，否则会发生数组越界错误
// go的数组是值类型，数组赋值会复制整个数组，修改一个数组不会影响另一个数组
// 如果想用引用类型修改原本的数组，可以使用引用传递
// 在go中，长度是数组类型的一部分，所以不同长度的数组是不同的类型，不能直接赋值
func test01(arr [3]int) {
	arr[0] = 8
}

// 此时arr是一个指针，指向一个长度为3的int数组，修改arr[0]会修改原本的数组
func test02(arr *[3]int) {
	arr[0] = 8
}
func reverseArr(arr *[10]int) {
	leng := len(arr)
	for i := 0; i < leng/2; i++ {
		arr[i], arr[leng-i-1] = arr[leng-i-1], arr[i]
	}
}
func main() {
	var arr [6]float64
	fmt.Printf("%p %p", &arr, &arr[1])

	var arr2 [3]int
	arr2[0] = 1
	test01(arr2)
	fmt.Println(arr2[0])
	test02(&arr2)
	fmt.Println(arr2[0])
}
