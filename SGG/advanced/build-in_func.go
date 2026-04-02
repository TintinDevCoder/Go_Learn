package main

// 内置函数
// Go语言中有一些内置函数，这些函数不需要导入任何包就可以直接使用，常见的内置函数有：
// len()：返回字符串、数组、切片、map等的长度
// new()：创建一个类型的零值，并返回指向它的指针
// cap()：返回切片、数组等的容量
// make()：创建切片、map、channel等
// append()：向切片添加元素
// copy()：复制切片
// delete()：删除map中的键值对
// panic()：触发一个运行时错误，导致程序崩溃
// recover()：从panic中恢复，捕获到panic后可以继续执行程序

func main() {
	// new
	p1 := new(int) // 创建一个int类型的零值，并返回指向它的指针
	*p1 = 10
	println(*p1) // 输出10

}
