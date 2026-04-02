package main

import "fmt"

/*Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。*/
/*
for key, value := range oldMap {
    newMap[key] = value
}
*/
// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func slice_range() {
	// 遍历 pow 切片，i 是索引，v 是值
	for i, v := range pow {
		// 打印 2 的 i 次方等于 v
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func string_range() {
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}
func map_range() {
	// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
	map1 := make(map[int]float32)

	// 向 map1 中添加 key-value 对
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 遍历 map1，读取 key 和 value
	for key, value := range map1 {
		// 打印 key 和 value
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 遍历 map1，只读取 key
	for key := range map1 {
		// 打印 key
		fmt.Printf("key is: %d\n", key)
	}

	// 遍历 map1，只读取 value
	for _, value := range map1 {
		// 打印 value
		fmt.Printf("value is: %f\n", value)
	}
}

// range 遍历从通道接收的值，直到通道关闭
func channel_range() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	channel_range()
}
