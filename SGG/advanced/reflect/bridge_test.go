package main

import (
	"reflect"
	"testing"
)

func TestReflectFunc(t *testing.T) {
	// 定义两个参数不同的待调用函数
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	// 定义桥接函数
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		// 初始化反射 Value 切片，长度为参数个数
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			// 将每一个传入的 interface{} 参数转为 reflect.Value
			inValue[i] = reflect.ValueOf(args[i])
		}

		// 获取目标函数的反射对象
		function = reflect.ValueOf(call)
		// 通过反射进行动态调用
		function.Call(inValue)
	}

	// 动态调用测试
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}
