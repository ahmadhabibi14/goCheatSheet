// Immediately-Invoked Function Expression (IIFE)
/*
Closure jenis ini dieksekusi langsung pada saat deklarasinya. Biasa digunakan untuk
membungkus proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak.

Di bawah ini merupakan contoh sederhana penerapan metode IIFE untuk filtering data array.
*/

package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	// var filterArray = newNumbers(3)

	fmt.Println("Original Numbers : ", numbers)
	fmt.Println("Filtered Numbers : ", newNumbers)
}

/*
Ciri khas IIFE adalah adanya kurung parameter tepat setelah deklarasi closure berakhir.
Jika ada parameter, bisa juga dituliskan dalam kurung parameternya
*/

/*
Closure bisa juga dengan gaya manifest typing, caranya dengan menuliskan skema
closure-nya sebagai tipe data. Contoh:

var closure (func (string, int, []string) int)
closure = func (a string, b int, c []string) int {
	// ..
}
*/
