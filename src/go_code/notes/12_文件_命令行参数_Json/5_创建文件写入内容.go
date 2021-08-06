package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件，并写入内容
	filePath := "/tmp/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
		os.Exit(1)
	}
	// 及时关闭 file 描述符
	defer file.Close()

	str := "hello golang\n"
	// 写入时，使用带缓存的 *Write
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为 writer 是带缓存的，因此在调用 WriteString 方法时，其实内容是先写入到缓存的
	// 所以需要调用 Flush 方法，将缓存的数据真正写入到文件中。否则文件中会没有数据。
	writer.Flush()

}

/*
func OpenFile(name string, flat int, perm FileMode) (file *File, err error)
说明：os.OpenFile 是一个更一般性的文件打开函数，它会使用指定的选项（如 O_RDONLY 等）、指定的模式（如0666等）
打开指定名称的文件。如果操作成功，返回的文件对象可用于 I/O。如果出错，错误底层类型是 *PathError
flag
	O_RDONLY		// 只读模式打开文件
	O_WRONLy		// 只写模式打开文件
	O_RDWR			// 读写模式打开文件
	O_APPEND		// 写操作时将文件附加到文件末尾
	O_CREATE		// 如果不存在将创建一个文件
	O_EXEC			// 和 O_CREATE 配合使用，文件必须不存在
	O_SYNC			// 打开文件用于同步 I/O
	O_TRUNC			// 如果可能，打开时清空文件
perm
	权限控制（仅限于 Linux，Windows 下无效）
*/
