package main

import "fmt"

type Book struct {
	name string
	auth string
}

func (this *Book) Read() {
	fmt.Println("Read...")
}

type ProBook struct {
	Book
	level int
}

func (this *ProBook) Read() {
	fmt.Println("ReadPro...")
}

func (this *ProBook) Do() {
	fmt.Println("Do...")
}

func main() {
	book := Book{name: "nomal", auth: "hh"}
	book.Read()

	proBook := ProBook{Book{name: "Probook", auth: "qq"}, 12}
	proBook.Read()
	proBook.Do()
}
