package main

import "fmt"

func main() {
	map_traverse()
}

func notes() {
	/*
		map 的遍历不能用 for 循环，只能用 for-range
	*/
}

func map_traverse() {
	cites := make(map[string]string)
	cites["no1"] = "北京"
	cites["no2"] = "上海"
	cites["no3"] = "深圳"
	for k, v := range cites {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "hybfkuf"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["addr"] = "北京"
	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "jonas"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["addr"] = "上海"
	studentMap["stu03"] = make(map[string]string)
	studentMap["stu03"]["name"] = "hyb"
	studentMap["stu03"]["sex"] = "男"
	studentMap["stu03"]["addr"] = "杭州"
	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2=%v v2=%v\n", k2, v2)
		}
	}
}
