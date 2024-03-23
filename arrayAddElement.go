package main

import "fmt"

func insertElementAtIndex(arr []int, index, element int) []int {
    // Ensure that the index is within the bounds of the array.
    if index < 0 || index >= len(arr) {
      arr = append(arr, element)
			return arr
    }

    // Expand the slice by one element.
    arr = append(arr, 0)

    // Shift elements to the right to make space for the new element.
    copy(arr[index+1:], arr[index:])

    // Insert the new element at the specified index.
    arr[index] = element

    return arr
}

func main() {
    // Example usage
    array := []int{1, 2, 3, 4, 5}
    index := 5
    element := 10

    newArray := insertElementAtIndex(array, index, element)
    fmt.Println(newArray) // Output: [1 2 10 3 4 5]
}
