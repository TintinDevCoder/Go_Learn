package main

import "fmt"

func main() {
	// 十进制输出二进制
	var i int = 5
	fmt.Printf("%b\n", i) // 输出二进制

	// 八进制，前面加0表示八进制
	var j int = 010
	fmt.Println(j)

	// 十六进制，前面加0x表示十六进制
	var k int = 0x020 // 16 * 2 + 0 = 32
	fmt.Println(k)

	// 二进制，前面加0b表示二进制
	var m int = 0b1010
	fmt.Println(m)

	a := -1 << 2
	fmt.Println(a) // 结果是0，1的二进制表示是0001，右移两位后变成0000，即0
}
