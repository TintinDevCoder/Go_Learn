package main

import "fmt"

// 泛型栈实现
type Stack[T any] struct {
	elements []T
}

// 入栈
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// 出栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// 使用示例
func main() {
	// 整数栈
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop()) // 输出: 3 true

	// 字符串栈
	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop()) // 输出: world true
}
