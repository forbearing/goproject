package main

import "fmt"

type familyAccount struct {
	key     string  // 声明一个变量，保存接收用户输入的选项
	loop    bool    // 声明一个变量，控制是否退出 for
	balance float64 // 定义账户的余额
	money   float64 // 每次收支的金额
	note    string  // 每次收支的说明
	flag    bool    // 定义一个变量，记录是否有收支行为
	details string  // 收支的详情，当有收支时，只需要对 details 进行拼接处理即可
}

// 显示账户明细
func (this *familyAccount) ShowDetails() {
	if this.flag {
		fmt.Println("-------------------- 当前收支明细 --------------------")
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

// 登记收入
func (this *familyAccount) income() {
	fmt.Println("登记收入...")
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 登记支出
func (this *familyAccount) pay() {
	fmt.Println("登记支出...  ")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	if this.money > this.balance {
		fmt.Println("剩余金额不够")
	} else {
		this.balance -= this.money
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	}
}

// 退出系统
func (this *familyAccount) exit() {
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
		this.loop = false
	}
}

// 显示主菜单
func (this *familyAccount) mainMenu() {
	for {
		fmt.Println("\n\n==================== 家庭手持记账软件 ====================")
		fmt.Println("                     1. 收支明细")
		fmt.Println("                     2. 登记收入")
		fmt.Println("                     3. 登记支出")
		fmt.Println("                     4. 退出软件")
		fmt.Println("请选择 (1-4): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.ShowDetails() // 收支明细
		case "2":
			this.income() // 登记收入
		case "3":
			this.pay() //登记支出
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
	}
}

// 编写一个工厂模式的构造方法，返回一个 *familyAccount 实例
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

func main() {
	fmt.Println("这是一个面向对象的方式完成...")
	NewFamilyAccount().mainMenu()
}
