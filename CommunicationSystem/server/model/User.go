package model

type User struct {
	UserId   int    `json:"userId" redis:"userId"`
	UserPwd  string `json:"userPwd" redis:"userPwd"`
	UserName string `json:"userName" redis:"userName"`
}
