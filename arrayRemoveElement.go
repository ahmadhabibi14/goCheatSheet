package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var children = []any{
		18, 10, 73, 44, 28, 35, 60,
	}
	fmt.Println("Children Before:", children)
	children, err := removeElement(children, 44)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("After removed:", children)
}


func removeElement(slice []any, element any) ([]any, error) { 
	var elmIndex int = -1
	for i, v := range slice {
		if v == element {
			elmIndex = i
			break
		}
	}

	if elmIndex == -1 {
		return []any{}, errors.New("element not found")
	}

	result := make([]any, len(slice)-1)
	
	copy(result, slice[:elmIndex]) 
	copy(result[elmIndex:], slice[elmIndex+1:])
	
	return result, nil
}