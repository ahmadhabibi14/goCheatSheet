package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArrayMoveElementToLast(t *testing.T) {
	children := []int{198, 349, 825, 902, 627}
	fmt.Println("Children Before:", children)
	
	moveChildToLast(&children, 902)
	
	fmt.Println("After moving:", children)
}

func moveChildToLast(children *[]int, elm int) {
	elmIdx := 0

	for i, v := range *children {
		if v == elm {
			elmIdx = i
		}
	}

	element := (*children)[elmIdx]
	*children = append((*children)[:elmIdx], (*children)[elmIdx+1:]...)

	*children = append(*children, element)
}
