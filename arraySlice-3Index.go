package main

import "fmt"

func main() {
	/*
		3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
		Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang,
		seperti fruits[0:1:1] . Angka kapasitas yang diisikan tidak boleh melebihi
		kapasitas slice yang akan di slicing.
	*/
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
