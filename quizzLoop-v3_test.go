package cheatsheet

import (
	"fmt"
	"testing"
)

func toPrint(total int, charc string) {
	for i := 0; i < total; i++ {
		fmt.Print(charc)
		if i != (total) {
			fmt.Print(" ")
		}
	}
}

func TestQuizLoopV3(t *testing.T) {
	var (
		x int = 3
		y int = 5
	)

	for i := 0; i < y; i++ {
		toPrint(x, "O")
		toPrint(x, "*")
		toPrint(x, "=")
		fmt.Println("")
	}
}

/*
	OUTPUT :
	O	O	O	*	*	*	=	=	=
	O	O	O	*	*	*	=	=	=
	O	O	O	*	*	*	=	=	=
	O	O	O	*	*	*	=	=	=
	O	O	O	*	*	*	=	=	=
*/
