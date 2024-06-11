package __funciton

import "fmt"

// 匿名返回
func foo1(a string) (int, int) {
	fmt.Println(a)
	return 12, 13
}

// 有形参返回
func foo2(a string) (r1 int, r2 int) {
	fmt.Println(a)
	r1 = 12
	r2 = 13
	return
}

// 多形参合并
func foo3(a string) (r1, r2 int) {
	fmt.Println(a)
	r1 = 12
	r2 = 13
	return
}
