package main

import "fmt"

// 定一个一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	map_details()
}

func notes() {
	/*
		***** map 使用细节 *****
		1. map 是引用类型，遵守引用类型传递的机制，在一个函数接收 map，修改后，会直接修改原来的 map
		2. map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动态的增长键值对
		3. map 的 value 也经常使用 struct 类型，更适合管理复杂的数据（比前面 value 是一个 map 更好）
	*/
}

func modify_map(map1 map[int]int) {
	map1[10] = 999
}
func map_details() {
	// 1. map 是引用类型
	map1 := make(map[int]int)
	map1[1] = 10
	map1[2] = 13
	map1[10] = 35
	map1[20] = 19
	fmt.Println("map1 =", map1)
	modify_map(map1)
	fmt.Println("map1 =", map1)

	// 2. map 是动态增长的，自动扩容
	fmt.Println()
	map2 := make(map[int]int, 2)
	map2[0] = 13
	map2[8] = 88
	fmt.Println("map2 =", map2)

	// 3. map 的 value 经常使用结构体，更适合管理复杂的数据
	fmt.Println()
	students := make(map[string]Stu, 10)
	stu1 := Stu{"hybfkuf", 18, "上海"}
	stu2 := Stu{"Jonas", 20, "北京"}
	students["no1"] = stu1
	students["no2"] = stu2
	for k, v := range students {
		fmt.Printf("学生的编号: %v\n", k)
		fmt.Printf("学生的名字: %v\n", v.Name)
		fmt.Printf("学生的年龄: %v\n", v.Age)
		fmt.Printf("学生的地址: %v\n", v.Address)
		fmt.Println()
	}
}
