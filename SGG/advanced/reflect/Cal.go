package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal *Cal) GetSub(name string) string {
	return fmt.Sprintf("%s完成了减法操作，%d - %d = %d", name, cal.Num1, cal.Num2, cal.Num1-cal.Num1)
}
func main() {
	c := Cal{Num1: 1, Num2: 2}
	// 使用反射遍历Cal结构体的所有字段信息
	ctype := reflect.TypeOf(&c)
	cvalue := reflect.ValueOf(&c)
	numF := cvalue.Elem().NumField()
	for i := 0; i < numF; i++ {
		field := ctype.Elem().Field(i)
		value := cvalue.Elem().Field(i)
		fmt.Printf("字段名称: %v, 字段类型: %v, 字段值: %v\n", field.Name, field.Type, value)
	}
	method := cvalue.Method(0)
	result := method.Call([]reflect.Value{reflect.ValueOf("tom")})
	fmt.Println(result)
}
