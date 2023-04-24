package main

import (
	"fmt"
)

func main() {
	//用户输入
	key := ""
	//控制退出for
	loop := true
	//用户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//每次收支的详情
	details := "收支\t账户金额\t收支金额\t说   明"
	//记录是否有收支行为
	flag := false
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.登记收入")
		fmt.Println("                   3.登记支出")
		fmt.Println("                   4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n-----------------收支明细-----------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("暂无收支明细")
			}

		case "2":
			fmt.Println("\n-----------------登记收入-----------------")
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			flag = true
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
		case "3":
			fmt.Println("\n-----------------登记支出-----------------")
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足,请重新输入")
				break
			}
			balance -= money
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			flag = true
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
		case "4":
			fmt.Println("确定退出吗? y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入错误,请重新输入")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("输入错误,请重新输入1-4")
		}
		if !loop {
			break
		}
	}
	fmt.Println("软件退出")
}
