package cheatsheet

import (
	"fmt"
	"testing"
)

func TestQuizLoopV2(t *testing.T) {
	/*
		OUTPUT :
		O	O	O	O	O	O
		 O	 O  O	 O  O
		  O  O  O  O
			O	O	O
			 O  O
			  O
	*/
	fmt.Println("Tugas Nomor 2 :")
	var angka int
	fmt.Print("Masukkan angka : ")
	// Ini untuk menginput angka dari user
	fmt.Scanln(&angka)
	// Menggunakan pointer
	for i := 0; i < angka; i++ {
		for j := 0; j < angka; j++ {
			if (i%2 == 1) && (j%2 == 1) {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}
