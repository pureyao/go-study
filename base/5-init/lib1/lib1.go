package lib1

import (
	"fmt"
	"go-study/base/5-init/lib2"
)

func Lib1Test() {
	fmt.Println("Lib1 Test")
	lib2.Lib2Test()
}

func init() {
	fmt.Println("Lib1 init")
}
