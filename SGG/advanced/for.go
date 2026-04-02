package main

import "fmt"

// for1
func t7_test1() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	for sum = 1; sum <= 100; sum++ {
		fmt.Printf("%d ", sum)
	}
	fmt.Println()
	// For-each range 循环 可以对字符串、数组、切片等进行迭代输出元素
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}

// 循环嵌套
func t7_test2() {
	for i := 1; i < 100; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("i: %d j: %d i + j = %d\n", i, j, i+j)
			fmt.Println()
		}
	}
}

func t7_test3() {
	// 传统的对字符串的遍历是按字节遍历的
	// 中文占三个字节，因此在数组中也占3个位置
	str := "北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d, %c\n", i, str[i])
	}
	// range范围遍历可以正常，是按照字符的方式遍历的
	for i, v := range str {
		fmt.Printf("%d, %c\n", i, v)
	}

	// 将字符串转换为rune切片，rune是Go语言中的一种数据类型，表示一个Unicode码点，可以正确处理中文字符
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d, %c\n", i, str2[i])
	}
}
func main() {
	t7_test3()
}
