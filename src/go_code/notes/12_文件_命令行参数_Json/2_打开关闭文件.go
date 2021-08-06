package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("Open file err: ", err)
	} else {
		fmt.Println(file)
		err = file.Close()
		if err != nil {
			fmt.Println("Close file err: ", err)
		}
	}
}

/*
1. os.Open()
	Open 打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有
	O_RDONLY 模式。如果出错，错误底层类型是 *PathError
2. func (f *File) Close() error
	Close 关闭文件 f，使文件不能用于读写，它返回可能出现的错误
*/
