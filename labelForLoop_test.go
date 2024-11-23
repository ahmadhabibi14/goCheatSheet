package cheatsheet

import (
	"fmt"
	"testing"
)

func TestLabelForLoop(t *testing.T) {
	/*
		Di perulangan bersarang, break dan continue akan berlaku pada
		blok perulangan dimana ia digunakan saja. Ada cara agar
		kedua keyword ini bisa tertuju pada perulangan terluar
		atau perulangan tertentu, yaitu dengan memanfaatkan teknik pemberian label.
	*/
	// outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// if i == 3 {
			// 	break outerLoop
			// }
			/*
				Tepat sebelum keyword for terluar, terdapat baris kode outerLoop: .
				Maksud dari kode itu adalah disiapkan sebuah label bernama
				outerLoop untuk for dibawahnya. Nama label bisa diganti dengan
				nama lain (dan harus diakhiri dengan tanda titik dua atau colon ( : ) ).
			*/
			fmt.Print("[", i, "][", j, "] ")
		}
		fmt.Println()
	}
}
