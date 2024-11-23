package oop

import (
	"fmt"
	"testing"
)

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

func TestAbstraction(t *testing.T) {
	dog := Dog{}
	cat := Cat{}

	dog.Speak()
	cat.Speak()
}
