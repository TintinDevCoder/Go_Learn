package main

// string底层是一个byte数组，因此也可以进行切片处理
// 字符串保存了一个地址ptr，指向byte数组首地址，字符串的长度len，字符串的容量cap等于字符串的长度
func main() {
	str := "hello world"
	// 切片表达式从字符串中创建切片
	slice1 := str[0:5]  // 从索引0到索引5（不包括索引5）创建切片
	slice2 := str[6:11] // 从索引6到索引11（不包括索引11）创建切片
	println(slice1)     // 输出：hello
	println(slice2)     // 输出：world

}
