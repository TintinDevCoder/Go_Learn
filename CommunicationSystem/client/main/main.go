package main

import (
	"Go_Learn/CommunicationSystem/client/process"
	"fmt"
)

var userId int
var userPwd string

func main() {
	key := 0
	loop := true
	for loop {
		fmt.Println("--------------------欢迎登录多人聊天系统--------------------")
		fmt.Println("1.登录聊天室")
		fmt.Println("2.注册用户")
		fmt.Println("3.退出系统")
		fmt.Print("请输入你的选择：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			// 登录逻辑
			loop = false // 登录成功后退出循环
			fmt.Println("请输入用户的id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码：")
			fmt.Scanf("%s\n", &userPwd)

			// 创建一个UserProcess实例，调用Login方法
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)

			if err != nil {
				fmt.Println("登陆失败")
			} else {
				fmt.Println("登陆成功")
			}
		case 2:
			// 注册逻辑
		case 3:
			loop = false
		default:
			fmt.Println("输入无效，退出系统！")
			loop = false
		}
	}
}
