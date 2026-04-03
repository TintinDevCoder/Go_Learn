package process

import (
	"Go_Learn/CommunicationSystem/common"
	"Go_Learn/CommunicationSystem/server/model"
	"Go_Learn/CommunicationSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

// 登录处理
func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {
	// 取出用户信息，并反序列化
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	// 定义返回信息
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	// 声明LoginResMes
	var loginResMes common.LoginResMes
	// 登录校验
	err = model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err == model.ERROR_USER_PWD {
		loginResMes.Code = 500
		loginResMes.Error = "账号或密码错误"
		return
	} else if err != nil {
		return
	}
	// 登录未出错
	loginResMes.Code = 200
	// 将声明LoginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	resMes.Data = string(data)

	// 将返回信息序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	// 发送信息
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	return
}
