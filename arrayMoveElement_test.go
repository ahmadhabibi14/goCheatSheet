package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// Sample slice
	slice := []any{1, 2, 3, 4, 5}
	
	// Move element 3 to index 1
	toSlice, err := moveElement(slice, 2, 200)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(toSlice) // Output: [1 3 2 4 5]
}

// moveElement moves the element at the given value to the desired index in the slice
func moveElement(slice []any, element any, newIndex int) ([]any, error) {
	// Find the index of the element
	var currentIndex int = -1
	for i, v := range slice {
		if v == element {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return []any{}, errors.New("element not found")
	}

	if newIndex < 0 {
		newIndex = 0
	}
	
	// Remove the element from the current index
	slice = append(slice[:currentIndex], slice[currentIndex+1:]...)
	log.Println(slice)
	
	// Insert the element at the desired index
	if newIndex >= len(slice) {
		// If the new index is beyond the length of the slice, append the element
		slice = append(slice, element)
	} else {
		// Otherwise, insert the element at the desired index
		slice = append(slice[:newIndex], append([]any{element}, slice[newIndex:]...)...)
	}

	return slice, nil
}
