package main

import (
	"fmt"
	"io"
	"net"
)

// 客户端,监听8888端口，如果有连接就开一个协程处理连接
// net包用来处理网络连接，监听端口，发送和接收数据等
// net.Listen函数用来监听指定的网络地址和端口，返回一个Listener对象，
// 该对象可以接受来自客户端的连接请求
// Listener.Accept方法用来接受一个连接请求，返回一个Conn对象，
// 该对象表示与客户端的连接，可以用来发送和接收数据
// 处理连接的函数handleConnection，接收一个Conn对象作为参数，
// 在函数中可以使用Conn对象的方法来发送和接收数据，例如Read和Write方法
func main() {
	fmt.Println("服务器开始监听8888端口...")
	listen, err1 := net.Listen("tcp", "0.0.0.0:8888") // 设置是tcp协议，在本地监听8888端口
	if err1 != nil {
		fmt.Println("监听失败，错误信息：", err1)
		return
	}
	fmt.Println("服务器监听成功", listen)
	// Listener.Accept方法会阻塞，直到有一个连接请求到来。当有连接请求时，Accept方法会返回一个Conn对象，表示与客户端的连接。可以在循环中处理每个连接请求，例如启动一个新的协程来处理每个连接。
	// 第一次运行到 Accept() 就会阻塞等待，而不是在 for 循环里疯狂空转
	for {
		// 等待连接
		fmt.Println("等待连接...")
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("接受连接失败，错误信息：", err2)
			continue
		}
		fmt.Printf("接受连接成功 con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		// 通过协程处理连接
		go handleConnection(conn)
	}

	defer listen.Close() // 关闭监听
}
func handleConnection(conn net.Conn) {
	// 循环接受客户端发送的数据
	buf := make([]byte, 1024) // 创建一个字节切片用来存储接收到的数据，大小为1024字节
	for {
		// 接收数据
		// 等待客户端通过conn发送信息，如果客户端没有write，那么一直阻塞等待
		fmt.Printf("等待接收服务端 %s 的数据\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // n是接收到的字节数，err是接收数据时的错误信息
		if err != nil {
			if err == io.EOF {
				fmt.Println("连接已关闭，无法接收数据")
			} else {
				fmt.Println("接收数据失败，错误信息：", err)
			}
			return
		}
		// 显示接收到的数据，使用string()函数将字节切片转换为字符串，注意只转换接收到的字节数
		fmt.Printf("接收数据成功，接收了%d个字节，数据内容：%s", n, string(buf[:n]))
	}
	defer conn.Close() // 关闭连接
}
