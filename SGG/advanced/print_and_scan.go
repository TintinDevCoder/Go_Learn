package main

import (
	"fmt"
	"strconv"
)

// 终端获取用户输入
func main() {
	// fmt.Scanln 函数用于从标准输入读取一行文本，并将其存储在指定的变量中。
	// 它会等待用户输入，直到用户按下回车键为止。使用 fmt.Scanln 可以方便地获取用户输入的数据，并进行处理。
	var name string
	var age int
	var is string
	// 提示用户输入姓名和年龄
	println("请输入您的姓名：")
	fmt.Scanln(&name) // 使用 &name 获取变量的地址，以便 Scanln 可以将输入的值存储在 name 变量中
	println("请输入您的年龄：")
	fmt.Scanln(&age) // 使用 &name 获取变量的地址，以便 Scanln 可以将输入的值存储在 name 变量中
	println("请输入是否通过考试（true/false）")
	fmt.Scanln(&is)

	fmt.Println("您好，", name, "！您的年龄是", age, "岁。")

	isPassed, err := strconv.ParseBool(is)
	if err != nil {
		fmt.Println("输入的值无法转换为布尔值，请输入 true 或 false")
	} else {
		if isPassed {
			fmt.Println("恭喜您通过了考试！")
		}
	}

	// fmt.Scanf 函数用于从标准输入读取格式化的文本，并将其存储在指定的变量中。它允许你指定输入的格式，以便正确地解析用户输入的数据.
	var name2 string
	var age2 int
	var isPassed2 bool

	println("请输入您的姓名和年龄（格式：姓名 年龄）")
	fmt.Scanf("%s %d", &name2, &age2) // 使用 &name 获取变量的地址，以便 Scanf 可以将输入的值存储在 name 变量中
	println("请输入是否通过考试（true/false）")
	fmt.Scanf("%t", &isPassed2)

	fmt.Println("您好，", name2, "！您的年龄是", age2, "岁。")

	if isPassed2 {
		fmt.Println("恭喜您通过了考试！")
	}
}
