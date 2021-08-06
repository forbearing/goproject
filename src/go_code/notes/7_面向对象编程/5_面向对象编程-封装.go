package main

import (
	"fmt"
	"hybfkuf/account"
	"hybfkuf/myobject"
	"hybfkuf/person"
	"hybfkuf/utils"
)

/*
封装就是把抽象出的字段和诶对字段的操作封装在一起，数据被保护在内部，程序的其它包只有通过被授权的
操作（方法）才能对字段进行操作

封装的理解和好处
	- 隐藏实现细节
	- 可以IDUI数据进行验证，保证安全合理（Ag e）
如何体现封装
	- 对结构体中国呢的属性进行封装
	- 通过方法，包实现封装
封装的实现步骤
	1:将结构体、字段（属性）的首字母小些（不能导出，其它包不能使用，类似 private）
	2:给结构体所在包提供一个工厂模式的函数，首字母大写。类似一个构造函数
	3:提供一个首字母大写的 Set 方法（类似其它语言的 public），用于对属性判断并赋值
		func (var 结构体类型名) SetXXX(参数列表) (返回值列表) {
			// 加入数据验证的业务逻辑
			var.字段 = 参数
		}
	4:提供一个首字母大写的 Get 方法（类似其它语言的 public），用于获取属性的值
		func (var 结构体类型名) GetXXX() {
			return var.字段 ;
		}
	特别说明：在 Golang 开发中并没有特别强调封装，这点并不像 Java。所以不用总是用 Java 的
			语法特性来看待 golang。Golang 本身对面向对象的特性做了简化。
*/

func main() {

	fmt.Println("hello golang")
	res := utils.Cal(10, 20, '+')
	fmt.Println(res)
	myobject.Hello()

	p := person.NewPerson("hybfkuf")
	fmt.Println(p)
	p.SetAge(100)
	p.SetSal(10000)
	fmt.Println(p)
	fmt.Println("Name:", p.Name, "Age:", p.GetAge(), "Sal:", p.GetSal())

	fmt.Println()
	account := account.NewAccount("8888888", "123456", 100)
	if account != nil {
		fmt.Println("创建成功")
		fmt.Println(account)
	} else {
		fmt.Println("创建失败")
	}
}
