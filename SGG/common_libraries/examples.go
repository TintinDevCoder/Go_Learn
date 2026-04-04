package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go 标准库示例 ===")

	// 1. 字符串操作示例
	stringExamples()

	// 2. 文件操作示例
	fileExamples()

	// 3. JSON 操作示例
	jsonExamples()

	// 4. HTTP 客户端示例
	httpClientExamples()

	// 5. 并发示例
	concurrencyExamples()

	// 6. 排序示例
	sortExamples()

	// 7. 时间操作示例
	timeExamples()

	fmt.Println("\n=== 所有示例完成 ===")
}

// 1. 字符串操作示例
func stringExamples() {
	fmt.Println("\n--- 字符串操作示例 ---")

	// 字符串连接
	s1 := "Hello"
	s2 := "World"
	result := s1 + " " + s2
	fmt.Printf("字符串连接: %s\n", result)

	// 使用 strings.Builder
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	fmt.Printf("Builder构建: %s\n", builder.String())

	// 字符串分割
	str := "apple,banana,orange"
	parts := strings.Split(str, ",")
	fmt.Printf("分割结果: %v\n", parts)

	// 字符串替换
	replaced := strings.ReplaceAll("hello world", "world", "Go")
	fmt.Printf("替换结果: %s\n", replaced)

	// 字符串转换
	numStr := "123"
	num, err := strconv.Atoi(numStr)
	if err == nil {
		fmt.Printf("字符串转整数: %d\n", num)
	}

	// 整数转字符串
	strFromInt := strconv.Itoa(456)
	fmt.Printf("整数转字符串: %s\n", strFromInt)
}

// 2. 文件操作示例
func fileExamples() {
	fmt.Println("\n--- 文件操作示例 ---")

	// 写入文件
	content := "Hello, Go Standard Library!\nThis is a test file."
	err := os.WriteFile("test.txt", []byte(content), 0644)
	if err != nil {
		log.Printf("写入文件失败: %v", err)
		return
	}
	fmt.Println("文件写入成功: test.txt")

	// 读取文件
	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Printf("读取文件失败: %v", err)
		return
	}
	fmt.Printf("文件内容:\n%s\n", string(data))

	// 使用bufio逐行读取
	file, err := os.Open("test.txt")
	if err != nil {
		log.Printf("打开文件失败: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("逐行读取:")
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("第%d行: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("扫描文件失败: %v", err)
	}

	// 删除文件
	err = os.Remove("test.txt")
	if err != nil {
		log.Printf("删除文件失败: %v", err)
	} else {
		fmt.Println("文件已删除: test.txt")
	}
}

// 3. JSON 操作示例
func jsonExamples() {
	fmt.Println("\n--- JSON 操作示例 ---")

	// 定义结构体
	type Person struct {
		Name    string    `json:"name"`
		Age     int       `json:"age"`
		Email   string    `json:"email,omitempty"`
		Active  bool      `json:"active"`
		Created time.Time `json:"created_at"`
	}

	// 创建实例
	person := Person{
		Name:    "张三",
		Age:     30,
		Email:   "zhangsan@example.com",
		Active:  true,
		Created: time.Now(),
	}

	// 编码为JSON
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Printf("JSON编码失败: %v", err)
		return
	}
	fmt.Printf("JSON编码结果:\n%s\n", string(jsonData))

	// 解码JSON
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		log.Printf("JSON解码失败: %v", err)
		return
	}
	fmt.Printf("JSON解码结果: 姓名=%s, 年龄=%d\n", decodedPerson.Name, decodedPerson.Age)

	// 解码到map
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("解码到map失败: %v", err)
		return
	}
	fmt.Printf("Map中的name字段: %v\n", data["name"])
}

// 4. HTTP 客户端示例
func httpClientExamples() {
	fmt.Println("\n--- HTTP 客户端示例 ---")

	// 发送GET请求
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Printf("HTTP请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应头: Content-Type=%s\n", resp.Header.Get("Content-Type"))

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应体失败: %v", err)
		return
	}

	// 只打印前200个字符
	if len(body) > 200 {
		fmt.Printf("响应体(前200字符):\n%s...\n", string(body[:200]))
	} else {
		fmt.Printf("响应体:\n%s\n", string(body))
	}

	// 创建自定义请求
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return
	}
	req.Header.Set("User-Agent", "Go-Standard-Lib-Example/1.0")
	req.Header.Set("X-Custom-Header", "CustomValue")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("发送请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("自定义请求状态码: %d\n", resp.StatusCode)
}

// 5. 并发示例
func concurrencyExamples() {
	fmt.Println("\n--- 并发示例 ---")

	// 使用WaitGroup
	var wg sync.WaitGroup
	results := make(chan string, 5)

	// 启动多个goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			results <- fmt.Sprintf("Goroutine %d 完成", id)
		}(i)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	fmt.Println("等待goroutine完成...")
	for result := range results {
		fmt.Println(result)
	}

	// 互斥锁示例
	var mu sync.Mutex
	var counter int

	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("计数器最终值: %d\n", counter)
}

// 6. 排序示例
func sortExamples() {
	fmt.Println("\n--- 排序示例 ---")

	// 整数排序
	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("排序前: %v\n", numbers)
	sort.Ints(numbers)
	fmt.Printf("排序后: %v\n", numbers)

	// 字符串排序
	fruits := []string{"banana", "apple", "orange", "grape"}
	fmt.Printf("排序前: %v\n", fruits)
	sort.Strings(fruits)
	fmt.Printf("排序后: %v\n", fruits)

	// 自定义排序
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 28},
	}
	fmt.Printf("排序前: %v\n", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("按年龄排序后: %v\n", people)

	// 查找
	index := sort.SearchInts(numbers, 5)
	if index < len(numbers) && numbers[index] == 5 {
		fmt.Printf("数字5在位置: %d\n", index)
	} else {
		fmt.Println("数字5未找到")
	}
}

// 7. 时间操作示例
func timeExamples() {
	fmt.Println("\n--- 时间操作示例 ---")

	// 当前时间
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("时间戳: %d\n", now.Unix())

	// 创建特定时间
	specificTime := time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)
	fmt.Printf("特定时间: %s\n", specificTime.Format("2006-01-02 15:04:05"))

	// 时间计算
	oneHourLater := now.Add(1 * time.Hour)
	fmt.Printf("一小时后: %s\n", oneHourLater.Format("15:04:05"))

	twoDaysAgo := now.Add(-48 * time.Hour)
	fmt.Printf("两天前: %s\n", twoDaysAgo.Format("2006-01-02"))

	// 时间差
	duration := oneHourLater.Sub(now)
	fmt.Printf("时间差: %v\n", duration)

	// 解析时间
	parsedTime, err := time.Parse("2006-01-02", "2023-12-25")
	if err != nil {
		log.Printf("解析时间失败: %v", err)
	} else {
		fmt.Printf("解析的时间: %s\n", parsedTime.Format("2006年01月02日"))
	}

	// 定时器
	fmt.Println("等待2秒...")
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("2秒时间到!")

	// 打点器
	fmt.Println("开始打点器（3次，间隔1秒）:")
	ticker := time.NewTicker(1 * time.Second)
	count := 0
	for range ticker.C {
		count++
		fmt.Printf("Tick %d\n", count)
		if count >= 3 {
			ticker.Stop()
			break
		}
	}
}
