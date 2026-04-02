package main

import "fmt"

// 数组
/**
全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
*/
var balance [10]float32

/* 初始化数组元素 */
func initArray() {
	var nums1 [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range nums1 {
		fmt.Println(i, v)
	}

	nums2 := [5]string{"Google", "Runoob", "Taobao", "Baidu", "Sina"}
	for i, v := range nums2 {
		fmt.Println(i, v)
	}
	// 初始化数组中 {} 中的元素个数不能大于 [] 中的数字
	// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
	balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	balance[4] = 60.0
	fmt.Println(balance)

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素
	// 将索引为 1 和 3 的元素初始化
	nums3 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(nums3)
}

// 访问数组元素
func visitArrays() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

// 多维数组
// var arrayName [ x ][ y ] variable_type
func multiArrays() {
	// 创建多维数组
	values1 := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}
	// 利用append函数添加行
	values1 = append(values1, row1)
	values1 = append(values1, row2)
	values1 = append(values1, row3)
	fmt.Println(values1)
	row1[1] = 20
	fmt.Println(values1)
	// 初始化多维数组
	a := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(a)
	// 访问多维数组
	// 增强for
	for index1, i := range a {
		for index2, j := range i {
			fmt.Printf("a[%d][%d]=%d \n", index1, index2, j)
		}
	}
	// 普通for
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

// 访问和修改数组元素
func visitAndEditArrays() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// 访问单个元素
	fmt.Println("第一行第二列:", matrix[0][1]) // 输出: 2
	fmt.Println("第三行第三列:", matrix[2][2]) // 输出: 9

	// 访问整行
	fmt.Println("第二行:", matrix[1]) // 输出: [4 5 6]

	// 遍历所有元素
	fmt.Println("\n遍历所有元素:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}
	matrix[0][0] = 10
	matrix[0][1] = 20
	matrix[1][0] = 30
	matrix[1][1] = 40
	fmt.Println("\n修改后的矩阵:", matrix)
	// 批量修改一行
	matrix[0] = [3]int{100, 200, 300}
	fmt.Println("修改第一行后:", matrix)
}

// 多维数组的常用操作
// 使用 range 遍历
func f1() {
	// 创建一个二维数组
	scores := [3][4]int{
		{85, 90, 78, 92},
		{88, 76, 95, 89},
		{92, 85, 88, 90},
	}
	fmt.Println("学生成绩表:")
	// 使用 range 遍历二维数组
	for i, row := range scores {
		fmt.Printf("学生 %d 的成绩: ", i+1)
		for j, score := range row {
			fmt.Printf("%d ", score)
			// 如果需要索引和值都使用
			_ = j // 避免未使用变量警告
		}
		fmt.Println()
	}
	// 只关心值，不关心索引
	total := 0
	count := 0
	for _, row := range scores {
		for _, score := range row {
			total += score
			count++
		}
	}
	fmt.Printf("\n平均分: %.2f\n", float64(total)/float64(count))
}
func f2() {
	matrix := [4][5]int{}
	// 获取行数
	rows := len(matrix)
	fmt.Println("行数:", rows) // 输出: 4
	// 获取第一行的列数（所有行长度相同）
	cols := len(matrix[0])
	fmt.Println("列数:", cols) // 输出: 5
	// 获取总元素数
	totalElements := rows * cols
	fmt.Println("总元素数:", totalElements) // 输出: 20
}
func f3() {
	// 创建两个相同的二维数组
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{1, 2}, {3, 4}}
	c := [2][2]int{{1, 2}, {3, 5}}

	// 数组可以直接比较（只有当维度完全相同时）会比较其中内容是否完全相同
	fmt.Println("a == b:", a == b) // 输出: true
	fmt.Println("a == c:", a == c) // 输出: false

	// 注意：不同维度的数组不能比较
	// d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(a == d)  // 编译错误：类型不匹配
}

// 向函数传递数组
// Go 语言中的数组是值类型，因此在将数组传递给函数时，实际上是传递数组的副本
// 如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组
func myFunc1(arr [10]int) {

}
func myFunc2(arr []int) {

}

// 值传递
func getAverage(arr [5]int, size int) float32 {
	fmt.Printf("函数内数组地址: %p\n", &arr)
	fmt.Printf("函数内首元素: %d, 地址: %p\n", arr[0], &arr[0])
	arr[0] = 0
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	return float32(sum) / float32(size)
}
func t10_test1() {
	balance := [5]int{1000, 2, 3, 17, 50}
	fmt.Printf("函数外数组地址: %p\n", &balance)
	fmt.Printf("函数外首元素: %d, 地址: %p\n", balance[0], &balance[0])
	getAverage(balance, len(balance))
}

// 引用传递
// 函数接受一个数组的指针作为参数
// 通过传递数组的指针，函数可以直接修改原始数组的值
// 值传递，传递的是数组的地址，go会通过新建一个新的指针变量来接收这个地址
func modifyArrayWithPointer(arr *[5]int) {
	fmt.Printf("函数内数组指针变量的地址: %p，存放的函数外数组地址: %p\n", &arr, arr)
	fmt.Printf("函数内首元素: %d, 地址: %p\n", (*arr)[0], &(*arr)[0])
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}
func t10_test2() {
	// 创建一个包含5个元素的整数数组
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array:", myArray)
	fmt.Printf("函数外数组地址: %p\n", &myArray)
	fmt.Printf("函数外首元素: %d, 地址: %p\n", myArray[0], &myArray[0])
	// 传递数组的指针给函数，可以修改原始数组的值
	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArrayWithPointer:", myArray)
}

func main() {
	t10_test2()
}
