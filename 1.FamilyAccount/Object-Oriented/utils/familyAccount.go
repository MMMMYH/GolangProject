package utils

import "fmt"

type FamilyAccount struct {
	//用户输入
	key string
	//控制退出for
	loop bool
	//用户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//每次收支的详情
	details string
	//记录是否有收支行为
	flag bool
}

//工厂模式
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   10000.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说   明",
		flag:    false,
	}
}

//显示明细
func (this *FamilyAccount) ShowDetails() {
	fmt.Println("\n-----------------收支明细-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("暂无收支明细")
	}
}

//登记收入
func (this *FamilyAccount) AddIncome() {
	fmt.Println("\n-----------------登记收入-----------------")
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

//登记支出
func (this *FamilyAccount) AddOutcome() {
	fmt.Println("\n-----------------登记支出-----------------")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足,请重新输入")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

//退出系统
func (this *FamilyAccount) Exit() {
	fmt.Println("确定退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入错误,请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
	fmt.Println("软件退出")
}

//主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.登记收入")
		fmt.Println("                   3.登记支出")
		fmt.Println("                   4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.ShowDetails()
		case "2":
			this.AddIncome()
		case "3":
			this.AddOutcome()
		case "4":
			this.Exit()
		default:
			fmt.Println("输入错误,请重新输入1-4")
		}
		if !this.loop {
			break
		}
	}
}
