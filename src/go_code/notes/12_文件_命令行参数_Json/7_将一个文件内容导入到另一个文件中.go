package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 将一个文件内容导入到另一个文件中
	// 使用 ioutil.ReadFile / ioutil.WriteFile 完成

	// 1. 先将文件 file1.txt 读取到内存
	// 2. 将读取到到内容写入到 file2.txt

	file1Path := "/tmp/file1.txt"
	file2Path := "/tmp/file2.txt"
	context, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("Read File Error: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(file2Path, context, 0640)
	if err != nil {
		fmt.Printf("Write File Error: %v\n", err)
		os.Exit(1)
	}
}
