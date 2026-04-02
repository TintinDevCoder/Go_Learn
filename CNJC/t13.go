package main

import "fmt"

func getAverage2(arr []int, size int) float32 {
	// 【阶段 1】
	// 此时 arr 是 main 中 balance 的一份 Header 拷贝（拷贝了指针、长度和容量）
	// arr[0] 和 balance[0] 指向的是同一个底层内存地址
	fmt.Printf("1函数内首元素: %d, 地址: %p\n", arr[0], &arr[0])

	// 【阶段 2】
	// 直接通过指针修改底层数组。因为没有搬家，这个操作会直接影响 main 函数里的 balance
	arr[0] = 0
	fmt.Printf("2函数内首元素: %d, 地址: %p\n", arr[0], &arr[0])

	// 【阶段 3：关键转折点】
	// append 发现当前切片的 cap（容量）不足以容纳新元素
	// 1. Go 会在堆内存中开辟一块新的、更大的空间
	// 2. 将旧数据拷贝到新空间
	// 3. arr 的内部指针指向这个新地址（即“搬家”）
	// 注意：此时 main 里的 balance 依然指向旧地址，两者正式“分家”
	arr = append(arr, 999)
	fmt.Printf("3函数内首元素: %d, 地址: %p\n", arr[0], &arr[0])

	// 【阶段 4】
	// 此时修改的是新家（新数组）里的数据，旧数组（main 里的 balance）不会受到任何影响
	arr[0] = 1
	fmt.Printf("4函数内首元素: %d, 地址: %p\n", arr[0], &arr[0])

	return 0
}
func t13_test1() {
	// 初始化一个长度为 5，容量也为 5 的切片
	balance := []int{1000, 2, 3, 17, 50}

	// 传递 balance 时，实际上是把 balance 的描述符（Header）复制了一份给 getAverage
	getAverage2(balance, len(balance))

	// 【阶段 5】
	// 验证结果：
	// 我们发现 balance[0] 变成了 0（受阶段 2 影响）
	// 但没有变成 1（阶段 4 的修改因为扩容搬家，改到新内存去了）
	fmt.Printf("5函数外首元素: %d, 地址: %p\n", balance[0], &balance[0])
}
