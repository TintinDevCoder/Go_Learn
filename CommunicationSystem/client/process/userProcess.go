package process

import (
	"Go_Learn/CommunicationSystem/client/utils"
	"Go_Learn/CommunicationSystem/common"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

// 登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 准备发送消息给服务端
	var mes common.Message
	mes.Type = common.LoginMesType

	// 封装一个LoginMes结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Mershal err=", err)
		return
	}
	// 将用户信息封装进消息
	mes.Data = string(data)

	// mes的序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mershal err=", err)
		return
	}

	trl := utils.Transfer{Conn: conn}
	// 将消息发送
	err = trl.WritePkg(data)
	if err != nil {
		return
	}

	// 处理返回信息
	mes, err = trl.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	// 将返回LoginResMes消息反序列化
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if loginResMes.Code != 200 {
		fmt.Println(loginResMes.Error)
		return
	}
	return
}
