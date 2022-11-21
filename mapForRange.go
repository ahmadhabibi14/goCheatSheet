package main

import "fmt"

func main() {
	/*
		Item variabel map bisa di iterasi menggunakan for - range . Cara penerapannya masih
		sama seperti pada slice, pembedanya data yang dikembalikan di tiap perulangan adalah key
		dan value, bukan indeks dan elemen
	*/

	var chicken = map[string]int{
		"January":  50,
		"February": 40,
		"Maret":    34,
		"April":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "	\t:", val)
	}
}
