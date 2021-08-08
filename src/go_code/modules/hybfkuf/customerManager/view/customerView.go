package main

import (
	"fmt"
	"hybfkuf/customerManager/model"
	"hybfkuf/customerManager/service"
)

type customerView struct {
	// 定义必要的字段
	key             string // 接收用户的输入
	loop            bool   // 表示是否循环显示菜单
	customerService *service.CustomerService
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("==================== 客户信息管理软件 ====================")
		fmt.Println("                     1. 添加账户")
		fmt.Println("                     2. 修改客户")
		fmt.Println("                     3. 删除客户")
		fmt.Println("                     4. 客户列表")
		fmt.Println("                     5. 退出")
		fmt.Println("请选择 (1-5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出了客户关系管理系统")
}

// 显示所有的客户信息
func (this *customerView) list() {
	// 首先，获取当前所有的客户信息（在切片中）
	customers := this.customerService.List()
	fmt.Println("-------------------- 客户列表 --------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n-------------------- 客户列表完成 --------------------\n\n")
}

// 得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) add() {
	name := ""
	gender := ""
	age := 0
	phone := ""
	email := ""
	fmt.Println("-------------------- 添加客户 --------------------")
	fmt.Println("姓名:")
	fmt.Scanln(&name)
	fmt.Println("性别:")
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	fmt.Scanln(&age)
	fmt.Println("电话")
	fmt.Scanln(&phone)
	fmt.Println("电邮")
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------------------- 添加完成 --------------------")
	}
}

// 得到用户的输入id，删除该 id 对应的客户
func (this *customerView) delete() {
	fmt.Println("-------------------- 删除客户 --------------------")
	fmt.Println("请选择待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "y" {
		// 调用 customerService 的 Delete 方法
		if this.customerService.Delete(id) {
			fmt.Println("-------------------- 删除成功 --------------------")
		} else {
			fmt.Println("-------------------- 删除失败, ID不存在 --------------------")
		}
	}
}

// 退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N): ")
	for {
		fmt.Println(&this.key)
		if this.key == "y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N): ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

func main() {
	fmt.Println("hello golang")
	// 在主函数中创建一个 customerView，并运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
