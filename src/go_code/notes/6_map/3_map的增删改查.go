package main

import "fmt"

func main() {
	map_crud()
}

func notes() {
	/*
		***** map 增加和更新 ******
			map["key"] = value			// 如果 key 没有，就是增加，如果 key 存在就是修改
		***** map 的删除 *****
			- delete(map, "key")，delete 是一个内置函数，如果 key 存在，就删除 key-value
			  如果 key 不存在，不操作， 但是也不会报错
			- 如果要删除 map 的所有 key，没有一个专门的方法一次性删除，可以遍历一下 key，逐个删除
			  或者 map = make(...), make 一个新的，让原来的成为垃圾，被 gc 回收
	*/
}

func map_crud() {
	// map 增加、修改
	cites := make(map[string]string)
	cites["no1"] = "北京"
	cites["no2"] = "上海"
	cites["no3"] = "深圳"
	fmt.Println("cites =", cites)
	cites["no3"] = "~"
	fmt.Println("cites =", cites)

	// map 删除
	fmt.Println()
	delete(cites, "no4") // 删除不存在的 key-value 不会操作也不会报错
	fmt.Println("cites =", cites)
	delete(cites, "no3")
	fmt.Println("cites =", cites)
	cites = make(map[string]string) // 删除整个 map
	fmt.Println("cites =", cites)

	// map 查找
	fmt.Println()
	cites = make(map[string]string)
	cites["no1"] = "北京"
	cites["no2"] = "上海"
	cites["no3"] = "深圳"
	val, ok := cites["no1"] // 有 no1 返回 true，否则返回 false
	if ok {
		fmt.Println("value =", val)
	}

	// map 的长度
	fmt.Println()
	fmt.Printf("cites len: %v", len(cites))
}
