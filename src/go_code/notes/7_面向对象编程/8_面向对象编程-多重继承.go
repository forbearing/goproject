package main

// 多重继承
type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}

func main() {

}

func notes() {
	/*
		多重继承说明
			如果一个 struct 嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，
			从而实现了多重继承

		多重继承细节说明
			1:如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分
			2:为了保证代码的简洁性，建议尽量不使用多重继承
	*/
}
