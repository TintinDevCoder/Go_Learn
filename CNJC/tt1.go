package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// 情况 A：留在栈上
// 虽然使用了 &User，但这个变量没有离开函数的作用域
func stayOnStack() {
	u := &User{ID: 1, Name: "Local"}
	fmt.Println(u.ID)
}

// 情况 B：逃逸到堆上
// 函数返回了局部变量的指针，外部还要用，必须放堆里
func escapeToHeap() *User {
	u := &User{ID: 2, Name: "Escaped"}
	return u
}

func main() {
	stayOnStack()
	u := escapeToHeap()
	fmt.Println(u.Name)
}
