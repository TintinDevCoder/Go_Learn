package main

import "fmt"

// go语言中的转义字符

func main() {
	// \t:表示制表符，用于排版对齐
	fmt.Println("123\t456")
	// \\:表示反斜杠本身
	fmt.Println("C:\\Users\\Admin")
	// \":表示双引号本身
	fmt.Println("She said, \"Hello!\"")
	// \n:表示换行符
	fmt.Println("Line 1\nLine 2")
	// \r:表示回车符，将光标移动到行首
	fmt.Print("Hello\rWorld\n") // 输出 "World" 覆盖 "Hello"
	// \b:表示退格符，将光标向后移动一个位置
	fmt.Print("Hello\b World\n") // 输出 "Hell World"
}
