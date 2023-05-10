package main

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) ShowInfo() string {
	return fmt.Sprintf("The %q article is written by %s.", a.Title, a.Author)
}

/* We now add a second type called Book. It also has the String method defined.
This means it also satisfies the Stringer interface. Because of this,
we can also send it to our Print function:
*/

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) ShowInfo() string {
	return fmt.Sprintf("The %q book was written by %s and it has %d pages", b.Title, b.Author, b.Pages)
}

type Stringer interface {
	ShowInfo() string
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Shammy Shark",
	}

	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	info := b.ShowInfo()
	fmt.Println(info)

	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.ShowInfo())
}
