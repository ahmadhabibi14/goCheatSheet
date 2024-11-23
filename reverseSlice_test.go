package cheatsheet

import (
	"fmt"
	"testing"
)

func reverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func TestReverseSlice(t *testing.T) {
	originalSlice := []int{1, 85, 20, 19, 5, 12, 39, 31, 72, 3, 8}
	fmt.Println("Original Slice:", originalSlice)
	reversedSlice := reverseSlice(originalSlice)

	fmt.Println("Reversed Slice:", reversedSlice)
}
