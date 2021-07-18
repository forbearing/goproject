package utils

import "fmt"

var Age int
var Name string

func init() {
	Age = 20
	Name = "hybfkuf"
}

func Cal(num1 float64, num2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Println("操作符错误...")
	}
	return res
}
