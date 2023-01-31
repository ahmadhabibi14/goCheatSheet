package main

import (
	"errors"
	"fmt"
	"strings"
)

type user struct {
	name string
}

func (u user) getName() {
	fmt.Println("Halo", u.name)
}

func emptyString(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		// Ini akan membuat error baru
		return false, errors.New("Name cannot be empty")
	}

	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name :: ")
	fmt.Scanln(&name)

	var n = user{
		name,
	}
	if valid, err := emptyString(name); valid {
		n.getName()
	} else {
		// disini akan di tampilkan error tersebut
		fmt.Println(err.Error())
	}
}
