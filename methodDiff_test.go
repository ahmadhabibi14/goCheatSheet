package main

import "fmt"

type progLang struct {
	compiled    string
	interpreted string
}

type library struct {
	frontend string
	backend  string
}

// Uses method pointer
func (pl *progLang) bestProgLang() {
	fmt.Println("Best language compiled :: ", pl.compiled)
	fmt.Println("Best language interpreted :: ", pl.interpreted)
}

// Not uses method pointer
func (lb library) bestLibrary() {
	fmt.Println("Best library frontend :: ", lb.frontend)
	fmt.Println("Best library backend :: ", lb.backend)
}

func main() {
	var S1 = &progLang{
		"Golang", "JavaScript",
	}
	var S2 = library{
		"Reactjs", "Gin-gonic",
	}
	S1.bestProgLang()
	S2.bestLibrary()
}
