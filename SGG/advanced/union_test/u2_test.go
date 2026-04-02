package main

import "testing"

func TestSub(t *testing.T) {
	res := sub(1, 2)
	if res != -1 {
		// Fatalf函数会停止测试，并打印错误信息
		// 该方法的第一个参数必须是常量字符串，不能是变量，而后面的参数可以是变量，填充前面的参数
		t.Fatalf("测试失败，结果为：%d", res)
	}
	// 正确，输出日志
	t.Logf("测试成功，结果为：%d", res)
}
