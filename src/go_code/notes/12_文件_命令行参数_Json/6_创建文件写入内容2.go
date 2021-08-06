package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/tmp/2.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0640)
	if err != nil {
		fmt.Printf("Open File Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 读取文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("NewReader Error: %v\n", err)
		}
		fmt.Print(str)
	}

	// 写文件
	str := "hello golang\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 2; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
