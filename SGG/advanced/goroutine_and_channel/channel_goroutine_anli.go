package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 如下例子中，writeData用于往intChan写入数据，readData用于从intChan读取数据。
// closeChan用于通知主 goroutine 何时完成读取数据。
// 当writeData完成写入数据后，关闭intChan；当readData完成读取数据后，向closeChan发送一个信号，通知主 goroutine 结束等待。
// 主线程等待closeChan的数据，此时两个协程在工作，程序不会报错，当readData完成读取数据后，向closeChan发送一个信号，并关闭closeChan，主线程接收到信号后结束等待，程序正常结束。
func anli1() {
	intChan := make(chan int, 50)
	closeChan := make(chan int, 1)
	writeData := func() {
		for i := range 50 {
			intChan <- i
		}
		close(intChan)
	}
	readData := func() {
		for {
			i, ok := <-intChan
			if !ok {
				break
			}
			fmt.Println(i)
		}
		closeChan <- -1
		close(closeChan)
	}
	go writeData()
	go readData()
	for {
		if i, ok := <-closeChan; !ok || i == -1 {
			break
		}
	}
}

// 8个协程分别计算1-2000的数字分别的累加
func anli2() {
	var wg sync.WaitGroup // 1. 引入等待组
	numChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 1)
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
	// 初始化结果集
	resChan <- make(map[int]int)

	xc := func() {
		m := make(map[int]int)
		for {
			if n, ok := <-numChan; !ok {
				break
			} else {
				res := 0
				j := n
				for j != 0 {
					res += j
					j--
				}
				m[n] = res
			}
		}
		for {
			// 先从中拿去结果集
			if ma, ok := <-resChan; !ok {
				continue
			} else {
				// 一次性将计算的结果放入结果集
				for key, value := range m {
					ma[key] = value
				}
				// 重新放入
				resChan <- ma
				break
			}
		}
		defer wg.Done() // 函数退出说明协程完成工作，等待组-1
	}
	for i := 0; i < 8; i++ {
		wg.Add(1) // 增加计数
		go xc()
	}
	wg.Wait()      // 阻塞等待 8 个协程全部真正执行完
	close(resChan) // 此时关闭才是绝对安全的
	// 遍历结果集
	res := <-resChan
	for k := 1; k <= 2000; k++ {
		fmt.Printf("res[%d] = %d\n", k, res[k])
	}
}

// 使用一个协程生成1000个数据，写入file1文件
// 再用一个协程读取file1文件，并排序数据，放入file2
func anli3() {
	file1Chan := make(chan bool, 1)
	file2Chan := make(chan bool, 1)
	writeDataToFile := func() {
		file1, err1 := os.OpenFile("./advanced/goroutine_and_channel/file1.txt", os.O_WRONLY|os.O_CREATE, 0777)
		if err1 != nil {
			fmt.Println(err1)
			file1Chan <- false
		} else {
			writer := bufio.NewWriter(file1)
			// 生成1000个数据
			for i := 0; i < 1000; i++ {
				n := rand.IntN(10000)                     // 生成0-10000的随机数字
				writer.WriteString(strconv.Itoa(n) + " ") // 写入文件
			}
			writer.Flush()
			file1Chan <- true
		}
		// 完成生成
		close(file1Chan)
		// 关闭文件
		defer func() {
			if closeErr := file1.Close(); closeErr != nil {
				// 处理关闭文件时的错误，例如记录日志等
				println("关闭文件时发生错误：", closeErr)
			} else {
				println("文件已成功关闭")
			}
		}()
	}
	sort := func() {
		// 循环等待写完文件
		for {
			if v, ok := <-file1Chan; ok {
				if v { // 成功写完
					// 读取文件数据
					bytes, err1 := os.ReadFile("./advanced/goroutine_and_channel/file1.txt")
					if err1 != nil {
						fmt.Println(err1)
						break
					}
					// 去掉两头空格，并分割为单个数字
					split := strings.Split(strings.TrimSpace(string(bytes)), " ")
					nums := make([]int, len(split))
					// 转化为数字类型
					for i := range split {
						nums[i], _ = strconv.Atoi(split[i])
					}
					// 排序
					slices.Sort(nums)
					fmt.Println("文件排序成功！", nums)
					// 写入文件2
					file2, err1 := os.OpenFile("./advanced/goroutine_and_channel/file2.txt", os.O_WRONLY|os.O_CREATE, 0777)
					if err1 != nil {
						fmt.Println(err1)
						file2Chan <- false
					} else {
						writer := bufio.NewWriter(file2)
						for _, num := range nums {
							writer.WriteString(strconv.Itoa(num) + " ") // 写入文件
						}
						fmt.Println("文件2写入成功！")
						writer.Flush()
						file2Chan <- true
					}
					// 关闭文件
					defer func() {
						if closeErr := file2.Close(); closeErr != nil {
							// 处理关闭文件时的错误，例如记录日志等
							println("关闭文件时发生错误：", closeErr)
						} else {
							println("文件已成功关闭")
						}
					}()
				}
				close(file2Chan)
				break
			}
		}
	}
	// 开启协程
	go writeDataToFile()
	go sort()
	// 主线程等待两个协程完成
	for {
		if v, ok := <-file2Chan; ok {
			if v {
				fmt.Println("写入并排序完成！")
			} else {
				fmt.Println("出现错误！")
			}
			break
		}
	}
}

// 用n个协程计算1-80000中，哪些是素数
func anli4(n int) {
	numChan := make(chan int, 80000)
	primeChan := make(chan int, 80000)
	exitChan := make(chan bool, n)
	go func() {
		// 写入80000个数据
		for i := 1; i <= 80000; i++ {
			numChan <- i
		}
		close(numChan)
	}()
	// 判断数字是否是素数
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	xc := func() {
		for {
			if v, ok := <-numChan; !ok {
				break
			} else {
				if isPrime(v) {
					primeChan <- v
				}
			}
		}
		exitChan <- true
	}
	for i := 0; i < n; i++ {
		go xc()
	}
	for {
		if len(exitChan) >= n {
			close(exitChan)
			close(primeChan)
			fmt.Println("1-80000中素数有：")
			for i := range primeChan {
				fmt.Print(i, " ")
			}
			fmt.Println()
			break
		}
	}
}
func main() {
	// 统计效率
	start := time.Now().UnixNano()
	anli4(4) // 4个协程耗时：27ms
	//anli4(8) // 8个协程耗时：17ms
	end := time.Now().UnixNano()
	fmt.Printf("耗时：%dms\n", (end-start)/1e6)
}
