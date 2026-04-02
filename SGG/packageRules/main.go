package main

// 只有在main包下的main函数才能作为程序的入口点，其他包中的函数不能直接运行
// 同一个包下的文件可以直接访问包内的变量和函数，不需要导入包
// 使用 import 语句引入 mathutils 包
import (
	mathutil "SGG/packageRules/mathutils" // 可以取别名
	"fmt"
)

// 程序运行必须在 main 包中，并且 main 包必须有一个 main 函数作为程序的入口点
func main() {
	// 可以访问 mathutils 包中导出的变量和函数
	fmt.Println(mathutil.Say)

	mathutil.SayHello("Gemini")
}

// 一般来说，编译推荐按照包进行编译，而不是按照文件进行编译，这样可以避免一些编译错误和依赖问题
// 例如，使用 go build 命令编译整个包，而不是单个文件
// go build SGG/packageRules/mathutils
// 这样会编译整个 mathutils 包，并且生成一个可执行文件，或者使用 go install 命令安装包到 GOPATH 中，不会漏掉一些文件或者依赖
// 如果一个包内有多个main函数，那么编译时会报错，因为一个程序只能有一个入口点，必须保证每个包内只有一个 main 函数
/*
# SGG/packageRules
packageRules\main2.go:5:6: main redeclared in this block
        packageRules\main.go:12:6: other declaration of main
*/

// go install SGG/packageRules/mathutils
// 这样会将 mathutils 包安装到 GOPATH 中，可以在其他包中直接 import mathutils 来使用包中的函数和变量
