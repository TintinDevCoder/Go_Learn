package main

import "fmt"

func main() {
	// 二维数组
	var arr [4][6]int
	fmt.Println(arr)
	// 遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = i + j
		}
	}
	fmt.Println(arr)
	for i, v1 := range arr {
		for j, v2 := range v1 {
			fmt.Printf("arr[%d][%d] = %d ", i, j, v2)
		}
		fmt.Println()
	}

	// var arr [4][6]int
	// 二维数组声明后，内存中会在一级数组中分配4个元素，每个元素都是一个指针，指向一个长度为6的数组的首地址，二维数组的内存布局是连续的
	fmt.Printf("%p %p %p %p %p\n", &arr, &arr[0], &arr[0][0], &arr[0][1], &arr[0][5]) // 二维数组的地址和一级数组的地址相同
	// 0xc000126000 0xc000126000 0xc000126000 0xc000126008 0xc000126028
	fmt.Printf("%p %p %p\n", &arr[1], &arr[1][0], &arr[1][1])
	// 0xc000126030 0xc000126030 0xc000126038

	// 四种初始化方法
	var arr1 [4][6]int
	arr2 := [4][6]int{}
	arr3 := [...][6]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}}
	arr4 := [4][6]int{{1}, {7}, {13}, {19}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
