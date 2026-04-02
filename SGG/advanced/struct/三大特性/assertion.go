package main

import "fmt"

// 类型断言：类型断言是go语言中的一个重要特性，它允许我们在运行时检查一个接口变量的动态类型，并将其转换为具体的类型
type Point struct {
	x int
	y int
}

func (p *Point) showAdd() {
	fmt.Println(p.x + p.y)
}
func main() {
	// 类型断言如果类型不匹配，就会报panic错误，因此在使用类型断言时，我们需要确保接口变量的动态类型与我们要转换的类型匹配，否则就会引发运行时错误
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	//b = a // 直接将接口变量a赋值给b会编译错误，因为a是一个接口类型，而b是一个具体的类型，因此需要使用类型断言来将a转换为Point类型
	b = a.(Point) // 使用类型断言将a转换为Point类型，如果a的动态类型不是Point类型，则会引发运行时错误
	fmt.Println(b)
	b.showAdd() // 通过类型断言将接口变量a转换为Point类型后，我们就可以调用Point类型的方法了

	// 其他案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y := x.(float32) // 将接口变量x转换为float32类型
	fmt.Printf("y的类型：%T，值：%v\n", y, y)

	// 在进行断言时，带上检测机制，如果成功就ok，否则就会返回一个错误
	var z interface{}
	z = "hello"
	// 带上检测机制进行类型断言
	if str, ok := z.(string); ok {
		fmt.Printf("z的类型：%T，值：%v\n", str, str)
	} else {
		fmt.Println("类型断言失败，z不是string类型")
	}
}
