package main

func main() {

}

func notes() {
	/*
		1、调用方式不一样
			函数的调用方式：函数名(实参列表)
			方法的调用方式：变量.方法名(实参列表)
		2、对于普通函数，接受者为值类型，不能将指针类型的数据直接传递，反之毅然
		3、对于方法（如 struct 的方法）接受者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
			不管调用方式如何，真正决定是值拷贝还是地址拷贝，看这方法是和哪个类型绑定的
			如果是值类型，比如 (p Person) 则是值拷贝，如果是和指针类型，比如 (p *Person) 则是地址拷贝
	*/
}
