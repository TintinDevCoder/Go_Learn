package main

import (
	"errors"
	"fmt"
	"sync"
)

// 1. 定义一个自定义错误结构体 (体现结构体和字段)
type NumberError struct {
	Number  int
	Message string
}

// 2. 隐式实现 error 接口 (只要有 Error() 方法就是实现了接口)
func (e *NumberError) Error() string {
	return fmt.Sprintf("错误: 数字 %d -> %s", e.Number, e.Message)
}

// 3. 结构体内嵌 (体现组合优于继承)
type Stats struct {
	TotalProcessed int
}

type Processor struct {
	Stats      // 内嵌结构体，Processor 直接拥有 TotalProcessed 字段
	Multiplier int
}

// 4. 定义方法 (体现 Method Receivers)
// test_number 函数：演示核心特性
func (p *Processor) test_number(input int) (int, error) {
	// 5. Defer 特性：无论函数如何结束，都会执行（常用于释放资源）
	defer fmt.Println("--- 本次数字处理逻辑结束 ---")

	fmt.Printf("开始处理数字: %d\n", input)

	// 6. 显式错误处理与多返回值
	if input < 0 {
		return 0, &NumberError{Number: input, Message: "不能处理负数"}
	}

	if input == 0 {
		return 0, errors.New("输入不能为零")
	}

	// 模拟计算
	result := input * p.Multiplier
	p.TotalProcessed++

	return result, nil
}

func main() {
	// 初始化结构体
	proc := &Processor{
		Multiplier: 10,
	}

	// 7. Channel：用于在 Goroutine 之间通信
	resultsChan := make(chan int)
	var wg sync.WaitGroup // 用于等待所有并发任务完成

	numbers := []int{5, -2, 10, 0, 8}

	fmt.Println(">>> 启动并发处理...")

	for _, n := range numbers {
		wg.Add(1)

		// 8. Goroutine：并发执行任务
		go func(num int) {
			defer wg.Done()

			// 调用核心函数
			res, err := proc.test_number(num)

			if err != nil {
				// 错误处理逻辑
				fmt.Printf("[异常报告] %v\n", err)
				return
			}

			// 将结果通过 Channel 发送出去
			resultsChan <- res
		}(n)
	}

	// 另外启动一个 Goroutine 负责关闭 Channel
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// 9. 循环从 Channel 接收数据
	for finalRes := range resultsChan {
		fmt.Printf("[处理成功] 结果为: %d\n", finalRes)
	}

	fmt.Printf("\n所有任务完成，总计成功处理次数: %d\n", proc.TotalProcessed)
}
