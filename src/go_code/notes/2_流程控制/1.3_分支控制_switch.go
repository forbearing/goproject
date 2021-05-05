package main

import "fmt"

func main() {
	// notes()
	// my_func1()
	my_func2()
}

func notes() {
	/*
		1. switch 语句用于基于不同条件执行不同动作，每一个 case 动作都是唯一的，从上到下逐一测试，直到匹配为止
		2. 匹配项后面也不需要再加 break
		3. switch 的执行的流程是，先执行表达式，得到值，然后和 case 的表达式进行比较，如果相等，就匹配到，
		   然后执行对应的 case 的语句块，然后退出 switch 语句。
		4. 如果 switch 的表达式的值没有和任何的 case 的表达式匹配成功，则执行 default 语句块。执行后退出 switch 的控制
		5. golang 的 case 后的表达式可以有多个，使用逗号间隔
		6. golang 中的 case 语句块不需要写 break，因为默认会有，即，当程序执行完 case 语句块后，就直接退出 switch 控制结构
	*/

	/*
		*** 细节讨论 ***
		1. case 后是一个表达式（即：常量值、变量、一个有返回值的函数等都可以）
		2. case 后的各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致
		3. case 后面可以带多个表达式，使用逗号间隔。比如 case 表达式1，表达式2 ...
		4. case 后面的表达式如果是常量值（字面量）则要求不能重复。
		5. case 后面不需要带 break，程序匹配到一个 case 后就会执行对应的代码块，然后退出 switch，
		   如果一个都匹配不到，则执行 default。
		6. default 语句不是必须的。
		7. switch 后也可以不带表达式，类似 if-else 分支来使用
		8. switch 后也可以直接声明/定义一个变量，分号结束，不推荐
		9. switch 穿透 fallthrough，如果在 case 语句块后增加 fallthrough，则会继续执行下一个 case，
		   也叫 switch 穿透。
		10. Type switch: switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的变量类型
	*/

	/*
		**** switch 和 if 的比较 ****
		1. 如果判断的具体数值不多，而且符合整数，浮点数，字符，字符串这几种类型。建议使用 switch 语句，简介高效
		2. 其它情况，对区间判断和结果为 bool 类型的判断，使用if，if的适用范围更广泛。
	*/
}

func my_func1() {

	var key byte
	fmt.Println("请输入一个字符: a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期天")
	default:
		fmt.Println("输入的不是我指定的范围")
	}

	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 5, 10:
		fmt.Println("ok")
	default:
		fmt.Println("没有匹配到")
	}

	// switch 后不带表达式，类似 if-else 分支来使用。
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配")
	}

	// case 也可以对范围做判断
	var score int = 90
	switch {
	case score >= 90:
		fmt.Println("成绩优秀..")
	case score >= 70 && score < 90:
		fmt.Println("成绩优良..")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格..")
	default:
		fmt.Println("你的成绩太烂了...")
	}

	// switch 后也可以直接声明一个变量，分号结束，不推荐
	switch grade := 90; {
	case grade >= 90:
		fmt.Println("优秀")
	case grade >= 70 && grade < 90:
		fmt.Println("良")
	case grade >= 60 && grade < 70:
		fmt.Println("及格")
	default:
		fmt.Println("都没及格")
	}

	// switch 穿透 fallthrough，默认只能穿透一层
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配成功")
	}
}
func my_func2() {
	//使用 switch 把小写类型的 char 型转换为大写（键盘输入），只转换 a，b，c，d，e，其它的输出 “other”
	var char byte
	fmt.Println("请输入一个字符:")
	fmt.Scanf("%c", &char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("others")
	}

	var score float64
	fmt.Println("请输入你的成绩")
	fmt.Scanf("%f", &score)
	switch int(score / 60) {
	case 1:
		fmt.Println("及格")
	case 0:
		fmt.Println("不及格")
	}
}
