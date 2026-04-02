package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup是Go语言中的一个同步原语，用于等待一组协程完成。它提供了Add、Done和Wait方法来管理协程的数量和等待状态。
func fetchData(site string, wg *sync.WaitGroup) {
	// 3. 任务结束时，告知计数器减 1
	// 使用 defer 确保即使函数中间报错，也能执行 Done，避免死锁
	defer wg.Done()

	fmt.Printf("正在从 %s 爬取论文数据...\n", site)

	// 模拟网络耗时
	time.Sleep(time.Duration(2) * time.Second)

	fmt.Printf("【完成】%s 数据下载完毕\n", site)
}

func main() {
	var wg sync.WaitGroup
	sites := []string{"arXiv", "IEEE", "ACM"}

	for _, site := range sites {
		// 1. 启动协程前，计数器加 1
		wg.Add(1)

		// 2. 开启并发任务（注意传指针 &wg）
		go fetchData(site, &wg)
	}

	fmt.Println("主程序：正在等待所有实验室任务完成...")

	// 4. 阻塞在此，直到计数器归零
	wg.Wait()

	fmt.Println("主程序：所有数据已汇总，准备开始分析。")
}
