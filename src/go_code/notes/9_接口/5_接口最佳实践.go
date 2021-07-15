package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 功能：实现对 Hero 结构体切片的排序: sort.Sort(data interface)

// 1. 声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个 Hero 结构体切片类型
type HeroSlice []Hero

// 3.实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	// 按 Hero 的年龄从小到大进行排序
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i] // 等同于上面的交换
}

func main() {
	// 先定义一个数组/切片
	intSlice := []int{0, -1, 10, 7, 90}
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	// 查看排序前的顺序
	fmt.Println("排序前的顺序")
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println()
	sort.Sort(heroes)
	// 查看排序后的顺序
	fmt.Println("排序后的顺序")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
