package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 将一个图片拷贝到另外一目录下

// 自己编写一个函数，接收两个文件路径，srcFileName dstFileName
func copyFile(dstFileName string, srcFileName string) (writer int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("Open File Error: %v\n", err)
	}
	defer srcFile.Close()
	// 通过 sfile 获得 Reader
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		fmt.Printf("Open File Error: %v\n", err)
	}
	defer dstFile.Close()
	// 通过 dfile 获得 Writer
	writer2 := bufio.NewWriter(dstFile)

	return io.Copy(writer2, reader)
}

func main() {
	srcFile := "/Users/hybfkuf/Downloads/1.png"
	dstFile := "/tmp/1.png"
	_, err := copyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 err= %v\n", err)
	}
}
