package main

import "fmt"

func main() {

	key := ""                       // 声明一个变量，保存接收用户输入的选项
	loop := true                    // 声明一个变量，控制是否退出 for
	balance := 10000.0              // 定义账户的余额
	money := 0.0                    // 每次收支的金额
	note := ""                      // 每次收支的说明
	flag := false                   // 定义一个变量，记录是否有收支行为
	details := "收支\t账户金额\t收支金额\t说明" // 收支的详情，当有收支时，只需要对 details 进行拼接处理即可

	// 显示主菜单
	for {
		if !loop {
			break
		}
		fmt.Println("\n\n==================== 家庭手持记账软件 ====================")
		fmt.Println("                     1. 收支明细")
		fmt.Println("                     2. 登记收入")
		fmt.Println("                     3. 登记支出")
		fmt.Println("                     4. 退出软件")
		fmt.Println("请选择 (1-4): ")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if flag {
				fmt.Println("-------------------- 当前收支明细 --------------------")
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细... 来一笔吧!")
			}
		case "2":
			fmt.Println("登记收入...")
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			if money < 0 {
				fmt.Println("输入金额有问题...")
			} else {
				balance += money
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = true
			}
		case "3":
			fmt.Println("登记支出...  ")
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			if money < 0 {
				fmt.Println("输入金额有问题")
			} else {
				if money > balance {
					fmt.Println("剩余金额不够")
				} else {
					balance -= money
					details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				}
			}
		case "4":
			fmt.Println("你确定要退出吗? y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项...")
		}
	}
	fmt.Println("你退出了家庭记账软件的使用...")
}
