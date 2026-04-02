package main

import (
	"fmt"
	"reflect"
)

// 反射（Reflection）是 Go 语言中的一种强大机制，允许程序在运行时检查和操作类型和值。通过反射，程序可以动态地获取变量的类型信息、修改变量的值，以及调用函数等。这使得反射成为实现一些高级功能（如序列化、依赖注入、ORM 等）的重要工具。
// 反射可以在运行时动态获取变量的各种信息，比如变量的类型、类别
// 如果是结构体变量，还可以获取到结构体本身的信息，比如结构体的字段、方法
// 通过反射，可以修改变量的值，可以调用关联的方法
// 反射中，变量、interface{}、类型、值等都是通过 reflect 包中的方法相互转换的
// 比如 reflect.TypeOf() 和 reflect.ValueOf() 等函数可以将变量转换为 reflect.Type 和 reflect.Value 类型的值
// 反过来也可以通过 reflect.Value 的方法将其转换回原始类型的值。

// reflect 包提供了许多函数和类型来实现反射功能，其中最常用的函数是 reflect.TypeOf() 和 reflect.ValueOf()。
// 1. reflect.TypeOf()：用于获取一个变量的类型信息，返回一个 reflect.Type 类型的值。
// 2. reflect.ValueOf()：用于获取一个变量的值信息，返回一个 reflect.Value 类型的值。

// type和kind可能相同，也可能不同
// type是表示唯一的类型信息，比如自己声明的main.Stu类型和main.Student是不同的类型
// kind是表示类型的类别，比如main.Stu和main.Student的kind都是struct

// 对基本数据类型的反射
func reflectTest01() {
	num := 100
	// 通过反射获取传入变量的type，kind，value值信息
	// TypeOf()函数返回一个reflect.Type类型的值，表示变量的类型信息，比如int、string、struct等
	// reflect.Type 是 Go 语言中用于描述类型的接口，它提供了许多方法来获取类型的详细信息，比如名称、大小、对齐方式等。
	// 这个的类型是reflect.Tpye
	nt := reflect.TypeOf(num)
	fmt.Println("类型: ", nt)
	// Kind()方法返回一个reflect.Kind类型的值，表示变量的类别，比如int、string、struct等
	nk := nt.Kind()
	fmt.Println("类别: ", nk)

	// ValueOf()函数返回一个reflect.Value类型的值，表示变量的值信息，可以通过这个值来获取变量的具体值
	nv := reflect.ValueOf(num)
	nk = nv.Kind()

	// 需要通过reflect.Value的Int()方法来获取num的整数值
	// 必须是int类型的变量才能调用Int()方法，否则会报错
	num2 := 1 + nv.Int()
	fmt.Println("num2: ", num2)
	// 将value转换为interface{}类型的值
	iv := nv.Interface()
	// 将interface{}类型的值转换为int类型的值
	num3 := iv.(int) + 1
	fmt.Println("num3: ", num3)
}

// 对结构体的反射
type Stu struct {
	Name  string `json:"name"`
	Age   int
	Score float64
	Sex   string
}

func (s Stu) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}
func (s Stu) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Stu) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func reflectTest02() {
	stu := Stu{Name: "Alice", Age: 30}

	types := reflect.TypeOf(stu)
	fmt.Println("类型: ", types) // main.Stu

	kinds := types.Kind()
	fmt.Println("类别: ", kinds) // struct

	values := reflect.ValueOf(stu)
	fmt.Println("值: ", values) // {Alice 30}

	// 运行的时候才能知道是Student类型，因此需要通过断言才能获取到Student类型的具体变量值
	stu2, err := values.Interface().(Stu)
	if !err {
		fmt.Println("类型断言失败: ", err)
		return
	}
	fmt.Printf("value=%v,type=%T\n", stu2, stu2)
	fmt.Println("stu2.Name: ", stu2.Name)

}

// 通过反射修改基本数据类型和修改结构体字段的值
func reflectTest03() {
	num := 100
	// 通过反射获取num的值信息，必须传入num的地址，否则无法修改num的值
	rVal := reflect.ValueOf(&num)
	fmt.Printf("rVal: %v,rVal kind: %v\n", rVal, rVal.Kind()) // rVal:  0xc0000160b8,rVal kind: ptr kind是ptr，说明rVal是一个指针类型的值
	// 通过反射修改变量的值，必须传入变量的地址，并且通过Elem()方法获取到变量的值信息
	//rVal.SetInt(200) // SetInt()方法只能修改int类型的变量的值，否则会报错，此时会报错
	rVal.Elem().SetInt(200) // 通过Elem()方法获取到num的值信息，然后调用SetInt()方法修改num的值
	fmt.Println("num: ", num)

	// 修改结构体的变量的值(方式一)
	stu := Stu{Name: "Alice", Age: 30, Score: 90}
	rVal = reflect.ValueOf(&stu)
	fmt.Printf("rVal: %v,rVal kind: %v\n", rVal, rVal.Kind()) // rVal:  0xc0000160b8,rVal kind: ptr kind是ptr，说明rVal是一个指针类型的值
	// 通过反射修改结构体类型的值，必须传入结构体变量的地址，并且通过Elem()方法获取到结构体变量的值信息
	// 通过FieldByName()方法获取到结构体变量的字段信息，然后调用SetXXX()方法修改字段的值
	rVal.Elem().FieldByName("Name").SetString("Bob") // 通过FieldByName()方法获取到结构体变量的字段信息，然后调用SetString()方法修改字段的值
	rVal.Elem().FieldByName("Age").SetInt(25)        // 通过FieldByName()方法获取到结构体变量的字段信息，然后调用SetInt()方法修改字段的值
	fmt.Printf("stu: %v\n", stu)
}

// 通过反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
func reflectTest04(stu interface{}) {
	tStu := reflect.TypeOf(stu)
	vStu := reflect.ValueOf(stu)
	kStu := tStu.Kind()
	if kStu != reflect.Struct {
		fmt.Println("不是结构体类型")
		return
	}
	// 获取结构体的字段数量
	num := vStu.NumField() // NumField()方法返回结构体类型的字段数量
	fmt.Printf("结构体的字段数量: %v\n", num)
	// 通过反射获取结构体的字段信息，遍历结构体的字段和标签
	for i := 0; i < num; i++ {
		field := tStu.Field(i)          // Field()方法返回结构体类型的第i个字段的信息
		value := vStu.Field(i)          // Field()方法返回结构体变量的第i个字段的值信息
		tagVal := field.Tag.Get("json") // 获取结构体字段的标签值，Get()方法传入标签的名称，返回标签的值
		// 如果标签值不为空，则输出字段名称、字段类型、字段值和字段标签，否则只输出字段名称、字段类型和字段值
		if tagVal != "" {
			fmt.Printf("字段名称: %v, 字段类型: %v, 字段值: %v, 字段标签: %v\n", field.Name, field.Type, value, tagVal)
		} else {
			fmt.Printf("字段名称: %v, 字段类型: %v, 字段值: %v\n", field.Name, field.Type, value)
		}
	}
	// Method默认按方法名排序对应i值，i从0开始，调用Method(i)方法获取到结构体变量的方法信息，然后调用Call()方法调用方法
	// 可以通过Value来调用方法，也可以通过Type来调用方法
	//
	// 方法是绑定在结构体变量上的，而不是绑定在结构体类型上的，所以只能通过Type来获取方法的信息，通过Value来调用方法时，无法获取到方法的名称和方法的类型信息，因为方法是绑定在结构体变量上的，而不是绑定在结构体类型上的，所以只能通过Type来获取方法的信息
	// 通过Type调用方法时，必须传入结构体变量的值信息vStu，否则会报错，因为方法是绑定在结构体变量上的，而不是绑定在结构体类型上的

	// 通过Type调用方法
	mnum1 := tStu.NumMethod() // NumMethod()方法返回结构体类型的方法数量
	fmt.Printf("通过Type调用方法\n")
	fmt.Printf("结构体的方法数量: %v\n", mnum1)
	for i := 0; i < mnum1; i++ {
		method := tStu.Method(i) // Method()方法返回结构体类型的第i个方法的信息
		fmt.Printf("方法名称: %v, 方法类型: %v\n", method.Name, method.Type)
		// 调用方法，Call()方法传入一个reflect.Value类型的切片，表示方法的参数，返回一个reflect.Value类型的切片，表示方法的返回值
		var result []reflect.Value
		// Type.Type.NumIn()方法返回方法的参数数量，如果方法没有参数，则直接调用方法，如果方法有参数，则调用方法并传入参数
		if method.Type.NumIn() == 1 { // 如果方法没有参数，则直接调用方法
			// Type.Func.Call()方法调用方法，传入一个reflect.Value类型的切片，表示方法的参数，返回一个reflect.Value类型的切片，表示方法的返回值
			// 通过Type调用方法时，必须传入结构体变量的值信息vStu，否则会报错，因为方法是绑定在结构体变量上的，而不是绑定在结构体类型上的
			result = method.Func.Call([]reflect.Value{vStu})
		} else if method.Type.NumIn() == 3 { // 如果方法有两个参数，则调用方法并传入参数
			result = method.Func.Call([]reflect.Value{vStu, reflect.ValueOf(10), reflect.ValueOf(20)})
		}
		if result != nil {
			fmt.Printf("方法的返回值: %v\n", result[0]) // 通过索引获取方法的返回值，result[0]表示方法的第一个返回值
		}
	}

	// 通过Value调用方法
	// Value调用方法，无法获取到方法的名称和方法的类型信息，因为方法是绑定在结构体变量上的，而不是绑定在结构体类型上的，所以只能通过Type来获取方法的信息
	mnum2 := vStu.NumMethod() // NumMethod()方法返回结构体类型的方法数量
	fmt.Printf("通过Value调用方法\n")
	fmt.Printf("结构体的方法数量: %v\n", mnum2)
	// Type.Method()方法返回结构体类型的第i个方法的信息，Method.Func.Call()方法调用方法，传入一个reflect.Value类型的切片，表示方法的参数，返回一个reflect.Value类型的切片，表示方法的返回值
	method := vStu.Method(0) // Method()方法返回结构体类型的第i个方法的信息
	// Type.Call()方法传入一个reflect.Value类型的切片，表示方法的参数，返回一个reflect.Value类型的切片，表示方法的返回值
	var result []reflect.Value
	result = method.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)})
	if result != nil {
		fmt.Printf("方法的返回值: %v\n", result[0]) // 通过索引获取方法的返回值，result[0]表示方法的第一个返回值
	}
}

// 通过反射修改结构体的变量的值（方式二）
// 参数为指针类型，则方法需要先通过Elem()方法获取到结构体变量的值信息，然后调用方法，否则会报错
func reflectTest05(stu interface{}) {
	fmt.Printf("stu: %v\n", stu)
	// 修改结构体的变量的值
	rVal := reflect.ValueOf(stu)
	// 通过Field()方法获取到结构体变量的第0个字段的值信息，然后调用SetString()方法修改字段的值
	rVal.Elem().Field(0).SetString("Bob")
	rVal.Elem().Field(1).SetInt(25)
	fmt.Printf("stu: %v\n", stu)
}

// 使用反射创建结构体
func reflectTest06() {
	// 通过反射创建结构体变量的值信息，必须传入结构体类型的指针，否则无法创建结构体变量的值信息
	var (
		model *Stu
		st    reflect.Type
		sv    reflect.Value
	)
	st = reflect.TypeOf(model)    // 获取model的类型信息，model是一个指向Stu类型的指针，所以st的类型信息是*main.Stu
	st = st.Elem()                // Elem()方法返回指针类型的元素类型的信息，st是一个指向Stu类型的指针，所以st.Elem()返回Stu类型的信息
	sv = reflect.New(st)          // New()方法创建一个新的指向st类型的指针类型的值信息，sv是一个指向Stu类型的指针，所以sv的类型信息是*main.Stu
	model = sv.Interface().(*Stu) // Interface()方法将sv转换为interface{}类型的值，然后通过断言将其转换为*Stu类型的值，赋值给model变量。此时赋给model的是一个已经创建好的指向Stu类型的指针类型的值，所以model和sv指向同一个Stu类型的变量
	sv = sv.Elem()
	sv.FieldByName("Name").SetString("dd") // 通过FieldByName()方法获取到结构体变量的字段信息，然后调用SetString()方法修改字段的值
	sv.FieldByName("Age").SetInt(252)
	// 此时sv是一个指向Stu类型的Value类型的值，model是一个指向Stu类型的指针类型的值，所以model和sv指向同一个Stu类型的变量，所以修改sv的字段值也会修改model指向的变量的字段值
	fmt.Println(model)
}
func main() {
	stu := Stu{Name: "Alice", Age: 30, Score: 90}
	reflectTest04(stu)

	// 可以传入结构体指针来修改其字段的值
	reflectTest05(&stu)

	reflectTest06()
}
func yongli1() {
	var v float64 = 1.2
	vValue := reflect.ValueOf(v)
	vType := reflect.TypeOf(v)
	vKind := vType.Kind()
	iValue := vValue.Interface()
	bValue := iValue.(float64)

	fmt.Println("kind:", vKind, "value:", bValue)
}
