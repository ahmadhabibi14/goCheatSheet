package main

import (
	"fmt"
)

/*
	Method pointer adalah method yang variabel objeknya dideklarasikan dalam bentuk pointer.
	Kelebihan method jenis ini adalah manipulasi data pointer pada property milik variabel
	tersebut bisa dilakukan.
*/

type student struct {
	name  string
	grade int
}

func (s *student) sayHello() {
	fmt.Println("Halo", s.name)
}

func main() {
	var s1 = student{
		"Ahmad Habibi",
		18,
	}
	s1.sayHello()
}
