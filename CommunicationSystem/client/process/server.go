package process

import (
	"Go_Learn/CommunicationSystem/client/utils"
	"fmt"
	"net"
	"os"
)

// 显示登陆后的界面
func ShowMenu() {
	fmt.Println("--------------------恭喜xxx登录成功--------------------")
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.信息列表")
	fmt.Println("4.退出系统")
	fmt.Print("请输入你的选择：")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入无效！")
	}
}

// 和服务端保持通讯
func serverProcessMes(conn net.Conn) {
	// 创建一个Transfer实例，不停的读取服务端发送的信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务端发送的信息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("trl.ReadPkg err=", err)
			return
		}
		// 这里我们又根据mes的类型来处理不同的信息
		fmt.Println("mes=", mes)
	}

}
