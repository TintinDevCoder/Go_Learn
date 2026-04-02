package main

import (
	"fmt"
	"time"
)

type Account struct {
	accountValue float64
	value        float64
	remark       string
	time         string
}
type FamilyAccount struct {
	accounts   []Account
	totalValue float64
}

// 收支明细
func (f *FamilyAccount) transactionDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	// 使用 \t 进行物理分隔
	fmt.Println("类型\t\t账户金额\t\t收支金额\t\t说明\t\t\t时间")
	for _, account := range f.accounts {
		typeName := "收入"
		val := account.value
		if val < 0 {
			typeName = "支出"
			val = -val
		}
		// \t\t 确保金额和说明之间有足够空间
		fmt.Printf("%s\t %10.2f\t %10.2f\t\t%s\t\t\t%s\n", typeName, account.accountValue, val, account.remark, account.time)
	}
}

// 登记收入和支出
func (f *FamilyAccount) log(is bool) {
	var v float64
	var r string
	fmt.Printf("本次输入金额：")
	fmt.Scanln(&v)
	fmt.Printf("本次支出说明：")
	fmt.Scanln(&r)
	if is {
		f.totalValue += v
	} else {
		f.totalValue -= v
	}
	now := time.Now()
	str := now.Format("2006/01/02 15:04:05")
	f.accounts = append(f.accounts, Account{accountValue: f.totalValue, value: v, remark: r, time: str})
}
func main() {
	choice := 0
	familyaccount := FamilyAccount{accounts: make([]Account, 0)}
	for choice != 4 {
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1 收支明细")
		fmt.Println("                 2 登记收入")
		fmt.Println("                 3 登记支出")
		fmt.Println("                 4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			familyaccount.transactionDetails()
		case 2:
			familyaccount.log(true)
		case 3:
			familyaccount.log(false)
		case 4:
			break
		default:
			fmt.Println("输入有误，请重新输入")
			choice = 0
			continue
		}
	}
}
