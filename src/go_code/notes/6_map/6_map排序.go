package main

import (
	"fmt"
	"sort"
)

func main() {
	map_sort()
}

func notes() {
	/*
		***** map 排序 *****
			1. golang 中没有一个专门的方法针对 map 的 key 进行排序
			2. golang 的 map 默认是无序的，也不是按添加顺序存放的，你每次遍历得到的输出都不一样。
			3. golang 中的 map 的排序，是先将 key 进行排序，然后根据 key 值遍历输出即可
	*/
}

func map_sort() {
	map1 := make(map[int]int)
	map1[0] = 11
	map1[2] = 13
	map1[4] = 35
	map1[8] = 10
	fmt.Println("map1 =", map1)

	// 声明一个切片用来存放 map 的 key
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	// 排序 map 的 key，然后根据 key 值遍历输出即可。
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
	}
}
