/*
Setelah di bab sebelumnya kita belajar mengenai fungsi yang mengembalikan nilai balik
berupa fungsi, kali ini topiknya tidak kalah unik, yaitu fungsi yang digunakan sebagai
parameter.

Di Golang, fungsi bisa dijadikan sebagai tipe data variabel. Dari situ sangat memungkinkan
untuk menjadikannya sebagai parameter juga.
*/

package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{
		"Wick", "Jason", "Ethan", "Bruce",
	}
	var datacontains0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var datalength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli \t\t:", data)
	// data asli : [Wick Jason Ethan Bruce]

	fmt.Println("Filter ada huruf \"i\"\t:", datacontains0)
	// filter ada huruf "i" : Wick

	fmt.Println("Filter jumlah huruf  \"5\"\t:", datalength5)
	// filter jumlah huruf "5" :[Jason Ethan]
}
