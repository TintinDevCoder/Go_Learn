package main

import "fmt"

func main() {
	bscore := [3]int{}
	result := 0
	for i := range 3 {
		fmt.Printf("请输入第%d个班级的成绩\n", i+1)
		for j := range 5 {
			var n int
			fmt.Printf("第%d个学生的成绩：", j+1)
			fmt.Scanf("%d", n)
			bscore[i] += n
		}
		result += bscore[i]
	}
	for i := range 3 {
		fmt.Printf("第%d个班级的成绩为：%d\n", i+1, bscore[i])
	}
	fmt.Printf("三个班级的总成绩为：%d\n", result)
}
