package main

import (
	"fmt"
	"time"
)

// channel 是 Go 语言中的一种数据结构，用于在不同的 goroutine 之间进行通信和同步。
// 它提供了一种安全的方式来传递数据，避免了共享内存带来的竞争条件和死锁问题。
// channel 本质就是一个队列，数据是先进先出的。它是线程安全的，多个 goroutine 可以同时发送和接收数据，而不需要担心数据竞争。
// channel 是有类型的，必须指定要传递的数据类型。可以使用 make 函数创建一个 channel，例如：
// ch := make(chan int) // 创建一个传递 int 类型数据的 channel
// channel是引用类型，本身有地址，但存放的是队列的首地址
func channel() {
	// 初始化管道
	ch1 := make(chan int, 3) // 创建一个传递 int 类型数据的 channel，并指定缓冲区大小为 3
	var ch2 chan int         // 声明一个传递 int 类型数据的 channel，但未初始化

	// 查看管道内容
	fmt.Println(ch1) // 输出: <nil>
	fmt.Println(ch2) // 输出: 0xc0000a8000 (channel 的地址)

	// 向 channel 中发送数据
	ch1 <- 10
	ch1 <- 20

	// 查看管道的长度和cap
	fmt.Println("channel len: ", len(ch1), "cap: ", cap(ch1))

	ch1 <- 30
	//ch1 <- 40 // 这行代码会阻塞报错：fatal error: all goroutines are asleep - deadlock!，因为 channel 的缓冲区已满.

	// 从channel读取数据, 会返回两个值，一个是数据，一个是是否成功读取到数据的布尔值，如果 channel 中没有数据可读了，第二个值会返回 false。
	num1, ok := <-ch1
	if !ok {
		fmt.Println("channel 已经关闭了，无法读取数据了")
	} else {
		fmt.Println("从 channel 中读取的数据: ", num1)
		fmt.Println("channel len: ", len(ch1), "cap: ", cap(ch1))
	}

	// 如果channel内的数据已经全部取出，再取就会报错
	num2 := <-ch1
	num3 := <-ch1
	//num4 := <-ch1 // 这行代码会阻塞报错：fatal error: all goroutines are asleep - deadlock!，因为 channel 中没有数据可读了.
	_, _ = num2, num3

	// 存放map类型数据的 channel
	var mapChan chan map[string]string        // 声明一个传递 map[string]string 类型数据的 channel，但未初始化
	mapChan = make(chan map[string]string, 2) // 创建一个传递 map[string]string 类型数据的 channel，并指定缓冲区大小为 2

	// 向 channel 中发送数据
	mapChan <- map[string]string{"name": "Alice", "age": "30"}
	mapChan <- map[string]string{"name": "Bob", "age": "25"}

	// 从 channel 中读取数据
	person1 := <-mapChan
	person2 := <-mapChan
	fmt.Println("从 channel 中读取的数据: ", person1)
	fmt.Println("从 channel 中读取的数据: ", person2)
	close(mapChan)

	// 存放任意类型数据的 channel
	var anyChan chan interface{}        // 声明一个传递 interface{} 类型数据的 channel，但未初始化
	anyChan = make(chan interface{}, 2) // 创建一个传递 interface{} 类型数据的 channel，并指定缓冲区大小为 2
	// 向 channel 中发送不同类型的数据
	anyChan <- 42
	anyChan <- "Hello, Go!"
	// 从 channel 中读取数据
	data1 := <-anyChan
	data2 := <-anyChan
	fmt.Println("从 channel 中读取的数据: ", data1.(int))
	fmt.Println("从 channel 中读取的数据: ", data2.(string))
	close(anyChan)

	// channel的关闭，只能关闭初始化了的channel
	// 使用内置的函数close可以关闭channel的信道，此时可以读数据，但不能再写入数据
	ch1 <- 56
	close(ch1)
	// 关闭后的channel仍然可以读取数据，但不能再写入数据了
	num4 := <-ch1
	fmt.Println("从 channel 中读取的数据: ", num4)
	// 关闭后的channel不能再写入数据了，否则会报错：panic: send on closed channel
	//ch1 <- 78
}

func visitChannel() {
	// channel的遍历
	intChan1 := make(chan int, 5)
	intChan1 <- 1
	intChan1 <- 2
	intChan1 <- 3
	intChan1 <- 4
	intChan1 <- 5

	// 遍历 channel 中的数据
	// channel只会返回一个数值，没有index
	// 如果在遍历时channel已经关闭，则会正常遍历数据，遍历结束退出遍历
	// 如果在遍历时channel没有关闭，则会报错：fatal error: all goroutines are asleep - deadlock!，因为 channel 中没有数据可读了.
	close(intChan1)

	for num := range intChan1 {
		fmt.Println("从 channel 中读取的数据: ", num)
		// 遍历时，如果channel已经全部取完了，会报错：fatal error: all goroutines are asleep - deadlock!
		// 因此在遍历前需要先关闭 channel，或者在遍历时使用 select 语句来判断 channel 是否已经关闭。
	}

	// for range 遍历 channel 时，其底层的行为逻辑如下：
	//  尝试读取：range 会不断尝试从 channel 中拉取数据。
	//  数据存在：正常取出数据，进入循环体。
	//  数据为空：
	//   如果 channel 已关闭：range 会读取完缓冲区（buffer）里的剩余数据，然后感知到关闭状态，优雅地退出循环。
	//   如果 channel 未关闭：range 会认为“现在没数据，但以后可能会有”，于是它会让当前协程进入 Gwaiting（等待状态），并交出 CPU 执行权。

	// 当调度器发现所有的 Goroutines（包括 main）都处于休眠/阻塞状态，为了防止程序白白耗电，Go 会直接抛出 fatal error: all goroutines are asleep - deadlock!
	// 实验验证如下：intChan2 遍历前未关闭，主函数每过5秒向其发送一条数据
	// 协程遍历时会不断尝试从 channel 中拉取数据，直到主函数发送完10条数据后，关闭intChan2，此时协程会读取完缓冲区里的剩余数据，然后感知到关闭状态，优雅地退出循环。
	intChan2 := make(chan int, 5)
	go func() {
		for num := range intChan2 {
			fmt.Println("从 channel 中读取的数据: ", num)
			// 遍历时，如果channel已经全部取完了，会报错：fatal error: all goroutines are asleep - deadlock!
			// 因此在遍历前需要先关闭 channel，或者在遍历时使用 select 语句来判断 channel 是否已经关闭。
		}
	}()

	for i := 1; i <= 10; i++ {
		intChan2 <- i
		time.Sleep(time.Second * 5)
	}
	close(intChan2)
}
func channelType() {
	// channel可以定义为只读或者只写，默认情况下管道是双向的
	// 只读管道：只能从管道中读取数据，不能向管道中写入数据。声明方式为：var ch chan<- int
	var chan1 <-chan int
	chan1 = make(chan int, 3)
	fmt.Println(chan1)
	//chan1 <- 10 // 报错
	//close(chan1) //报错，因为只读管道不能写入数据，也不能关闭管道

	// 只写管道：只能向管道中写入数据，不能从管道中读取数据。声明方式为：var ch <-chan int
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 10
	//num := <-chan2 // 报错
	close(chan2)
	// 只读和只写管道的使用可以增加代码的可读性和安全性，明确了数据流向，避免了误用。

	chan3 := make(chan int, 3)
	send(chan3)
	receive(chan3)
}

// 参数接收的channel会变为只写的
func send(ch chan<- int) {
	//num := <-ch //报错
	ch <- 1
}

// 参数接收的channel会变为只读的
func receive(ch <-chan int) {
	num := <-ch //报错
	//ch <- 1 // 报错
	fmt.Println("从 channel 中读取的数据: ", num)
}

// channel阻塞问题，使用select可以防止在管道没有数据时，阻塞数据
// 传统方法在遍历管道时，如果不关闭会阻塞而导致死锁
func selectF() {
	ch1 := make(chan int, 10)
	for i := range 10 {
		ch1 <- i
	}
	ch2 := make(chan string, 5)
	for i := range 5 {
		ch2 <- fmt.Sprintf("str%d", i)
	}
	for {
		// 用select语句，如果管道一直没关闭，也不会一直阻塞而死锁
		// 会自动到下一个case匹配，直到其中一个管道有数据可读了，才会继续执行。
		select {
		case v := <-ch1:
			fmt.Println("从 ch1 中读取的数据: ", v)
		case v := <-ch2:
			fmt.Println("从 ch2 中读取的数据: ", v)
		// 如果两个管道都没有数据可读了，select会阻塞等待，直到其中一个管道有数据可读了，才会继续执行。
		default:
			fmt.Println("两个管道都没有数据可读了，等待中...")
			time.Sleep(time.Second * 2)
		}
	}
}

// 在协程中使用recover函数来捕获和处理 panic，可以防止程序崩溃，并允许协程继续执行。
func recoverF() {
	go func() {
		for i := range 10 {
			time.Sleep(time.Second)
			fmt.Println("i: ", i)
		}
	}()
	go func() {
		// 使用defer+recover来捕获 panic，防止程序崩溃
		// defer要写在可能发生错误的语句之前面，这样才能捕获到 panic。
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("捕获到 panic: ", r)
			}
		}()
		var a []int
		a[0] = 1 // 这行代码会引发 panic，因为访问了一个空切片的索引
	}()
	for i := range 10 {
		time.Sleep(time.Second)
		fmt.Println("main i: ", i)
	}
}
func main() {
	// channel的创建、存入、读取和关闭
	//channel()
	// channel的遍历
	//visitChannel()
	// channel的类型
	//channelType()
	// select方法防止管道阻塞死锁
	//selectF()
	// recover函数捕获协程中的panic
	recoverF()
}

// 一个channel中，取出数据可以慢于写入数据，此时如果channel满了，会异步阻塞，但并不会发生死锁，因为还在读取
// 如果只写入，没有取出，那么运行时，编译器会发现死锁，并报错：fatal error: all goroutines are asleep - deadlock!，因为 channel 中没有数据可读了.
