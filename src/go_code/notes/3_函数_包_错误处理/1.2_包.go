package main

import (
	"fmt"
	"hybfkuf/utils"
)

func main() {
	my_func1()
}

func notes() {
	/*
		***** 包的介绍 *****
		包的本质实际上就是创建不同的文件夹，来存放程序文件。
		tree /usr/local/go/src/hybfkuf/                                                  (hybfkuf) 0 < 20:59:17
			/usr/local/go/src/hybfkuf/
			└── utils
				└── utils.go

		1. go 的每一个文件都属于一个包，也就是说 go 是以包的形式来管理文件和项目目录结构的
		2. 包的三大作用
			1. 区分相同名字的函数、变量等标识符
			2. 当程序文件很多时，可以很好的管理项目
			3. 控制函数、变量等访问范围，即作用域。
		3. 包的相关说明
			1. 打包基本语法: package [包名]
			2. 引入包的基本语法: import "包的路径"


		***** 包的注意事项和细节说明 *****
		1. 在给一个文件打包时，该包对应一个文件夹，比如这里的 utils 文件夹对应的包名就是 utils，
		   文件的包名通常和文件所在的文件夹名一致，一般为小写字母。
		2. 当一个文件要使用其它包函数或变量时，需要引入对应的包
			1. 引入方式1: import "包名"
			2. 引入方式2:
				import (
					"包名"
					"包名"
				)
			3. package 指令在文件第一行，然后是 import 指令
			4. 在 import 包时，路径从 $GOPATH 的 src 下开始，不用带 src,
		   	   编译器会自动从 src 下开始引入。
		3. 为了让其它包的文件，可以访问到本包的函数，则该函数的首字母需要大写，类似其它语言的 public，
		   这样才能跨包访问。
		4. 在访问其它包函数、变量时，其语法是 包名.函数名.
		5. 如果包名较长，Go 支持给包取别名，注意细节：取别名后，原来的包名就不能使用了
			import (
				util "hybfkuf/utils"
			)
			hybfkuf/utils 的别名是 util 了。
		6. 在同一包下，不能有相同的函数名，否则报重复定义。
		7. 如果你要编译成一个可执行程序文件，就需要将这个包声明为 main，即 package main，
		   这就是一个语法规范，如果你是写一个库，包名可以自定义。
	*/
}

func my_func1() {
	res := utils.Cal(1, 10, '+')
	fmt.Println(res)
}
