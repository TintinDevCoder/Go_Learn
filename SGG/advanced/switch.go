package main

import "fmt"

/**
 * switch 语句
 */
// switch 语句
func t6_test1() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

// Type Switch
func t6_test2() {
	var x interface{}
	x = 7
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

// fallthrough 语句 fallthrough 会强制执行后面的一条 case 语句
func t6_test3() {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

func main() {
	t6_test2()
}
func lianxi1(c byte) byte {
	switch c {
	case 'a':
		return 'A'
	case 'b':
		return 'B'
	case 'c':
		return 'C'
	case 'd':
		return 'D'
	default:
		return c - 'a' + 'A'
	}
}
func lianxi2(score int) {
	switch {
	case score <= 100 && score >= 60:
		fmt.Println("合格")
	case score < 60 && score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误")
	}
}
func lianxi3(month int) {
	switch month {
	case 3, 4, 5:
		println("春季")
	case 6, 7, 8:
		println("夏季")
	case 9, 10, 11:
		println("秋季")
	case 12, 1, 2:
		println("冬季")
	default:
		fmt.Println("输入有误")
	}
}
