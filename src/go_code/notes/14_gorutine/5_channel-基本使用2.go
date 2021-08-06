package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// ===== map channel
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "beijing"
	m1["city2"] = "shanghai"
	m2 := make(map[string]string, 20)
	m2["name1"] = "jonas"
	m2["name2"] = "hybfkuf"
	mapChan <- m1
	mapChan <- m2
	m3 := <-mapChan
	m4 := <-mapChan
	fmt.Println(m3, m4)

	// ===== struct channel
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	c1 := Cat{Name: "hybfkuf", Age: 10}
	c2 := Cat{Name: "jonas", Age: 20}
	catChan <- c1
	catChan <- c2
	c3 := <-catChan
	c4 := <-catChan
	fmt.Println(c3, c4)

	// ===== *struct channel
	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)
	cc1 := Cat{Name: "hybfkuf", Age: 10}
	cc2 := Cat{Name: "jonas", Age: 20}
	catChan2 <- &cc1
	catChan2 <- &cc2
	cc3 := <-catChan2
	cc4 := <-catChan2
	fmt.Println(cc3, cc4)

	// ===== interface{} channel
	cat := Cat{Name: "hybfkuf"}
	// var allChan chan interface{}
	// allChan = make(chan interface{}, 3)
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	allChan <- cat

	// 我们希望获得管道中的第三个元素，则需要先将前面2个推出，不能跳过
	<-allChan
	<-allChan
	newCat := <-allChan // 从管道中取出 Cat 是什么
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	a := newCat.(Cat) // 需要一个类型断言
	fmt.Printf("newCat.Name=%v\n", a.Name)

}
