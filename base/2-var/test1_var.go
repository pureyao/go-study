package main

import "fmt"

var (
	xx int    = 123
	yy string = "345"
)

func main() {
	a := 100
	fmt.Println("a:", a)
	fmt.Printf("a的数据类型:%T\n", a)

	fmt.Println("xx:", xx)
	fmt.Printf("xx的数据类型:%T\n", xx)
	fmt.Println("yy:", yy)
	fmt.Printf("yy的数据类型:%T\n", yy)
}
