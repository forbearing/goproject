package service

import "hybfkuf/customerManager/model"

// 该结构体完成对 customer 的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片包含多少个客户，该字段后面，还可以作为新客户的 id+1
	customerNum int
}

// 编写一个方法，可以发挥 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@mail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到切片
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据 id 查找客户在切片中对应下标，如果没有该客户，返回 -1
func (this *CustomerService) FindById(id int) int {
	index := -1
	// 遍历 this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 根据 Id 删除客户(从切片中删除)
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
