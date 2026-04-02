package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 客户端，连接服务器，发送数据，接收数据
// net.Dial函数用来连接服务器，返回一个Conn对象，该对象表示与服务器的连接，可以用来发送和接收数据
// Conn.Write方法用来发送数据，参数是一个字节切片，可以使用[]byte()函数将字符串转换为字节切片
func main() {
	conn, err1 := net.Dial("tcp", "127.0.0.1:8888") // 参数分别为协议类型和服务器地址和端口
	if err1 != nil {
		fmt.Println("连接服务器失败，错误信息：", err1)
		return
	}
	fmt.Println("连接服务器成功", conn)
	// 发送数据
	reader := bufio.NewReader(os.Stdin) // 从终端读取输入
	var line string = "test"
	var err2 error
	for {
		// 从终端读取一行数据，直到遇到换行符为止
		line, err2 = reader.ReadString('\n')
		if err2 != nil {
			fmt.Println("读取输入失败，错误信息：", err2)
			return
		}
		if line == "exit\n" { // 如果输入的是exit，就退出循环
			fmt.Println("退出客户端")
			break
		}
		// 将读取到的数据发送给服务器
		n, err3 := conn.Write([]byte(line)) // n是发送的字节数，err3是发送数据时的错误信息
		if err3 != nil {
			fmt.Println("发送数据失败，错误信息：", err3)
			return
		}
		fmt.Println("发送数据成功,发送了", n, "个字节")
	}

	defer func() {
		err4 := conn.Close() // 关闭连接
		if err4 != nil {
			fmt.Println("关闭连接失败，错误信息：", err4)
			return
		}
		fmt.Println("关闭连接成功")
	}()
}
