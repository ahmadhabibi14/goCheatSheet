package cheatsheet

import (
	"fmt"
	"testing"
)

type Item struct {
	// Define the fields of your struct
	ID   int
	Name string
	// Add more fields as needed
}

func removeItem(arr []Item, index int) []Item {
	// Create a new slice with length reduced by 1
	result := make([]Item, len(arr)-1)

	// Copy elements before the index
	copy(result, arr[:index])

	// Copy elements after the index
	copy(result[index:], arr[index+1:])

	return result
}

func TestArrayRemoveElementV2(t *testing.T) {
	arr := []Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
	}
	index := 1

	newArr := removeItem(arr, index)

	fmt.Println(newArr) // Output: [{1 Item 1} {3 Item 3}]
}
