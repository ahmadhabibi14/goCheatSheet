package cheatsheet

import (
	"fmt"
	"testing"
)

func batasXC() {
	fmt.Println("+=========================+")
}

func TestQuizLoop(t *testing.T) {
	/*
		OUTPUT :
				O
				O
		O	O	O	O	O
				O
				O
	*/
	fmt.Println("Tugas Nomor 1 :")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var operasi int = 5%2 + 1
			if i == operasi {
				fmt.Print("O ")
			} else if j == operasi {
				fmt.Print("O ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	batasXC()

	/*
		OUTPUT :
		O	O	O	O	O
		O	X	O	X	O
		O	O	O	O	O
		O	X	O	X	O
		O	O	O	O	O
	*/
	fmt.Println("Tugas Nomor 2 :")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (i%2 == 1) && (j%2 == 1) {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
	batasXC()

	/*
		OUTPUT :
		O	X	O	X	O
		X	X	X	X	X
		O	X	O	X	O
		X	X	X	X	X
		O	X	O	X	O
	*/
	fmt.Println("Tugas Nomor 3 :")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (i%2 == 0) && (j%2 == 0) {
				fmt.Print("O ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
	batasXC()
}
