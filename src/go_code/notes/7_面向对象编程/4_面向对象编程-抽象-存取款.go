package main

import "fmt"

type Account struct {
	AccountNo string
	Passwd    string
	Balance   float64
}

func (account *Account) Deposite(money float64, passwd string) {
	// 密码是否正确
	if passwd != account.Passwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	// 查看存款余额是否正确
	if money < 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) Withdraw(money float64, passwd string) {
	if passwd != account.Passwd {
		fmt.Println("你的密码不正确")
		return
	}
	if money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(passwd string) {
	if passwd != account.Passwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号=%v 余额=%f\n", account.AccountNo, account.Balance)
}

func main() {
	account := &Account{
		AccountNo: "88888888",
		Passwd:    "123456",
		Balance:   1000000,
	}
	account.Query("123456")
	account.Deposite(8888, "123456")
	account.Query("123456")
	account.Withdraw(7777, "123456")
	account.Query("123456")

}
