package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func lianxi1() {
	val1 := rand.IntN(101) + 1
	i := 1
	is := false
	for ; i <= 10; i++ {
		var val2 int
		println("请输入一个1-100的整数：")
		_, err := fmt.Scanln(&val2)
		if err != nil {
			fmt.Println("输入错误，请输入一个整数")
			return
		}
		if val2 > 100 || val2 < 1 {
			fmt.Println("输入的数不合法，请输入一个1-100的整数")
			return
		}
		if val2 == val1 {
			is = true
			break
		} else if val2 > val1 {
			fmt.Println("你输入的数太大了")
		} else {
			fmt.Println("你输入的数太小了")
		}
	}
	if is {
		fmt.Printf("恭喜你猜对了，你一共猜了%d次\n", i)
		fmt.Print("评分：")
		if i == 1 {
			fmt.Println("你是个天才！")
		} else if i >= 2 && i <= 3 {
			fmt.Println("你很聪明！")
		} else if i == 10 {
			fmt.Println("可算猜对了")
		} else {
			fmt.Println("一般般")
		}
	} else {
		fmt.Printf("很遗憾，你没有猜对，正确的数是%d\n", val1)
	}
}

// checkStatus 判断在指定日期是在打鱼还是晒网
func checkStatus(year, month, day int) string {
	// 1. 设置起始日期：1990年1月1日
	startDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local)

	// 2. 设置目标日期
	targetDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	// 3. 计算天数差
	// Sub 返回 Duration，通过 Hours() 除以 24 转换为天数
	// 注意：如果是同一天，Hours 为 0，即第 1 天
	duration := targetDate.Sub(startDate)
	totalDays := int(duration.Hours()/24) + 1

	if totalDays < 1 {
		return "日期早于起始时间"
	}

	// 4. 对 5 取模
	remainder := totalDays % 5

	// 5. 判断状态
	// 余数为 1, 2, 3 表示打鱼；余数为 4, 0 表示晒网
	if remainder >= 1 && remainder <= 3 {
		return "打鱼 (Fishing)"
	}
	return "晒网 (Drying)"
}
func lianxi2() {
	// 1. 获取用户输入的日期
	fmt.Println("请输入一个日期 (格式: 年 月 日，例如: 2024 6 15)：")
	year := 0
	month := 0
	day := 0
	fmt.Scanf("%d %d %d", &year, &month, &day)

	// 2. 调用函数判断状态
	status := checkStatus(year, month, day)

	// 3. 输出结果
	fmt.Printf("在 %d-%02d-%02d 是: %s\n", year, month, day, status)
}
func printPrimes(limit int) {
	isPrime := make([]bool, limit+1)
	for p := 2; p*p <= limit; p++ {
		if !isPrime[p] {
			for pp := p * p; pp <= limit; pp += p {
				isPrime[pp] = true
			}
		}
	}
	fmt.Printf("%d 以内的素数有：\n", limit)
	k := 1
	for i := 1; i <= limit; i++ {
		// 一行打印5个
		if !isPrime[i] {
			if k == 5 {
				fmt.Println()
				fmt.Print(i, " ")
				k = 1
			} else {
				fmt.Print(i, " ")
				k++
			}
		}
	}
	fmt.Println()
}
func main() {
	printPrimes(10000)
}
