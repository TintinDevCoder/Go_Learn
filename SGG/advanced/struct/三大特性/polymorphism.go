package main

import "fmt"

// 多态：多态是面向对象编程中的一个重要概念，它允许一个接口有多个实现，从而使得同一操作可以作用于不同的对象上
// 可以按照统一的接口来调用不同的实现，这时接口变量就呈现出不同的形态。
type Usbp interface {
	Start()
	Stop()
}
type Phonep struct {
	Name string
}

// Phonep实现Usbp接口
func (p Phonep) Start() {
	fmt.Println("Phone start")
}
func (p Phonep) Stop() {
	fmt.Println("Phone stop")
}

// 让Camera实现Usb接口
type Camerap struct {
	Name string
}

// Phone实现Usb接口
func (c Camerap) Start() {
	fmt.Println("Camera start")
}
func (c Camerap) Stop() {
	fmt.Println("Camera stop")
}

// Computer结构体
type Computerp struct {
	Name string
}

// 多态参数的体现：
func (c Computerp) Working(usb Usbp) { // usb变量会根据传入的实例，判断到底是Phone还是Camera，从而调用对应的方法
	usb.Start()
	fmt.Printf("Computer %s is working with %s\n", c.Name, usb)
	usb.Stop()
}

// 多态数组的体现：
func main() {
	usbps := make([]Usbp, 0)                      // 创建一个Usbp类型的切片
	usbps = append(usbps, Phonep{Name: "iPhone"}) // 将Phonep类型的实例添加到切片中
	usbps = append(usbps, Camerap{Name: "Canon"}) // 将Camerap类型的实例添加到切片中
	for _, usb := range usbps {                   // 遍历切片，调用Computerp的Working方法
		computer := Computerp{Name: "MacBook"}
		computer.Working(usb) // 传入不同的实例，调用不同的方法
	}

}
