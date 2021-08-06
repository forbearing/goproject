package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("Open file error", err)
		os.Exit(1)
	}

	// 创建一个 *Reader，是带缓存的，默认缓存大小为 4096 字节
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到换行就结束一次
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}

	// 当函数退出时，要及时关闭 file
	defer file.Close()
}

/*
读取文件内容并显示在终端（带缓存的读取方式）使用
os.Open, file.Close, bufio.NewReader(), reader.ReadString 函数和方法
*/
