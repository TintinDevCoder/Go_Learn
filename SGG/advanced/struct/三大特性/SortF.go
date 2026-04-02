package main

import (
	"fmt"
	"sort"
)

// Sort方法接收一个接口类型，这个接口类型定义了Len、Less和Swap三个方法，这三个方法是sort包中定义的接口方法
// 任何类型只要实现了这三个方法，就可以被视为实现了sort.Interface接口，从而可以使用sort包中的Sort函数对该类型的切片进行排序
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
type mySortStruct struct {
	data []int
}

func (s *mySortStruct) Len() int {
	return len(s.data)
}
func (s *mySortStruct) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}
func (s *mySortStruct) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func main() {
	// 通过实现Sortable接口，可以使用sort包中的Sort函数对mySortStruct类型的切片进行排序
	s := &mySortStruct{data: []int{5, 2, 6, 3, 1, 4}}
	fmt.Println(s.data)
	sort.Sort(s) // 需要导入sort包
	fmt.Println(s.data)
}
