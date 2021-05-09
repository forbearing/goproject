package main

import (
	"errors"
	"fmt"
)

func main() {
	// capture_err()
	custom_error()
}

func notes() {
	/*
		当发生错误后，并进行处理，保证程序可以继续进行，还可以在捕获到错误后，给管理员一个特殊的提升
		***** 基本说明 *****
		1. Go 语言追求简洁优雅，所以，Go 语言不支持传统的 try... catch... finally 这种处理
		2. Go 中引入的处理方法为: defer, panic, recover
		3. 这几个异常的使用场景可以这么简单描述：go 中可以抛出一个 panic 的异常，然后在 defer 中
		   通过 recover 捕获这个异常，然后正常处理。
		4. 错误处理的好处: 进行错误处理后，程序不会轻易挂掉，如果加入预警代码，就可以让程序更加健壮

		***** 自定义错误 *****
		1. Go 程序中，也支持自定义错误，使用 errors.New 和 panic 内置函数
		2. panic 内置函数，接受一个 interface{} 类型的值(也就是任何值)作为参数。可以接收 error
		   类型的变量，输出错误信息，并退出程序。

	*/
}

func capture_err() {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover() 内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到异常
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		// 文件名是 config.ini，返回 nil
		return nil
	} else {
		// 文件名不是 config.ini 返回自定义错误
		return errors.New("读取文件错误...")
	}
}

func custom_error() {
	// 函数函去读取配置文件 config.ini 的信息，如果文件名传入不正确，就返回一个自定义的错误
	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("继续执行 ...")
}
