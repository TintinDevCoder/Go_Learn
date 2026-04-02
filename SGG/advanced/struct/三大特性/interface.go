package main

import "fmt"

// 接口：接口是go语言中实现多态的方式，接口是一种抽象类型。高内聚低耦合
// 接口定义了一组方法，任何类型只要实现了接口定义的方法，就可以被视为实现了该接口，耦合性低，更为灵活
// 实现接口，就是实现了接口定义的方法，go语言中没有显式的implements关键字
// 接口不允许有字段，接口只能定义方法，接口的方法没有方法体，接口的方法只能声明方法的签名
// 只有实现了接口的所有方法，才算实现了接口，接口的实现是隐式的，不需要显式声明

// interface默认是指针类型，必须初始化后才能使用，接口变量默认值为nil，接口变量可以指向任何实现了该接口的类型的变量

type Usb interface {
	// 声明了两个方法，任何类型只要实现了这两个方法，就可以被视为实现了Usb接口
	Start()
	Stop()
}
type Phone struct {
	Name string
}

// Phone实现Usb接口
func (p Phone) Start() {
	fmt.Println("Phone start")
}
func (p Phone) Stop() {
	fmt.Println("Phone stop")
}

// 让Camera实现Usb接口
type Camera struct {
	Name string
}

// Phone实现Usb接口
func (c Camera) Start() {
	fmt.Println("Camera start")
}
func (c Camera) Stop() {
	fmt.Println("Camera stop")
}

// Computer结构体
type Computer struct {
	Name string
}

// Computer的方法Working,接收一个Usb接口类型的参数，表示Computer可以使用任何实现了Usb接口的类型作为参数，从而实现了多态
func (c Computer) Working1(usb Usb) {
	// 通过接口调用Usb接口的方法，实际调用的是具体类型的方法
	usb.Start()
	fmt.Printf("Computer %s is working with %s\n", c.Name, usb)
	usb.Stop()
}

// golang不关心一个类型是否实现了某个接口，只要这个类型实现了接口定义的所有方法，就可以被视为实现了该接口，这种特性被称为鸭子类型（Duck Typing）。
// 鸭子类型是一种动态类型系统的特性，它允许一个对象在运行时被视为某个类型，只要它具有该类型所需的方法和属性。
// 因此，Phone和Camera也实现了Usb2接口
type Usb2 interface {
	Start()
	Stop()
}

func (c Computer) Working2(usb Usb2) {
	// 通过接口调用Usb接口的方法，实际调用的是具体类型的方法
	usb.Start()
	fmt.Printf("Computer %s is working with %s~~\n", c.Name, usb)
	usb.Stop()
}

type Usb3 interface {
	Start()
	Stop()
	Test()
}

func (c Computer) Working3(usb Usb3) {
	// 通过接口调用Usb接口的方法，实际调用的是具体类型的方法
	usb.Start()
	fmt.Printf("Computer %s is working with %s~~\n", c.Name, usb)
	usb.Stop()
}

type MyInt int

func (m MyInt) Start() {
	fmt.Println("MyInt start")
}
func (m MyInt) Stop() {
	fmt.Println("MyInt stop")
}
func test() {
	phone := Phone{Name: "iphone17"}
	camera := Camera{Name: "Canon EOS R5"}
	computer := Computer{Name: "MacBook Pro"}

	computer.Working1(phone)
	computer.Working1(camera)

	computer.Working2(phone)
	computer.Working2(camera)

	// 接口有三个方法，Phone和Camera只实现了两个，因此并没有实现Usb3接口，编译报错
	//computer.Working3(phone)
	//computer.Working3(camera)

	// 接口本身是不能声明变量的，但是可以指向一个实现了接口的类型的变量
	var u Usb = phone // 接口变量u指向一个实现了Usb接口的类型的变量phone
	fmt.Println(u)

	// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

	var u2 Usb = MyInt(10) // 接口变量u2指向一个实现了Usb接口的类型的变量MyInt(10)
	fmt.Println(u2)
}

// 接口中所有方法都没有方法体，即都是没有实现的方法
// 一个自定义类型可以实现多个接口，一个接口也可以被多个自定义类型实现
// 一个接口可以继承多个别的接口，这时如果要实现这个接口，就必须实现它继承的所有接口的方法
type A interface {
	a()
	b()
}
type B interface {
	c()
	d()
}
type C interface {
	A
	B
	e()
	f()
}
type D struct {
}

func (d D) a() {
	fmt.Println("D a")
}
func (d D) b() {
	fmt.Println("D b")
}
func (d D) c() {
	fmt.Println("D c")
}
func (d D) d() {
	fmt.Println("D d")
}
func (d D) e() {
	fmt.Println("D e")
}
func (d D) f() {
	fmt.Println("D f")
}
func t(c C) {
	fmt.Println("c is ", c)
}
func test1() {
	dd := D{}
	t(dd)
}

// 一个结构体可以实现多个接口
type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}
type Monster struct {
}

// 实现AInterface接口
func (*Monster) Say() {
	fmt.Println("Monster say")
}

// 实现BInterface接口
func (*Monster) Hello() {
	fmt.Println("Monster hello")
}
func test2() {
	monster := Monster{}
	var a AInterface = &monster // Monster实现了AInterface接口
	var b BInterface = &monster // Monster实现了BInterface接口
	a.Say()
	b.Hello()
}

// 空接口interface{} 没有任何方法，因此所有类型都实现了空接口，即可以把任何变量赋给空接口类型
func test3() {
	var i interface{} = 10
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = true
	fmt.Println(i)
}

// 接口继承其他多个接口的时候，不允许出现同名但签名不同（例如返回值不同或参数不同）的方法，否则编译报错
/*
type t1 interface {
	a(int)
}
type t2 interface {
	a()
}
type t3 interface {
	t1
	t2
}
*/
type u1 interface {
	Say()
}
type fu1 struct {
}

// 此时实现u1接口的是fu1的接口类型，不是fu1类型
func (f *fu1) Say() {
	fmt.Println("Say")
}
func test4() {
	var f1 fu1 = fu1{}
	//var u1 u1 = f1 // 此时会编译错误，因为不是fu1类型实现了u1接口，而是*fu1类型实现了u1接口，因此需要将f1的地址赋给u1变量
	var u2 u1 = &f1 // 此时编译成功，因为*fu1类型实现了u1接口
	u2.Say()
}
func main() {
	test()
	fmt.Println("-----------------------------")
	test3()
	fmt.Println("-----------------------------")
	test4()

}
