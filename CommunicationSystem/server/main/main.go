package main

import (
	"Go_Learn/CommunicationSystem/server/model"
	"Go_Learn/CommunicationSystem/server/utils"
	"fmt"
	"net"

	"github.com/redis/go-redis/v9"
)

// 初始化用户数据访问对象
func initUserDao(pool *redis.Client) {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	// 初始化redis连接池
	pool := utils.InitPool("localhost:6379", 16, 0)
	initUserDao(pool)
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败，错误信息：", err)
		return
	}
	fmt.Println("服务器开始监听8888端口...", listen)
	for {
		fmt.Println("等待连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受连接失败，错误信息：", err)
			continue
		}
		fmt.Printf("接受连接成功 con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		go process(conn)
	}
	defer listen.Close()
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	processer := &Processor{
		Conn: conn,
	}
	err := processer.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协议错误：", err)
		return
	}
}
