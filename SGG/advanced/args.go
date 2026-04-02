package main

import (
	"flag"
	"fmt"
	"os"
)

// 运行文件的时候传入参数，获取参数
// os.Args是一个string的切片，用来存储所有的命令行参数
// os.Args是一个全局变量，存储了程序运行时的命令行参数
// 0位置存放的是文件的路径，后面依次是实参
func argsApplication() {
	// 1、使用os.Args获取参数
	fmt.Println("命令行的参数有", len(os.Args))
	// 遍历os.Args切片，得到所有的命令行输入参数
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	/*
		命令行的参数有 4
		args[0]=C:\Users\Administrator\AppData\Local\JetBrains\GoLand2025.3\tmp\GoLand\___2go_build_SGG_advanced.exe
		args[1]=1
		args[2]=2
		args[3]=3
	*/
}
func flagApplication() {
	// 使用flag包来解析命令行参数
	var user string
	var pwd string
	var host string
	var port int
	// mysql -u root -p **** -h 127.0.0.1 -P 3306
	// &user 用于接收用户命令行中 -u 后面的参数
	// name参数为命令行参数的参数名，value参数为命令行参数的默认值，usage参数为命令行参数的提示信息
	flag.StringVar(&user, "u", "root", "用户名")
	flag.StringVar(&pwd, "p", "123456", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "主机地址")
	flag.IntVar(&port, "P", 3306, "端口号")
	// 必须调用flag.Parse()，才能解析命令行参数
	flag.Parse()
	fmt.Println("用户名：", user)
	fmt.Println("密码：", pwd)
	fmt.Println("主机地址：", host)
	fmt.Println("端口号：", port)
}
func main() {

}
