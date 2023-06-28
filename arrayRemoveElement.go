package main

import "fmt"

var Arr = []int{
	18, 10, 73, 44, 28, 35, 60,
}

func removeElement(arr []int, index int) []int {
	result := make([]int, len(arr)-1)

	copy(result, arr[:index])           // semua elemen hingga sebelum indeks ke-index
	copy(result[index:], arr[index+1:]) // semua elemen mulai indeks ke-index+1

	return result
}

func main() {
	indexToDelete := 2 // akan menghapus nilai 73

	newArr := removeElement(Arr, indexToDelete) // buat array baru
	Arr = newArr
	fmt.Println(Arr)
}
