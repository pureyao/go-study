package main

import "fmt"

type Animal interface {
	Play()
	Eat()
}

type Cat struct {
	food string
}

func (cat *Cat) Play() {
	fmt.Println("玩毛线")
}

func (cat *Cat) Eat() {
	fmt.Println("吃", cat.food)
}

type Dog struct {
	food string
}

func (dog *Dog) Play() {
	fmt.Println("玩球")
}

func (dog *Dog) Eat() {
	fmt.Println("吃", dog.food)
}

func main() {
	cat := Cat{food: "小鱼干"}
	cat.Play()
	cat.Eat()

	dog := Dog{food: "骨头"}
	dog.Play()
	dog.Eat()
}
