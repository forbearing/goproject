package main

import "fmt"

// 嵌套匿名结构体
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

// 匿名字段是基本数据类型的使用
type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int
	n int
}

func main() {
	//嵌套你名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值。
	tv1 := TV{Goods{"电视机02", 4000}, Brand{"海尔", "北京"}}
	fmt.Println(tv1)
	tv2 := TV{
		Goods{
			Name:  "电视机001",
			Price: 5000,
		},
		Brand{
			Name:    "海尔",
			Address: "山东青岛",
		},
	}
	fmt.Println(tv2)
	fmt.Println()

	tv3 := TV2{&Goods{"电视机02", 4000}, &Brand{"海尔", "北京"}}
	fmt.Println(tv3)
	tv4 := TV2{
		&Goods{
			Name:  "电视机001",
			Price: 5000,
		},
		&Brand{
			Name:    "海尔",
			Address: "山东青岛",
		},
	}
	fmt.Println(tv4)
	fmt.Println()

	// 匿名字段是基本数据类型的使用
	// 如果一个结构体有 int 乐行的匿名字段，就不能有第二个
	// 如果需要多个 int 的字段，则必须给 int 字段指定名字
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println(e)
	fmt.Println()
}

func notes() {
	/*
		1. 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或小写的字段、方法、都可以使用
		2. 匿名结构体字段访问可以简化，例如
			var b B
			b.A.name = "myname" 	==> b.name = "myname"
			b.A.say() 				==> b.say()
		3. 当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则访问。如果希望访问匿名结构体
		   的字段和方法，可以通过匿名结构体名来区分。
		4. 编译器会先看 b 对应的类型有没有 Name，如果有则直接调用 B 类型的 Name 字段，如果没有则就去看
		   B 中嵌入的匿名结构体 A 有没有声明 Name 字段，如果有就调用，如果没有继续查找。
		   如果都找不到则报错
		5. 结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法）
		   在访问时，就必须明确指定匿名结构体名字，否则编译报错
		   type A struct {
			   Name string
			   Age int
		   }
		   type B struct {
			   Name string
			   Score int
		   }
		   type C struct {
			   A
			   B
		   }
		   C.Name 报错，C.A.Name 才行
		6. 如果一个 struct 嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体
		   的字段或方法时，必须带上结构体的名字
		   type A struct {
			   Name string
			   Age int
		   }
		   type C struct {
			   a A 		// 有名结构体
		   }
		   d.Name					错误，不会自动查找
		   d.a.Name = "jack"		正确

		7.嵌套你名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值。
	*/
}
