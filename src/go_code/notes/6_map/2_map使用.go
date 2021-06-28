package main

import "fmt"

func main() {
	// map_usage()
	example()
}

func notes() {

}

func map_usage() {
	// 方式一：
	var a map[string]string
	fmt.Println("a =", a)

	// 方式二：使用前 make
	fmt.Println()
	b := make(map[string]string, 10)
	b["name"] = "hybfkuf"
	b["age"] = "18"
	fmt.Println("b =", b)
	c := make(map[string]string) // 不指定 size
	c["no1"] = "北京"
	c["no2"] = "上海"
	c["no3"] = "杭州"
	c["no4"] = "深圳"
	fmt.Println("c =", c)

	// 方式三：
	fmt.Println()
	// var heros map[string]string = map[string]string{}
	heros := map[string]string{
		"hero1": "曹操",
		"hero2": "赵云",
		"hero3": "吕布",
	}
	fmt.Println("heros =", heros)
}

func example() {
	// 存放三个学生信息，每个学生有名字和性别
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

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"]["name"])

}
