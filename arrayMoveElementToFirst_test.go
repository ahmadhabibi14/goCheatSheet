package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArrayMoveElemenToFirst(t *testing.T) {
	children := []int{198, 349, 825, 902, 627}
	fmt.Println("Children Before:", children)
	
	moveChildToFirst(&children, 902)
	
	fmt.Println("After moving:", children)
}

func moveChildToFirst(children *[]int, elm int) {
	elmIdx := 0

	for i, v := range *children {
		if v == elm {
			elmIdx = i
		}
	}

	temp := (*children)[elmIdx]

	for i := elmIdx; i > 0; i-- {
		(*children)[i] = (*children)[i-1]
	}

	(*children)[0] = temp

	return
}