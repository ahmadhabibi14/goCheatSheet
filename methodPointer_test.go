package cheatsheet

import (
	"fmt"
	"testing"
	"time"
)

/*
	Method pointer adalah method yang variabel objeknya dideklarasikan dalam bentuk pointer.
	Kelebihan method jenis ini adalah manipulasi data pointer pada property milik variabel
	tersebut bisa dilakukan.
*/

type studentX struct {
	name string
	age  int
}

func (s *studentX) sayHello() {
	fmt.Println("Halo", s.name)
	fmt.Println("Your age is", s.age)
}

func TestMethodPointer(t *testing.T) {
	var now_year = time.Now().Year()
	// var year = strconv.Atoi(now_year)
	var my_age int = now_year - 2004
	var s1 = &studentX{
		"Ahmad Habibi",
		my_age,
	}
	s1.sayHello()
}
