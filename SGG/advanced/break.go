package main

import "fmt"

// break语句用于终止当前循环，跳出循环体，继续执行循环后面的代码
// 也可以跳转到标签处执行，标签必须在同一个函数内，并且在break语句之前定义
// 也就是会跳转出标签处的循环体，继续执行标签处后面的代码
func main() {
	for i := range 5 {
		if i == 3 {
			break
		}
	}

label1:
	for i := range 4 {
		for j := range 10 {
			if j == 2 {
				break label1 // 跳出标签处的循环体，继续执行标签处后面的代码
			}
			fmt.Printf("i: %d j: %d\n", i, j)
		}
	}
	fmt.Println("循环结束")
	breaklianxi1()
}
func breaklianxi1() {
	name := "张无忌"
	passw := 888
	is := false
	for i := range 3 {
		gassname := ""
		gasspassw := 0
		println("请依次输入用户名和密码：")
		fmt.Scanf("%s", &gassname)
		fmt.Scanf("%d", &gasspassw)
		if name == gassname && passw == gasspassw {
			println("登录成功")
			is = true
			break
		} else {
			println("登录失败，还有", 2-i, "次机会")
		}
	}
	if is {
		println("登陆成功")
	} else {
		println("登陆失败")
	}
}
