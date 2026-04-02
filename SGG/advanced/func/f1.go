package main

import "fmt"

// 函数可以返回多个值
// 支持对函数返回值命名，命名的返回值在函数体内就可以当做变量使用了
func getSumAndSub(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

// 函数参数可以是指针类型，可以修改外部变量的值
func test1(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("neibu n1:", *n1)
}

// 函数本身也是一种数据类型,可以赋给一个变量，通过变量对函数调用
func test2() {
	a1 := getSumAndSub
	fmt.Printf("a1的类型:%T\n", a1)
	a1(10, 20)
}

// 函数作为一个数据类型，可以作为形参，并且调用
func test3(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func getSum(a int, b int) int {
	return a + b
}

// 可以自定义数据类型
type myFuncType func(int, int) int

func test4() {
	type myint int // 给int取别名
	// 在go中 int和myint都是int数据类型，但go认为是两个不同的数据类型
	var num1 myint
	fmt.Printf("num1的类型：%T, 值：%v\n", num1, num1)

	var funvar myFuncType
	funvar = getSum
	fmt.Printf("funvar的类型:%T\n", funvar)
	fmt.Println(funvar(10, 20))
}

// go支持可变参数,参数是一个切片，在函数体内可以当做切片使用
// 可变参数必须放在参数列表的最后
func test5(args ...int) (sum int) {
	for i := range args {
		sum += args[i]
	}
	return sum
}
func main() {
	test2()

	test3(getSum, 1, 2)

	test4()

	a := test5(1, 2, 3, 4, 5)
	fmt.Println(a)
}
