package main

import (
	"encoding/json"
	"fmt"
)

// go语言的面向对象编程，并不是纯粹的面向对象语言。
// go中没有类，结构体就是对应的对象，用来实现oop特性
// go语言的面向对象编程去掉了oop语言的继承、方法重载、构造函数和析构函数等
// go仍有继承特性、封装、多态特性，只是是想方式与其他oop语言不同，比如继承是通过匿名字段实现的，封装是通过首字母大小写实现的，多态是通过接口实现的
// go通过接口关联不同类型的对象，接口是一种抽象类型，接口定义了一组方法，任何类型只要实现了接口定义的方法，就可以被视为实现了该接口，耦合性低，更为灵活
// 将一类事务特征提取出来，形成一个新的数据结构，就是结构体
// 通过结构体，可以创建多个变量
type Cat struct {
	Name  string
	Age   int
	Ctype string
	Color string
}

// 结构体的字段 = 属性 = field
// 字段的类型可以为任意类型，包括基本类型、数组、切片、map、函数、接口、结构体等
// 创建结构体变量后，若没有给变量赋值，则会对应一个零值
// 不同结构体变量的字段相互独立
func main() {
	cat1 := Cat{Name: "uu", Age: 2, Ctype: "中华田园猫", Color: "黄色"}
	fmt.Println(cat1)

	// 结构体是值类型，结构体变量存储的是结构体的值，结构体变量之间是相互独立的，修改一个结构体变量不会影响另一个结构体变量
	fmt.Printf("%p %p %p\n", &cat1, &cat1.Name, &cat1.Age) // 0xc00009e040 0xc00009e040 0xc00009e050

	// 不同结构体变量的字段是独立的，一个结构体变量字段的改变不影响另一个，结构体是值类型
	cat2 := cat1
	// 在内存中将完整复制一个cat1给cat2，cat2的地址和cat1不同，cat2.Name的地址和cat1.Name不同，cat2.Age的地址和cat1.Age不同
	fmt.Printf("%p %p %p\n", &cat2, &cat2.Name, &cat2.Age) // 0xc00009e060 0xc00009e060 0xc00009e070
	cat2.Name = "vv"
	cat2.Age = 3
	fmt.Println(cat1) // {uu 2 中华田园猫 黄色}
	fmt.Println(cat2) // {vv 3 中华田园猫 黄色}

	// 声明方式
	cat3 := Cat{}
	cat4 := Cat{Name: "ww"}

	var p *Cat = new(Cat)
	(*p).Name = "xx" //(*p).Name = "xx"可以等嫁与p.Name = "xx"，go语言中会自动将p.Name转换为(*p).Name

	var p2 *Cat = &Cat{}
	fmt.Println(cat3)
	fmt.Println(cat4)
	fmt.Println(*p)
	fmt.Println(*p2)

	neicun()

	// 结构体之间如果想要相互转换，必须满足两个条件：
	// 1.两个结构体的字段数量必须相同
	// 2.两个结构体的字段名称和类型必须相同，字段的顺序也必须相同
	type Dog struct {
		Name  string
		Age   int
		Ctype string
		Color string
	}
	var dog1 = Dog{Name: "dd", Age: 4, Ctype: "中华田园猫", Color: "黄色"}
	var cat5 = Cat(dog1)
	fmt.Println(cat5)

	// 结构体可以用type进行重新定义
	type catt Cat
	// catt和Cat是两种不同的类型，虽然它们的字段相同，但它们是不同的类型，不能直接赋值，需要进行类型转换
	// var catt1 catt = Cat{} // 报错
	// 此时需要强转
	catt1 := catt(Cat{Name: "1"})
	fmt.Println(catt1)

	// 结构体的每个字段上都可以写上一个tag，tag可以通过反射机制获取到，常见的用途是用来指定json序列化和反序列化时的字段名称
	// 结构体字段首字母小写，会使得别的包无法访问，也无法处理为json, 因为相当于在其他包访问结构体字段时会报错，无法访问到该字段，所以无法处理为json
	// 可以用tag解决问题
	type User struct {
		Name string `json:"name"` // 通过tag指定json序列化和反序列化时的字段名称为name
		Age  int    `json:"age"`
	}
	var user1 = User{Name: "dd", Age: 4}
	// 序列化为json
	jsonStr, err := json.Marshal(user1) // 返回一个byte切片
	if err == nil {
		fmt.Println(string(jsonStr)) // {"name":"dd","age":4}
	}
}
func neicun() {
	var p1 = Cat{Name: "dd", Age: 4, Ctype: "中华田园猫", Color: "黄色"}
	var p2 *Cat = &p1
	p2.Name = "cc"
	fmt.Printf("p1的值：%v 地址：%p\n", p1, &p1)
	fmt.Printf("p2指向的内容的值：%v p2的内容：%p p2的地址：%p\n", *p2, p2, &p2)

	// 结构体的字段在内存中是连续分布的
	fmt.Printf("p1.Name的地址：%p p1.Age的地址：%p p1.Ctype的地址：%p p1.Color的地址：%p\n", &p1.Name, &p1.Age, &p1.Ctype, &p1.Color)

}
