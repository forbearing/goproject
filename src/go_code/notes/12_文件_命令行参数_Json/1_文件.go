package main

func main() {

}

/*
1. 输入流和输出流
	流：数据在数据源(文件)和程序(内存)之间经历的路径
	输入流: 数据从数据源（文件）到程序（内存）的路径
	输出流：数据从程序（内存）到数据源（文件）的路径
2. os.File 封装了所有文件相关操作，File 是一个结构体
*/

/*
os.Open(filePath)
os.OpenFile(filePath)
ioutil.ReadFile(filePath)
ioutil.WriteFile(filePath)
bufio.NewReader(file)
bufio.NewWriter(file)
reader.ReadString('\n')
writer.WriteString(str)
*/