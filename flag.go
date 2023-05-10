package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your wonderful name")
var age = flag.Int("age", 18, "Your Age")

// First argument is a flag variable
// Second argument is a default variable
// Third argument is a description of the flag

func main() {
	flag.Parse()
	log.Printf("Hello %s, Welcome to the command line world\n", *name)
	log.Printf("Your age is %d", *age)
}

// How to run ?
//
//	./flag_example -name Habibi
//	or ...
//	./flag_example -name=Habibi
