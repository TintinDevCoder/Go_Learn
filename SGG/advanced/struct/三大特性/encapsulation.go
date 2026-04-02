package main

// golang有三大特性：封装、继承、多态，实现方式和其他oop语言不一样
// 封装：将数据和操作数据的函数绑定在一起，并隐藏内部实现细节，只暴露必要的接口给外部使用。
// 封装的具体步骤：
// 1. 定义一个结构体，包含需要封装的数据字段，首字母小写，表示私有字段
// 2. 定义一个工厂模式的构造函数，返回结构体的实例，首字母大写，表示公共函数
// 3. 定义一些方法，操作结构体的字段，首字母大写，表示公共方法，如：Setxxx、Getxxx等

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}
