package main

import (
	"fmt"
)

// Daftar nama dengan umur nya

func main() {
	ages := map[string]int{
		"Habibi": 19,
		"Aulia":  21,
		"Yatna":  20,
		"Ilham":  19,
		"Putri":  18,
		"Reva":   19,
		"Nauval": 20,
		"Doni":   22,
	}

	var name string
	fmt.Print("Tuliskan nama : ")
	fmt.Scanln(&name)

	if _, ok := ages[name]; ok {
		fmt.Println(ages[name])
	} else {
		fmt.Println("Gak ada")
	}

	return
}
