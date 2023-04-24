package main

import (
	"Project/2.CustomerService/model"
	"Project/2.CustomerService/service"
	"fmt"
)

type CustomerView struct {
	key             string //接受用户输入
	loop            bool   //判断是否继续循环主菜单
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (c *CustomerView) listCustomer() {
	customers := c.customerService.List()
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------客户列表完成-----------------")
}

// 添加客户
func (c *CustomerView) addCustomer() {
	fmt.Println("-----------------添加客户-----------------")
	fmt.Println("请输入客户姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入客户性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入客户年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入客户电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入客户邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if c.customerService.Add(customer) {
		fmt.Println("-----------------添加客户完成-----------------")
	} else {
		fmt.Println("-----------------添加客户失败-----------------")
	}
}
func (c *CustomerView) deleteCustomer() {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Println("请输入要删除的客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除?(y/n)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			fmt.Println("删除成功")
			break
		}
		if choice == "n" || choice == "N" {
			return
		}
		fmt.Println("输入有误,请重新输入")
	}
	if c.customerService.Delete(id) {
		fmt.Println("-----------------删除客户完成-----------------")
	} else {
		fmt.Println("-----------------删除客户失败-----------------")
	}
}

// 修改客户信息
func (c *CustomerView) modifyCustomer() {
	fmt.Println("-----------------修改客户-----------------")
	fmt.Println("请选择待修改客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := c.customerService.FindById(id)
	if index == -1 {
		fmt.Println("客户编号不存在")
		return
	}
	customer := c.customerService.List()[index]
	fmt.Printf("姓名(%v):", customer.Name)
	name := ""
	fmt.Scanln(&name)
	if name != "" {
		customer.Name = name
	}
	fmt.Printf("性别(%v):", customer.Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender != "" {
		customer.Gender = gender
	}
	fmt.Printf("年龄(%v):", customer.Age)
	age := 0
	fmt.Scanln(&age)
	if age != 0 {
		customer.Age = age
	}
	fmt.Printf("电话(%v):", customer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone != "" {
		customer.Phone = phone
	}
	fmt.Printf("邮箱(%v):", customer.Email)
	email := ""
	fmt.Scanln(&email)
	if email != "" {
		customer.Email = email
	}
	if c.customerService.Update(index, customer) {
		fmt.Println("-----------------修改客户完成-----------------")
	} else {
		fmt.Println("-----------------修改客户失败-----------------")
	}
}

func (c *CustomerView) exit() {
	fmt.Println("是否退出软件?(y/n)")
	for {
		fmt.Scanln(&c.key)
		if c.key == "y" || c.key == "Y" {
			c.loop = false
			break
		}
		if c.key == "n" || c.key == "N" {
			c.loop = true
			break
		}
		fmt.Println("输入有误,请重新输入(y/n)")
	}
}

func (c *CustomerView) mainMenu() {
	for {
		fmt.Println("\n-----------------客户信息管理软件-----------------")
		fmt.Println("                   1.添加客户")
		fmt.Println("                   2.修改客户")
		fmt.Println("                   3.删除客户")
		fmt.Println("                   4.客户列表")
		fmt.Println("                   5.退出软件")
		fmt.Println("请选择(1-5):")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.addCustomer()
		case "2":
			c.modifyCustomer()
		case "3":
			c.deleteCustomer()
		case "4":
			c.listCustomer()
		case "5":
			c.exit()
		default:
			fmt.Println("输入有误,请重新输入1-5")
		}
		if !c.loop {
			break
		}
	}
	fmt.Println("软件退出")
}

func main() {
	customerView := CustomerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
	customerView.mainMenu()
}
