package main

import "fmt"

func main() {
	map_slice()
}

func notes() {
	/*
		切片的数据类型如果是 map，则我们称为 slice of map，map 切片，这样使用则 map 个数就可以动态变化了
	*/
}

func map_slice() {
	// 1. 声明一个 map 切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	// 2. 增加第一个
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "孙悟空"
		monsters[1]["age"] = "600"
	}
	// 我们需要使用切片的 append 函数，可以动态的增加 monster
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string)
	// 	monsters[2]["name"] = "猪八戒"
	// 	monsters[2]["age"] = "1000"
	// }
	for _, value := range monsters {
		fmt.Printf("name=%v age=%v\n", value["name"], value["age"])
	}

	fmt.Println()
	newMonster := map[string]string{
		"name": "沙悟净",
		"age":  "800",
	}
	monsters = append(monsters, newMonster)
	for _, value := range monsters {
		fmt.Printf("name=%v age=%v\n", value["name"], value["age"])
	}
}
