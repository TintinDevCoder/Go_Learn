package main

import "fmt"

// 继承：继承是面向对象编程中的一个重要概念，它允许一个新的类（子类）从一个已有的类（父类）继承属性和方法。通过继承
// 子类可以重用父类的代码，并且可以添加自己的属性和方法来扩展功能。

// 嵌套匿名结构体：一个结构体可以嵌套另一个匿名结构体，这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性
type Goods struct {
	name  string
	price int
}

// get和set方法
func (g *Goods) GetName() string {
	return g.name
}
func (g *Goods) SetName(name string) {
	g.name = name
}
func (g *Goods) GetPrice() int {
	return g.price
}
func (g *Goods) SetPrice(price int) {
	g.price = price
}

// Goods的特有方法
func (g *Goods) show() {
	fmt.Println("商品名称：", g.name, "价格：", g.price)
}

// 1、结构体可以使用嵌套匿名结构体的所有字段和方法，不管是否首字母大写
type Book struct {
	Goods  // 嵌入匿名结构体
	Writer string
}

// Book独有的方法
func (b *Book) readbook() {
	fmt.Println("正在阅读", b.GetName(), "作者是", b.Writer)
}

// 也可以嵌入匿名结构体的指针
type Book2 struct {
	*Goods // 嵌入匿名结构体指针
	Writer string
}

// 如果结构体嵌套了一个有名的结构体，那么访问有名结构体的字段和方法时，必须通过有名结构体的名字来访问，而不能直接通过外层结构体来访问
type Car struct {
	good  *Goods // 嵌入有名结构体
	Brand string
	Model string
}

func main() {
	book := Book{Writer: "张三"}
	book.Goods.price = 100
	book.Goods.name = "Go语言编程"
	fmt.Println(book)
	book.Goods.show()
	book.readbook()
	
	// 2、匿名结构体字段访问可以简化
	// 当直接通过book访问匿名结构体的字段和方法时，执行流程如下：
	// 1.首先在Book结构体中查找是否有该字段或方法，如果有则直接访问
	// 2.如果Book结构体中没有该字段或方法，则继续在匿名结构体Goods中查找，如果有则访问，Goods中再查找自己的匿名结构体
	// 3.如果匿名结构体Goods中也没有该字段或方法，则访问失败，编译器会报错
	// 如果存在相同的字段和方法，会采用就近访问的原则，如果希望访问某个匿名结构体的字段，则可以通过结构体名访问
	book.SetName("Go")
	book.SetPrice(200)
	book.show()

	// 嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体的值
	b2 := Book{Goods: Goods{name: "Go语言编程", price: 100}, Writer: "张三"}
	fmt.Println(b2)

	// 匿名结构体指针的访问
	book2 := Book2{Goods: &Goods{name: "Go语言编程2", price: 100}, Writer: "张三2"}
	fmt.Println(book2)        // {0xc0000080a8 张三2}
	fmt.Println(*book2.Goods) // {Go语言编程2 100}
	book2.show()

	// 有名结构体嵌套
	car := Car{good: &Goods{name: "汽车", price: 10000}, Brand: "宝马", Model: "X5"}
	fmt.Println(car)
	fmt.Println(*car.good)
}
