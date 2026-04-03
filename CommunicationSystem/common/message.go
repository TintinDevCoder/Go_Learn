package common

// 客户端和服务端发送消息默认是先发送4个字节的信息，说明发送信息的长度，然后再发送Message类型的传输数据
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
