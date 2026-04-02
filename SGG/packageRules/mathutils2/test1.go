package mathutils

// 不同文件之间，即使包名相同，但因为路径不同，因此是不同的包，无法直接访问，需要使用 import 语句引入包
import (
	"SGG/packageRules/mathutils"
	"fmt"
)

func main() {
	// 可以访问 mathutils 包中导出的变量和函数
	fmt.Println(mathutils.Say)

	mathutils.SayHello("Gemini")
}
