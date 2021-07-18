package account

import "fmt"

type account struct {
	accountNo string
	passwd    string
	balance   float64
}

func NewAccount(accountNo string, passwd string, balance float64) *account {
	// 账号长度 6-10之间
	// 余额必须 > 20
	// 密码必须是六位数
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对...")
		return nil
	}
	if len(passwd) != 6 {
		fmt.Println("密码长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}

	return &account{
		accountNo: accountNo,
		passwd:    passwd,
		balance:   balance,
	}
}

func (a *account) SetPasswd(oldPass string, newPass string) {
	if oldPass != a.passwd {
		fmt.Println("旧密码错误，修改密码失败")
		return
	}
	a.passwd = newPass
}

func (a *account) GetAccountNo(passwd string) {
	if passwd != a.passwd {
		fmt.Println("密码输入错误，获取帐号信息失败")
		return
	}
	fmt.Printf("你的账户: %v\n", a.accountNo)
}
func (a *account) SetAccountNo(passwd string, newAccountNo string) {
	if passwd != a.passwd {
		fmt.Println("密码输入错误，修改账户信息失败")
		return
	}
	a.accountNo = newAccountNo
	fmt.Println("修改账户信息成功")
}

func (a *account) Deposite(passwd string, money float64) {
	if passwd != a.passwd {
		fmt.Println("密码输入错误，存款失败")
		return
	}
	if money < 0 {
		fmt.Println("存款金额错误，存款失败")
		return
	}
	a.balance += money
	fmt.Println("存款成功")
}
func (a *account) Withdraw(passwd string, money float64) {
	if passwd != a.passwd {
		fmt.Println("密码错误，取款失败")
		return
	}
	if money < 0 {
		fmt.Println("取款金额错误，取款失败")
		return
	}
	if money > a.balance {
		fmt.Println("余额不足，取款失败")
		return
	}
	a.balance -= money
	fmt.Println("取款成功")
}
