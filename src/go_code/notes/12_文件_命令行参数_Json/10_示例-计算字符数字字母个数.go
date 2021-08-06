package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 记录英文的个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main() {
	// 统计一个文件中含有的英文、数字、空格及其它字符数量
	// 思路：打开一个文件，创建一个 Reader
	// 每读取一行，就去统计该行有多少个英文、数字、空格和其它字符
	// 然后将结果保存到一个结构体

	filePath := "/tmp/1.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Open File Error: %v\n", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)

	// 开始循环读取 filePath 文件中的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数=%v 数字的个数=%v 空格的个数=%v 其它=%v\n", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
