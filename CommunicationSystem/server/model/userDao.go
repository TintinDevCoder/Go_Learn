package model

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9" // 确保 go.mod 中已拉取 v9
)

var (
	// 全局变量，方便各层调用
	MyUserDao *UserDao
	// go-redis 几乎所有操作都需要 ctx
	ctx = context.Background()
)

type UserDao struct {
	client *redis.Client
}

// NewUserDao 工厂模式：传入初始化好的 client
func NewUserDao(client *redis.Client) *UserDao {
	return &UserDao{
		client: client,
	}
}

// GetUserById 根据用户 id 返回 User 实例
// 注意：在 v9 中，推荐将 context 作为第一个参数传入，以便处理超时和链路追踪
func (this *UserDao) GetUserById(id int) (user *User, err error) {
	key := fmt.Sprintf("user:%d", id)

	// 1. 初始化一个结构体实例
	user = &User{}

	// 2. 使用 HGetAll 获取该 Hash 的所有字段
	// Scan 会根据结构体里的 `redis:"xxx"` 标签自动把字段对号入座
	err = this.client.HGetAll(ctx, key).Scan(user)

	if err != nil {
		return nil, err
	}

	// 3. 关键判断：HGetAll 在 Key 不存在时不会返回 redis.Nil，而是返回一个空结果
	// 所以我们需要通过判断某个必填字段（如 UserId）是否为初始值来确定是否查到了
	if user.UserId == 0 {
		return nil, fmt.Errorf("用户 %d 不存在", id)
	}

	return user, nil
}

// 登录校验Login
func (this *UserDao) Login(userId int, userPwd string) (err error) {
	user, err := this.GetUserById(userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
