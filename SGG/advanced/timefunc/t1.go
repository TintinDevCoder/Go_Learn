package main

import (
	"fmt"
	"strconv"
	"time"
)

// 设置某时间
func timeset() {
	// 解析时间字符串
	layout := "2006-01-02 15:04:05" // Go 特有的格式化模板
	str := "2026-03-06 16:00:00"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("解析时间失败:", err)
		return
	}
	// 使用 time.Date 设置特定时间
	// 参数依次为：年, 月, 日, 时, 分, 秒, 纳秒, 时区
	specificTime := time.Date(2026, time.March, 6, 15, 30, 0, 0, time.Local)
	fmt.Println("解析后的时间:", parsedTime)
	fmt.Println("设置的时间:", specificTime)
}

// 修改时间变量
func updatetime() {
	now := time.Now()
	// 1. 修改小时/分钟 (使用 Add)
	afterTwoHours := now.Add(2 * time.Hour)

	// 2. 修改年/月/日 (使用 AddDate)
	// 参数：年, 月, 日 (正数为加，负数为减)
	nextYear := now.AddDate(1, 0, 0)

	fmt.Println("当前时间:", now)
	fmt.Println("两小时后:", afterTwoHours)
	fmt.Println("明年此时:", nextYear)
}

// 时间和日期相关的函数
func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间", now)

	// 获取其他日期的信息
	fmt.Println("年份：", now.Year())
	fmt.Println("月份：", now.Month())
	fmt.Println("日份：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	// 格式化日期时间
	// 第一种方式
	fmt.Println("第一种方式")
	// 2026-3-5 16:29:53
	fmt.Printf("%01d-%01d-%01d %01d:%01d:%01d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 2026-03-05 16:29:53
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 2026-003-005 016:030:040
	fmt.Printf("%03d-%03d-%03d %03d:%03d:%03d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// d前面的数字是指这个字段最少的占位数
	fmt.Println()

	// 第二种方式
	// 2006/01/02 15:04:05 其中的数字是固定的，必须这样写
	fmt.Println("第二种方式")
	str1 := now.Format("2006/01/02 15:04:05")
	fmt.Println(str1)
	str2 := now.Format("2006-01-02")
	fmt.Println(str2)
	str3 := now.Format("15:04:05")
	fmt.Println(str3)

	// 时间的常量
	i := 0
	for {
		i++
		fmt.Println("i:", i)
		// time.Sleep(time.Second)            // 休眠1秒
		time.Sleep(time.Millisecond * 100) // 休眠0.1秒
		if i == 10 {
			break
		}
		// time包中定义了很多时间相关的常量，time.Second表示1秒，time.Minute表示1分钟，time.Hour表示1小时
	}
	// func Sleep(d Duration)
	// 让当前的goroutine休眠d的时间，d的类型是Duration，Duration是一个int64类型的别名，表示纳秒数，
	// time包中定义了很多时间相关的常量，time.Second表示1秒，time.Minute表示1分钟，time.Hour表示1小时

	// 时间戳 —— 可以获取随机的数字
	// Unix时间戳（秒） UnixNano（纳秒）
	fmt.Println("unix时间戳：", now.Unix())
	fmt.Println("unixnano时间戳：", now.UnixNano())

	timeset()
	updatetime()
	//lianxi1()
	lianxi2()
}
func lianxi1() {
	time1 := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	time2 := time.Now().Unix()
	fmt.Println("耗时：", time2-time1, "秒")
}
func lianxi2() {
	fmt.Println("请输入年：")
	year := 0
	fmt.Scanln(&year)
	fmt.Println("请输入月：")
	month := 0
	fmt.Scanln(&month)
	if month > 12 || month < 1 {
		fmt.Println("输入的月份不合法")
		return
	}
	layout := "2006-1" // Go 特有的格式化模板
	if month == 12 {
		month = 1
	} else {
		month++
	}
	str := strconv.Itoa(year) + "-" + strconv.Itoa(month)
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("解析时间失败:", err)
	}
	// 减去一天
	parsedTime = parsedTime.AddDate(0, 0, -1)
	fmt.Println("输入的月份有：", parsedTime.Day(), "天")
}
