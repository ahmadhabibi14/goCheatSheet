package main

import "fmt"

type Animal interface {
	Speak()
}

type (
	Dog struct{}
	Cat struct{}
)

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func (c Cat) Speak() {
	fmt.Println("Meowww")
}

func main() {
	dog := Dog{}
	cat := Cat{}

	dog.Speak()
	cat.Speak()
}
