package my_utils

import (
	"fmt"
	"time"
)

func main() {
	My_time()
}

func My_time() {
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
}
