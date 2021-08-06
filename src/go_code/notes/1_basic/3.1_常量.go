package main

func main() {
	const num int = 10
	const tax float64 = 0.0

}

/*
1.常量介绍
	1.常量使用 const 修饰
	2.常量在定义的时候，必须初始化
	3.常量不能修改
	4.常量只能修饰 bool、数据类型(int, float系列), string 类型
	5.语法: const identifier [type] = value
	6:举例
		const name = "hybfkuf" 		# 可以使用类型推导
		const tax float64 = 0.8
		const b = 9 / 3
		const c = getVal()			# error
2.常量使用注意事项
	1.比较简洁的写法
		func main() {
			const (
				a = 1
				b = 2
			)
			fmt.Println(a,b)
		}
	2.还有一种专业的写法
		func main() {
			const (
				a = iota
				b
				c
			)
			fmt.Println('a,b,c) // 0,1,2
			# iota 表示给 a 赋值为0，
		}
	3.Golang 中没有常量名必须字母大写的规范
	4.仍然通过首字母的大小写来控制常量的访问范围
*/
