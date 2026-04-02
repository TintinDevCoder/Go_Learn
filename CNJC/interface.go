package main

import (
	"fmt"
	"math"
)

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具。
接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
*/
/*
接口的特点
隐式实现：Go 中没有关键字显式声明某个类型实现了某个接口。只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。
接口类型变量：接口变量可以存储实现该接口的任意值。
接口变量实际上包含了两个部分：
动态类型：存储实际的值类型。  动态值：存储具体的值。
零值接口：接口的零值是 nil。一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。
空接口：定义为 interface{}，可以表示任何类型。

接口的常见用法
多态：不同类型实现同一接口，实现多态行为。
解耦：通过接口定义依赖关系，降低模块之间的耦合。
泛化：使用空接口 interface{} 表示任意类型。
*/
// 定义接口
type Shape2 interface {
	Area() float64
	Perimeter() float64
	SetRadius(r float64)
	SetRadius2(r float64)
}

// 定义一个结构体
type Circle2 struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle2) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle2) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func (c Circle2) SetRadius(r float64) {
	c.Radius = r
}
func (c *Circle2) SetRadius2(radius float64) {
	c.Radius = radius
	// 打印方法内部 c 的值，即它存储的地址 (简体中文注释)
	fmt.Printf("Inside method, pointer value: %p\n", c)
}

// 重要！
func interface_test1() {
	c := Circle2{Radius: 5}
	/*	var s Shape = c // 接口变量可以存储实现了接口的类型
		// s此处是 Circle 类型的一个副本，即复制出一个新的Circle，而后将s中的data指针指向它
		s.SetRadius(6)
		fmt.Println(c.Radius)
		fmt.Println("Area:", s.Area())
		fmt.Println("Perimeter:", s.Perimeter())*/

	var s2 Shape2 = &c // 使用指针类型实现接口
	// s2 此处是 *Circle 类型的一个副本，即复制出一个新的*Circle，而后将s2中的data指针指向它
	// 因此 s2 中存储的是 Circle 的地址，对 s2 的修改会影响到 c
	s2.SetRadius(6) //此处SetRadius是Circle，而s2是*Circle类型，所以会自动解引用
	// 自动将你的调用 s2.SetRadius(6) 转化为 (*s2).SetRadius(6)
	// 虽然编译器取出了指针指向的对象，但由于方法定义是 func (c Circle)
	// Go 依然会按照值传递的规则，为这个对象创建一个全新的副本（Copy）
	s2.SetRadius2(6) // 此处SetRadius2是*Circle，而s2是*Circle类型，所以不会自动解引用
	// 因为 SetRadius2 的接收者是指针类型，所以对 s2 的修改会直接影响到 c
	// 但是要注意，s2 仍然是一个接口类型的变量，它存储的是 *Circle 类型的一个副本
	// 此时因为存在了指针接收者的方法，因此无法用接口接收普通值类型的变量。即var s Shape = c会报错
	fmt.Println(c.Radius)
}

/*
空接口
空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。

任意类型都实现了空接口。
常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。
*/
func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}
func interface_test2() {
	printValue(42)          // int
	printValue("hello")     // string
	printValue(3.14)        // float64
	printValue([]int{1, 2}) // slice
}

// 类型断言用于从接口类型中提取其底层值
// value := iface.(Type)
func interface_test3() {
	var i interface{} = "hello"
	str := i.(string) // 类型断言
	fmt.Println(str)  // 输出：hello

	// 带检查的类型断言
	/*
		为了避免 panic，可以使用带检查的类型断言：

		value, ok := iface.(Type)
		ok 是一个布尔值，表示断言是否成功。
		如果断言失败，value 为零值，ok 为 false。
	*/
	var j interface{} = 42
	str, ok := j.(string)
	if ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}
}
func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}
func interface_test4() {
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})
}

// 复杂点的接口组合
type Readert interface {
	Read() string
}

type Writert interface {
	Write(data string)
}

type ReadWriter interface {
	Readert
	Writert
}

type File struct{}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}
func interface_test5() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello, Go!")
}

/*
动态值和动态类型
接口变量实际上包含了两部分：
动态类型：接口变量存储的具体类型。
动态值：具体类型的值。
*/
func interface_test6() {
	var i interface{} = 42
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}

func main() {
	interface_test6()
}
