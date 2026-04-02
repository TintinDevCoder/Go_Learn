package main

// 抽象是一种编程方法，思维方式
// 抽象是指将对象的共性提取出来，形成一个抽象的概念。
// 抽象可以通过接口来实现。接口是一种类型，它定义了一组方法，但不实现这些方法。接口可以被任何类型实现，只要它实现了接口定义的方法。
type Account struct {
	AccountNumber string
	Balance       float64
	pwd           string
}

// 方法
// 存款
func (a *Account) Deposit(amount float64, pwd string) {
	// 查看输入的密码是否正确
	if a.pwd != pwd {
		println("密码错误，存款失败")
		return
	}
	a.Balance += amount
	println("存款成功，当前余额：", a.Balance)
}

// 取款
func (a *Account) Withdraw(amount float64, pwd string) {
	// 查看输入的密码是否正确
	if a.pwd != pwd {
		println("密码错误，取款失败")
		return
	}
	if a.Balance < amount {
		println("余额不足，取款失败")
		return
	}
	a.Balance -= amount
	println("取款成功，当前余额：", a.Balance)
}

func main() {

}
