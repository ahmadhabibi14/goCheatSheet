package cheatsheet

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	var point = 8

	if point >= 10 {
		fmt.Println("Lulus dengan Nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("Tidak Lulus!, nilai anda %d\n", point)
	}
}