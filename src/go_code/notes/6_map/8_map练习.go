package main

import "fmt"

func main() {
	// 1. 使用 map[string]map[string]string 的 map 类型
	// 2. key 表示用户名，是唯一的，不可以重复
	// 3. 如果某个用户存在，就将其密码修改为 "888888", 如果不存在就增加这个用户信息（包括昵称 nickname 和密码 passwd)
	// 4. 编写一个函数 modifyUser(user map[string]map[string]string, name string)

	users := make(map[string]map[string]string, 10)
	users["tom"] = make(map[string]string)
	users["tom"]["passwd"] = "123456"
	users["tom"]["nickname"] = "hello"
	modifyUsers(users, "hybfkuf")
	modifyUsers(users, "Jonas")
	modifyUsers(users, "tom")
	fmt.Println("users =", users)
}

func modifyUsers(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["passwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["passwd"] = "888888"
		users[name]["nickname"] = "昵称" + name
	}
}
