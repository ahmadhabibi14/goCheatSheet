package main

import (
	"fmt"
)

func main() {
	// Daftar nama dengan umur nya
	ages := map[string]func(){
		"habibi": umurNya(19),
	}

	var name string
	fmt.Print("Tuliskan nama : ")
	fmt.Scanln(&name)

	if comFunc, ok := ages[name]; ok {
		comFunc()
	} else {
		fmt.Println("Gak ada")
	}

	return
}

func umurNya(umur int) func() {
	return func() {
		fmt.Printf("Umurnya adalah : %d", umur)
	}
}
