package main

import (
	"SGG/advanced/struct/model"
	"fmt"
)

// 工厂模式：定义一个创建对象的接口，让子类决定实例化哪个类。工厂方法使一个类的实例化延迟到其子类。
// 当结构体首字母是小写，但又想在其他包中使用时，可以通过工厂模式来实现。

func main() {
	// 通过工厂模式创建stu对象
	stu1 := model.NewStu("Alice", 90.5)
	fmt.Println(stu1)
}
