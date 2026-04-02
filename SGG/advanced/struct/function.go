package main

import "fmt"

// 结构体可以有自己的方法，方法和数据类型相互绑定
type A struct {
	Num int
}

// 在通过一个变量去调用方法时，调用机制和函数一样
// 不一样之处在于，变量调用方法时，该变量本身也会作为一个参数传递到方法
// 如果变量是值类型，那么传递的是该变量的副本，如果变量是指针类型，那么传递的是该变量的地址

// 方法的定义：func (接收者) 方法名(参数列表) (返回值列表) {函数体}
// 此处接收者没有指定为指针类型，所以方法内部对接收者的修改不会影响到外部调用者，此时是传入了一个副本，方法内部修改的是副本的值，不会影响到外部调用者的值
func (a A) test1(n int) {
	fmt.Println("test1", a.Num)
	a.Num = n
	fmt.Println("test1", a.Num)
}

// 首字母小写的方法只能在当前包内被调用，首字母大写的方法可以在当前包内和其他包内被调用
func (a *A) test2() {
	fmt.Println("test2", a.Num)
	a.Num = 100
	fmt.Println("test2", a.Num)
}

// 如果一个结构体实现了String方法，那么在打印该结构体变量时，会自动调用String方法，返回一个字符串来表示该结构体变量
func (a A) String() string {
	return fmt.Sprintf("A{Num: %d}", a.Num)
}
func main() {
	a := A{Num: 10}
	a.test1(100)   // test 10
	fmt.Println(a) // {10}

	(&a).test2() // 因为方法接收者是指针类型，所以需要传入一个指针类型的变量，此处通过&a将a的地址传入方法
	// 此处也可以写成a.test2()，go语言会自动将a转换为&a传入方法
	fmt.Println(a) // {100}

	(&a).test1(15) // 从形式上看是传了地址
	// 但编译器做了优化，会自动将&a转换为a传入方法，因为方法接收者是值类型，所以传入的是a的副本
	// 方法内部修改的是副本的值，不会影响到外部调用者的值
	fmt.Println(a) // {100}
	// 不管调用形式如何，真正决定时只拷贝还是地址拷贝，看这个方法是和那个类型绑定
}

type Caculator struct {
	Num int
}

func (this *Caculator) add(num int) {
	this.Num += num
}
func (this *Caculator) sub(num int) {
	this.Num -= num
}
func (this *Caculator) mult(num int) {
	this.Num *= num
}
func (this *Caculator) div(num int) {
	this.Num /= num
}
