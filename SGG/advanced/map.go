package main

import (
	"fmt"
	"sort"
)

// map 是key、value的数据结构，又称为字段或者关联数组。类似其他编程语言的集合
// map是引用类型的，函数引入修改后会影响原本的值
// map的长度会动态增加，能动态增长键值对
// map的value常为struct类型的
func main() {
	// map的声明
	var m1 map[string]int // map的零值是nil，nil map没有底层数据结构，不能存储键值对
	// 声明map是不会分配内存的，初始化需要make，分配内存后才可以赋值和使用
	m2 := make(map[string]int) // make函数创建map，make函数的语法是：make(map[K]V)，K表示map的键类型，V表示map的值类型，make函数会返回一个非nil的map，可以存储键值对
	m3 := map[string]int{}     // map的字面量创建，字面量的语法是：map[K]V{key1: value1, key2: value2}，K表示map的键类型，V表示map的值类型，key1、key2表示map的键，value1、value2表示map的值，字面量创建的map也是非nil的，可以存储键值对
	m4 := map[string]int{"a": 1, "b": 2, "c": 3}
	println(m1 == nil) // true
	println(m2 == nil) // false
	println(m3 == nil) // false
	println(m4 == nil) // false

	//m1["a"] = 1        // nil map不能存储键值对，会发生运行时错误
	m1 = make(map[string]int, 10)
	m1["a"] = 1
	fmt.Println(m1["a"])
	// map是一个无序的集合，map的键值对在内存中的存储顺序是不确定的，每次迭代map时，键值对的顺序可能会不同

	// 复杂使用
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "张三"
	stuMap["stu01"]["age"] = "18"
	fmt.Println(stuMap["stu01"])

	// map的crud
	// 创建和更新
	m := make(map[string]int)
	m["a"] = 1 // 创建键值对，如果键不存在，则创建键对应的值
	m["a"] = 2 // 创建键值对，如果键已经存在，则更新键对应的值
	fmt.Println(m)
	// 删除
	delete(m, "a") // delete函数删除键值对，delete函数的语法是：delete(map, key)，map表示要删除键值对的map，key表示要删除的键
	// 如果键存在，则删除键值对，如果键不存在，则什么都不做
	fmt.Println(m)
	// 全部删除
	m = make(map[string]int)

	// 查找
	m["a"] = 1
	value, ok := m["a"] // map的查找，map的语法是：value, ok := map[key]，map表示要查找的map，key表示要查找的键，value表示键对应的值，ok表示键是否存在，如果键存在，则ok为true，否则为false
	if ok {
		fmt.Println("键存在，值为：", value)
	} else {
		fmt.Println("键不存在")
	}

	// 遍历
	// map不能用for循环遍历，但可以用for range遍历
	m5 := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m5 {
		fmt.Println("key:", k, "value:", v)
	}
	for k1, v1 := range stuMap {
		fmt.Println("k1:", k1, " ")
		for k2, v2 := range v1 {
			fmt.Println(" ", "k2:", k2, "v2:", v2)
		}
	}

	// map的切片
	monsters := make([]map[string]string, 1)
	monsters[0] = make(map[string]string)
	monsters[0]["name"] = "牛魔王"
	monsters[0]["age"] = "500"
	fmt.Println(monsters) // [map[age:500 name:牛魔王] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
	monsters = append(monsters, make(map[string]string))
	monsters[1]["name"] = "铁扇公主"
	monsters[1]["age"] = "400"
	fmt.Println(monsters) // [map[age:500 name:牛魔王] map[age:400 name:铁扇公主] map[] map[] map[] map[] map[] map[] map[] map[]]

	fmt.Println()
	// map的排序
	// golang中的map默认是无序的，map的键值对在内存中的存储顺序是不确定的
	// 每次迭代map时，键值对的顺序可能会不同
	// 如果需要对map进行排序，可以将map的键存储在一个切片中，对切片进行排序，然后按照排序后的切片顺序访问map中的键值对
	// 先将map的key放入切片中
	// 对切片怕排序
	// 遍历切片，按照key，输出map
	m6 := make(map[int]string)
	m6[0] = "s1"
	m6[1] = "s2"
	m6[2] = "s3"
	keys := make([]int, 0)
	for k, _ := range m6 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println("key:", v, "value:", m6[v])
	}
}
