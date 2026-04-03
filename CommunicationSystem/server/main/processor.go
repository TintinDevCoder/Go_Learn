package main

import (
	"Go_Learn/CommunicationSystem/common"
	process2 "Go_Learn/CommunicationSystem/server/process"
	"Go_Learn/CommunicationSystem/server/utils"
	"errors"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的信息种类不同，决定调用函数信息
func (this *Processor) serverProcessMes(mes *common.Message) (err error) {
	// 根据类型处理不同
	switch mes.Type {
	case common.LoginMesType:
		up := process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case common.RegisterMesType:

	default:
		fmt.Println("消息类型不存在，无法处理")
		err = errors.New("消息类型不存在，无法处理")
	}
	return
}
func (this *Processor) process2() error {
	// 循环等待客户端发送信息
	for {
		tf := &utils.Transfer{Conn: this.Conn}
		fmt.Printf("等待接收客户端 %s 的数据\n", this.Conn.RemoteAddr().String())
		mesg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("连接已关闭，无法接收数据")
			} else {
				fmt.Println("ReadPkg err：", err)
			}
			return err
		}
		err = this.serverProcessMes(&mesg)
		if err != nil {
			return err
		}
	}
	return nil
}
