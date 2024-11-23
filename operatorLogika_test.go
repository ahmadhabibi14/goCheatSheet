package cheatsheet

import (
	"fmt"
	"testing"
)

func TestOperatorLogika(t *testing.T) {
	var left = false
	var right = true
	
	var leftAndRight = left && right
	fmt.Printf("Left && Right \t(%t) \n", leftAndRight)
	
	var leftOrRight = left || right
	fmt.Printf("Left || Right \t(%t) \n", leftOrRight)
	
	var leftReserve = !left
	fmt.Printf("!Left \t\t(%t) \n", leftReserve)
}