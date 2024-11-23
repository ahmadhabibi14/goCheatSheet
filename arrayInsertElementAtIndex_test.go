package cheatsheet

import (
	"fmt"
	"log"
	"testing"
)

func TestArrayInsertElm(t *testing.T) {
  var children = []any{}
	fmt.Println("Children Before:", children)

	toIdx := -100
	children, err := insertElementAtIndex(children, 59, toIdx)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("After insert to index %d: %v\n", toIdx, children)
}

func insertElementAtIndex(slice []any, element any, toIndex int) ([]any, error) {
	if len(slice) < 1 || toIndex < 0{
		slice = append(slice, element)
		return slice, nil
	}

	slice = append(slice, 0)

	copy(slice[toIndex+1:], slice[toIndex:])

	slice[toIndex] = element

	return slice, nil
}