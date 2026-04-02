package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 统计1-n的数字中，哪些是素数
// 传统方法是使用循环，判断各个数字是不是素数
// 使用并发或并行，将统计任务分配给多个goroutine

// 线程和协程：线程和协程都是并发执行，但线程是操作系统调度的，而协程是Go语言调度的。
// 线程： 是操作系统调度的，比如操作系统的线程调度算法是抢占式调度，即多个线程并发执行，当一个线程执行时间片用完时，操作系统会选择另一个线程继续执行。
// 协程： 是Go语言调度的，比如Go语言的协程调度算法是抢占式调度，即多个协程并发执行，当一个协程执行时间片用完时，Go语言会选择另一个协程继续执行。

// go主线程：主线程是Go语言的入口，所有Go语言程序都是从主线程开始执行。
// 主线程上可以有多个协程，协程是Go语言的并发执行单元，协程可以理解为轻量级的线程
// 协程的特点：1、有独立的栈 2、共享程序堆空间 3、调度由用户控制 4、是轻量的线程
// goroutine:
func hello(s string) {
	for i := range 10 {
		fmt.Println(s + " goroutine hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func test() {
	go hello("1") // 开启一个协程
	for i := range 10 {
		fmt.Println("main hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// 如果主线程退出，协程就算没有执行完，也会退出
	// 主线程是一个重量级的物理线程。协程是一个轻量级的逻辑线程，资源消耗小
	// golang的并发上有很大的性能优势，因为Go语言的协程是轻量级的逻辑线程，资源消耗小，可以快速创建和销毁，因此Go语言的并发性能高。
}

// goroutine的调度模型：MPG模型
// M (Machine)：内核线程 (Kernel Thread)。由操作系统调度，是真正的执行单元。M必须绑定一个 P才能执行G。
// P (Processor)：逻辑处理器/上下文。它是调度器的核心，维护着一个本地运行队列。P 的数量由 $GOMAXPROCS 决定，通常等于CPU核心数。即有多少个G可以同时运行
// G (Goroutine)：用户态协程。包含栈空间、指令指针以及调度所需的状态（如 _Grunning）。它非常轻量（初始仅 2KB）。
// 每启动一个G，就会在P维护的runqueue队列末尾加一个G。当P空闲时，会从runqueue队列中取出G并运行。

// MPG运行机制：
// 1.在 Go程序启动时，会在一个cpu中运行，系统会执行一系列初始化操作，然后 Runtime 会创建一个特殊的 Goroutine，被称为 Main Goroutine
//   当运行程序时，main 协程是第一个被创建并投入执行的 G，如果 main 协程执行结束（返回），整个进程就会立即退出，无论其他子协程（Child Goroutines）是否还在运行
// 2.程序启动时，P 的数量就固定了，默认等于CPU Cores数量。并发是在已有的 P 之间调度
//   当程序中创建的协程数量超过 P 的数量，那么多余的协程就会在已存在的 P 之间进行调度
//   这里并不是简单的“挂在队列”，而是M在执行一个 G 一段时间后（或遇到阻塞时），调度器会保存当前 G 的上下文，将其放回队列末尾，并取出下一个 G 执行。这保证了公平性。
//   程序中创建的协程数量小于 P 的数量，那么多余的协程就会在已存在的 P 之间进行调度
//   如果某个 P 的队列空了，它会触发 Work Stealing (工作窃取) 算法，去别的 P 那里偷一半的 G 过来跑
// 3.当程序中有新开的协程，新的 G 会被分配到负载较轻的 P 队列中，每个 P 都会寻找（或创建）一个 M 来绑定
//   此时，不同的 M 在不同的物理核心上同时运行不同的 G
//   虽然 P 的数量受 CPU 核心限制，但 M（内核线程）的数量默认上限可以达到 10000。这是为了防止大量阻塞调用导致程序彻底卡死。

// MPG的移交机制：
// 1. G0协程在M0主线程上执行，另外有三个协程在队列等待。
// 2. G0协程在M0主线程上阻塞，比如读文件等。
// 3. 这时会创建一个M1主线程（或使用已存在的线程），将等待的G1、G2、G3加入到M1主线程的队列中。M0仍然执行G0协程，G1、G2、G3会从M1主线程的队列中取出并执行。
// 4. 当 G0 的读文件操作完成后，M0 会被放到空闲的主线程继续执行（从已有的线程池取），同时G0被唤醒
// 这样做可以防止由于一个 G 的阻塞导致整个 P（及其队列中的其他 G）被“饿死”

// goroutine的问题
// 不同的goroutine之间可能会出现数据竞争，导致数据不一致。
// 此时就需要goroutine之间进行通信，使用channel解决

// 计算 1-200 各个数的阶乘，并把各个数的阶乘放入到 map 中，通过 goroutine
var (
	myMap = make(map[int]int, 10)
)

// 计算n的阶乘
func jc1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

// 通过goroutine分别计算1-200的阶乘
// 此时会报错 fatal error: concurrent map writes
/*
在 Go 的 MPG Model 下，当开启多个协程并行执行时，如果它们都在执行 main.jc 函数并尝试修改同一个全局 map：
	Race Condition (竞态条件)：两个或多个协程可能在同一时间修改哈希表的内部结构（例如触发扩容或修改同一个 Bucket）。
	Panic 保护：为了防止产生不可预知的内存损坏，Go 运行时检测到这种行为后会直接抛出 Fatal Error 强制退出。
*/
func test1() {
	// 开启多个协程，计算1-200的阶乘
	for i := 1; i <= 200; i++ {
		go jc1(i)
	}
	//输出结果
	for i := 1; i <= 200; i++ {
		fmt.Printf("%d! = %d\n", i, myMap[i])
	}
	// 休眠10秒
	time.Sleep(10 * time.Second)
}

// 解决方式一：加锁
// 将全局变量加锁，保证多个协程在修改全局变量时，只有一个协程可以访问
var (
	myMap2 = make(map[int]int, 10)
	// 声明一个全聚德互斥锁
	// lock是全局的互斥锁
	// syn是包：synchornized同步
	lock sync.Mutex
)

func jc2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}
func test2(n int) {
	// 开启多个协程，计算1-200的阶乘
	for i := 1; i <= n; i++ {
		go jc2(i)
	}
	// 休眠10秒,等待计算完成
	time.Sleep(10 * time.Second)
	//输出结果
	for i := 1; i <= n; i++ {
		fmt.Printf("%d! = %d\n", i, myMap[i])
	}
}

// 也可以使用channel来解决竞争问题，在channel.go中进行详解

func main() {
	test2(20)
}
