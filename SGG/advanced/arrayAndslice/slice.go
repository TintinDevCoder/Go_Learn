package main

import "fmt"

// 切片是go语言中特有的数据类型
// 切片是对数组的抽象，切片是一个动态数组，切片是一个引用类型
// 切片的底层是一个数组，切片的长度可以动态变化，切片的容量是切片底层数组的长度
// 切片的长度是切片中元素的个数，切片的容量是切片底层数组的长度
// 切片的零值是nil，nil切片没有底层数组，长度和容量都是0
// 切片可以通过make函数创建，也可以通过字面量创建
// 切片可以通过切片表达式从数组或切片中创建，切片表达式的语法是：arrayOrSlice[low:high]，low表示切片的起始索引，high表示切片的结束索引，low和high都必须在原数组或切片的范围内，low默认为0，high默认为原数组或切片的长度
// 切片是引用类型，切片赋值
func initslice() {
	// 通过make函数创建切片，make函数的语法是：make([]T, len, cap)，T表示切片元素的类型，len表示切片的长度，cap表示切片的容量，cap默认为len
	slice1 := make([]int, 5)
	// 通过字面量创建切片，字面量的语法是：[]T{v1, v2, v3}，T表示切片元素的类型，v1、v2、v3表示切片元素的值
	slice2 := []int{1, 2, 3, 4, 5}
	// 通过切片表达式从数组中创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:4] // 从索引1到索引4（不包括索引4）创建切片
	// 通过切片表达式从切片中创建切片
	slice4 := slice2[1:4] // 从索引1到索引4（不包括索引4）创建切片
	// 切片的长度和容量
	println("slice1:", len(slice1), cap(slice1))
	println("slice2:", len(slice2), cap(slice2))
	println("slice3:", len(slice3), cap(slice3))
	println("slice4:", len(slice4), cap(slice4))
	// 切片的长度可以动态变化，切片的的底层是一个数组，切片的长度可以通过append函数动态增加，
	// append函数的语法是：append(slice, v1, v2, v3)，slice表示要追加元素的切片，v1、v2、v3表示要追加的元素，
	// 可以追加一个或多个元素，append函数会返回一个新的切片，如果切片的容量不足以容纳新的元素，append函数会自动扩容，扩容后的切片容量通常是原来容量的两倍
	// 切片的容量是切片底层数组的长度
	// 切片的容量代表了 在不重新分配内存的情况下，该切片可以增长到的最大长度
	slice4 = append(slice4, 6, 7, 8)
	println("slice4:", len(slice4), cap(slice4))
}
func neicunbuju() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[0:4]
	// 切片的地址 输出两个值相同，切片的底层数组就是这个数组，所以切片的地址就是这个数组的地址
	fmt.Printf("%p %p\n", &arr, slice3) // 数组的地址
	// slice中存放了三个值：切片的长度、切片的容量、切片底层数组的地址，本身就是一种结构体
	fmt.Println(arr[0])
	slice3[0] = 8
	fmt.Println(arr[0]) // 切片修改了底层数组的值，所以数组的值也被修改了

	// make创建切片也会创建数组，是由切片在底层进行维护的
	slice4 := make([]int, 3, 5)
	slice4[0] = 1
}
func slicePro() {
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1[2]) // 在长度范围内，可以访问到
	// fmt.Println(slice1[3]) // 在长度范围外，访问不到，会发生数组越界错误

	// 切片还可以继续切片
	slice2 := slice1[0:2] // 从索引0到索引2（不包括索引2）创建切片
	fmt.Println(slice2)
	slice2[0] = 8
	fmt.Println(slice1) // 切片修改了底层数组的值，所以切片1的值也被修改了

	// append函数可以动态增加切片的长度，如果切片容量不足，就会动态扩容，扩容后的切片容量通常是原来容量的两倍
	println("slice2:", len(slice2), cap(slice2))
	fmt.Println(slice2)
	slice2 = append(slice2, 1, 2, 3, 4) // 追加元素，返回一个新的切片
	println("slice2:", len(slice2), cap(slice2))
	slice2[0] = 9
	fmt.Println(slice1[0]) // slice2此时已经不再和slice1共享底层数组了，所以修改slice2不会影响slice1

	// append 也可以追加切片
	fmt.Println("slice2:", slice2)
	slice2 = append(slice2, slice2...) // 追加切片，返回一个新的切片
	fmt.Println("slice2:", slice2)

	slice3 := make([]int, 2, 2)
	fmt.Printf("%p, %v\n", &slice3[0], slice3) // 0xc00010c140, [0 0]
	slice3 = append(slice3, 1, 2, 3)
	fmt.Printf("%p, %v\n", &slice3[0], slice3) // 0xc000128030, [0 0 1 2 3]
	// 通过上面代码可以知道，此时append发送扩容，切片底层数组是新创建的，因此第一个元素的地址发送变化

	slice4 := make([]int, 2, 5)
	fmt.Printf("%p, %v\n", &slice4[0], slice4) // 0xc000128060, [0 0]
	slice4 = append(slice4, 1, 2, 3)
	fmt.Printf("%p, %v\n", &slice4[0], slice4) // 0xc000128060, [0 0 1 2 3]
	// 此时append没有发生扩容，切片底层数组没有新创建，因此第一个元素的地址没有发生变化
	// 切片的地址不会发生变化，但切片所指向的数组地址会发生变化

	// 切片的拷贝
	slice5 := []int{1, 2, 3}
	slice6 := make([]int, 10)
	copy(slice6, slice5) // 必须都为切片类型
	fmt.Println(slice6)  // 前面三个元素是从slice5中拷贝的，后面都为0
	// 如果拷贝的目标元素小于源元素，那么只会拷贝目标元素个数的元素
	slice7 := make([]int, 2)
	copy(slice7, slice5)
	fmt.Println(slice7) // [1 2]
}
func testsl(slice []int) {
	slice[0] = 10
}
func main() {
	slicePro()
	s := make([]int, 3)
	s[0] = 1
	fmt.Println(s[0]) // 1
	testsl(s)
	fmt.Println(s[0]) // 10

}
