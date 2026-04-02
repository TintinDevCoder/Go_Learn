package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Customer struct {
	name  string
	sex   string
	age   int
	phone string
	email string
}
type CustomerManager struct {
	customers []Customer
}

// 添加客户信息
func (cm *CustomerManager) addCustomer() {
	var name string
	var sex string
	var age int
	var phone string
	var email string
	fmt.Println("------------添加客户------------")
	fmt.Printf("姓名：")
	fmt.Scanln(&name)
	fmt.Printf("性别：")
	fmt.Scanln(&sex)
	fmt.Printf("年龄：")
	fmt.Scanln(&age)
	fmt.Printf("电话：")
	fmt.Scanln(&phone)
	fmt.Printf("邮箱：")
	fmt.Scanln(&email)
	customer := Customer{name: name, sex: sex, age: age, phone: phone, email: email}
	cm.customers = append(cm.customers, customer)
	fmt.Println("------------添加成功------------！")
}

// 修改客户信息
func (cm *CustomerManager) alterCustomer() {
	var name string
	var sex string
	var age int
	var phone string
	var email string
	fmt.Println("------------修改客户------------")
	fmt.Println("--------直接回车表示不修改----------")
	fmt.Printf("请选择待修改客户编号(-1退出)：")
	var index int
	fmt.Scanln(&index)
	if index == -1 {
		return
	}
	if index < 0 || index > len(cm.customers) {
		fmt.Println("无效的客户编号")
		return
	}
	c := &cm.customers[index-1]
	fmt.Printf("姓名(%s)：", c.name)
	fmt.Scanln(&name)
	if name != "" {
		c.name = name
	}

	fmt.Printf("性别(%s)：", c.sex)
	fmt.Scanln(&sex)
	if sex != "" {
		c.sex = sex
	}

	fmt.Printf("年龄(%d)：", c.age)
	fmt.Scanln(&age)
	if age > 0 {
		c.age = age
	}

	fmt.Printf("电话(%s)：", c.phone)
	fmt.Scanln(&phone)
	if phone != "" {
		c.phone = phone
	}

	fmt.Printf("邮箱(%s)：", c.email)
	fmt.Scanln(&email)
	if email != "" {
		c.email = email
	}
	fmt.Println("------------修改成功------------！")
}

// 查看客户信息
func (cm *CustomerManager) showCustomer() {
	fmt.Println("------------客户列表------------")
	// 参数说明：输出流、最小列宽、tab宽度、填充空格数、填充字符、标志
	w := tabwriter.NewWriter(os.Stdout, 8, 1, 2, ' ', 0)

	// 列标题，使用 \t 分隔
	fmt.Fprintln(w, "编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i, c := range cm.customers {
		fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%s\n",
			i+1, c.name, c.sex, c.age, c.phone, c.email)
	}

	w.Flush() // 必须调用 Flush 才会写入终端
	fmt.Println("-----------客户列表完成-----------")
}

// 删除客户信息
func (cm *CustomerManager) deleteCustomer() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	var index int
	fmt.Scanln(&index)
	if index == -1 {
		return
	}
	if index < 0 || index > len(cm.customers) {
		fmt.Println("无效的客户编号")
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm != "Y" && confirm != "y" {
		fmt.Println("------------取消删除------------")
		return
	}
	cm.customers = append(cm.customers[:index-1], cm.customers[index:]...)
	fmt.Println("------------删除完成------------")
}
func main() {
	choice := 0
	cm := CustomerManager{}
	for choice != 5 {
		fmt.Println("------------客户信息管理系统------------")
		fmt.Println("1. 添加客户信息")
		fmt.Println("2. 删除客户信息")
		fmt.Println("3. 修改客户信息")
		fmt.Println("4. 查询客户信息")
		fmt.Println("5. 退出系统")
		fmt.Printf("请输入您的选择：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			cm.addCustomer()
		case 2:
			cm.deleteCustomer()
		case 3:
			cm.alterCustomer()
		case 4:
			cm.showCustomer()
		case 5:
			fmt.Println("退出系统")
			break
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
}
