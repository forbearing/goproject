package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 使用 ioutil.ReadFile 一次性将文件读取到位
	// 因为没有显式的 Open 文件，因此也不需要显式的 Close，因为文件的 Open 和 close 被封装到 ReadFile 函数内部
	file := "/tmp/test.txt"
	context, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Read file err: %v", err)
		os.Exit(1)
	}
	// 把读取到的内容显示到终端
	fmt.Printf("%v", string(context)) // context 为 []byte 类型，需要转换成 string 类型再输出
}

/*
读取文件的内容并显示在终端（使用 ioutils 一次性将整个文件读入到内存中），这种方式适用于文件不大的情况
相关方法和函数 ioutil.ReadFile
*/
