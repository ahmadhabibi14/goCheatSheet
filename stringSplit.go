package main

import (
	"fmt"
	"strings"
)

func main() {
	var nama string = "Ahmad Rizky Nusantara Habibi"
	var arrayNama = strings.Split(nama, " ")

	for i := 0; i < len(arrayNama); i++ {
		fmt.Printf("%d. %s\n", i+1, arrayNama[i])
	}
	// fmt.Println(arrayNama)
}
