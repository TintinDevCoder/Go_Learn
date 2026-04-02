package main

import (
	"encoding/json"
	"fmt"
)

// JSON是一种轻量级的数据交换格式，它基于JavaScript对象表示法（JSON）
// JSON易于用于机器解析和生成，并有效的提高网络传输效率。网络传输时先将数据序列化为JSON，再进行传输。在接收的时候，将JSON反序列化为对象。
// json语言中，一切皆为对象，所有数据都是对象，对象有属性和值。
// json是用键值对存储数据的，键是字符串，值可以是字符串、数字、布尔值、数组、对象等。
// 键值对组合中的键名写在前面，用冒号：隔开，键值对之间用逗号,隔开

// 可以使用tag对结构体的字段进行标记，以便在序列化时指定字段的序列化方式。
// 但只能用在首字母大写的字段上，因为小写的字段不能被json包引用，是私有的 。
type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday float64 `json:"birthday"`
	Sal      float64 `json:"sal"`
	skill    string  `json:"skill"`
}

// json序列化
func jsonSerialize() (s1 string, s2 string, s3 string, s4 string) {
	// 结构体序列化
	monster := Monster{
		Name:     "张三",
		Age:      18,
		Birthday: 1.2345,
		Sal:      1234.5678,
	}
	// 将monster序列化为json字符串
	jsonStr, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("json序列化错误：", err)
	} else {
		s1 = string(jsonStr)
		fmt.Println("结构体的json序列化结果：", s1)
	}

	//map序列化
	m := make(map[string]interface{})
	m["name"] = "张三"
	m["age"] = 18
	m["birthday"] = 1.2345
	m["sal"] = 1234.5678
	jsonStr, err = json.Marshal(m)
	if err != nil {
		fmt.Println("json序列化错误：", err)
	} else {
		s2 = string(jsonStr)
		fmt.Println("map的json序列化结果：", s2)
	}

	// 切片序列化
	slice := make([]map[string]interface{}, 0)
	m1 := make(map[string]interface{})
	m1["name"] = "张三"
	m1["age"] = 18
	m1["birthday"] = 1.2345
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "张e"
	m2["age"] = 11
	m2["birthday"] = 1
	slice = append(slice, m2)
	jsonStr, err = json.Marshal(slice)
	if err != nil {
		fmt.Println("json序列化错误：", err)
	} else {
		s3 = string(jsonStr)
		fmt.Println("切片的json序列化结果：", s3)
	}

	// 普通数据类型序列化
	var a int = 10
	jsonStr, err = json.Marshal(a)
	if err != nil {
		fmt.Println("json序列化错误：", err)
	} else {
		s4 = string(jsonStr)
		fmt.Println("普通数据的json序列化结果：", s4)
	}
	return
}

// json反序列化
func jsonDeserialize(json1 string, json2 string, json3 string, json4 string) {
	// 结构体反序列化
	var monster Monster
	err := json.Unmarshal([]byte(json1), &monster)
	if err != nil {
		fmt.Println("json反序列化错误：", err)
	} else {
		fmt.Println("结构体的json反序列化结果：", monster)
	}

	// map反序列化
	var m map[string]interface{}
	// 反序列化底层会自动make，所以不需要make
	err = json.Unmarshal([]byte(json2), &m)
	if err != nil {
		fmt.Println("json反序列化错误：", err)
	} else {
		fmt.Println("map的json反序列化结果：", m)
	}

	// 切片反序列化
	var slice []map[string]interface{}
	err = json.Unmarshal([]byte(json3), &slice)
	if err != nil {
		fmt.Println("json反序列化错误：", err)
	} else {
		fmt.Println("切片的json反序列化结果：", slice)
	}

	// 普通数据类型反序列化
	var a int
	err = json.Unmarshal([]byte(json4), &a)
	if err != nil {
		fmt.Println("json反序列化错误：", err)
	} else {
		fmt.Println("普通数据的json反序列化结果：", a)
	}
}

func main() {
	s1, s2, s3, s4 := jsonSerialize()
	jsonDeserialize(s1, s2, s3, s4)
}
