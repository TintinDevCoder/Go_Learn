package mathutils

import "fmt"

// 一般情况下，要保持包名和目录名相同
// 对于变量名、函数名、常量名等，首字母大写表示可导出（public），首字母小写表示不可导出（private）
// public状态下同模块内的其他包也可以访问，private状态下只能在当前包内访问
var Say int = 10

// 注意：函数名首字母必须大写才能被 main 包调用
func SayHello(name string) {
	fmt.Printf("Hello, %s from packageRules!\n", name)
}
