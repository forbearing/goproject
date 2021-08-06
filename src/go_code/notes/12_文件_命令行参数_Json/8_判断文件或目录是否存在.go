package main

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // 文件或目录存在
	}
	if os.IsNotExist(err) {
		return false, nil // 文件或目录不存在
	}
	return false, err // 出错
}

func main() {

}

/*
 golang 判断文件或文件夹是否存在的方法为使用 os.Stat() 函数返回的错误值进行判断
	1. 如果返回的错误为 nil，说明文件或文件夹存在
	2. 如果返回的错误类型为 os.IsNotExist() 判断为 true，说明文件或文件夹不存在
	3. 如果返回的错误为其它类型，则不确定是否存在
*/
