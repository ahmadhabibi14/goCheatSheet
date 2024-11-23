package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArrayForLoop(t *testing.T) {
	var fruits = [4]string{
		"Apple", "Grape", "Banana", "Melon",
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Elemen %d : %s\n", i, fruits[i])
	}
	/*
		Perulangan di atas dijalankan sebanyak jumlah elemen array fruits
		(bisa diketahui darikondisi i < len(fruits ). Di tiap perulangan,
		elemen array diakses dengan memanfaatkan variabel iterasi i .
	*/
}
