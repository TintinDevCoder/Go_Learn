package main

import (
	"fmt"
	"runtime"
	"syscall"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	getCurrentThreadId = kernel32.NewProc("GetCurrentThreadId")
)

// 获取 Windows 环境下的当前物理线程 ID
func getTID() uint32 {
	ret, _, _ := getCurrentThreadId.Call()
	return uint32(ret)
}

// 修改为计算密集型函数
func hello2(id string) {
	for i := 0; i < 10; i++ {
		// 模拟耗时计算，不让出 CPU，强制让调度器在多核间切换
		count := 0
		for j := 0; j < 1000000000; j++ {
			count += j
		}
		// 减少打印频率，只打印关键节点
		fmt.Printf("[Goroutine %s] 运行在线程 TID: %d | 进度: %d\n", id, getTID(), i)
	}
}

func testofcpu() {
	// 关键：限制为 2 个核心
	runtime.GOMAXPROCS(2)

	fmt.Println("Starting with 2 CPU cores...")

	// 开启两个计算密集型协程
	go hello2("A")
	go hello2("B")

	// 主协程也参与竞争
	for i := 0; i < 10; i++ {
		count := 0
		for j := 0; j < 1000000000; j++ {
			count += j
		}
		fmt.Printf("[Main Thread] 运行在线程 TID: %d | 进度: %d\n", getTID(), i)
	}
}
func main() {
	/*	// 设置golang运行的cpu核心数量
		// runtime包提供和go运行时环境的互操作，如控制go程序的函数
		num := runtime.NumCPU() // 获取cpu核心数量
		fmt.Println("cpu num:", num)
		runtime.GOMAXPROCS(num) // 设置golang运行时可同时运行的cpu核心数量
	*/
	testofcpu()
}
