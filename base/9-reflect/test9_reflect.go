package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty：pair<type:*os.File,value:"./test.txt">
	tty, err := os.OpenFile("./test.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}

	var r io.Reader
	// r：pair<type:*os.File,value:"./test.txt">
	r = tty

	var w io.Writer
	// r：pair<type:*os.File,value:"./test.txt">
	w = tty
	w.Write([]byte("hello world"))
	r.Read(make([]byte, 1))

}
