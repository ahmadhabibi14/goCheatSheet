/*
Definisi termudah Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
Dengan menerapkan konsep tersebut, sangat mungkin untuk membuat fungsi didalam
fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.

Closure merupakan anonymous function atau fungsi tanpa nama. Biasa dimanfaatkan untuk
membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja.
*/

package cheatsheet

import (
	"fmt"
	"testing"
)

// Closure Disimpan Sebagai Variabel
/*
Sebuah fungsi tanpa nama bisa disimpan dalam variabel. Variabel yang menyimpan closure
memiliki sifat seperti fungsi yang disimpannya. Di bawah ini adalah contoh program
sederhana untuk mencari nilai terendah dan tertinggi dari suatu array. Logika pencarian
dibungkus dalam closure yang ditampung oleh variabel getMinMax .
*/

func TestFunctionClosure(t *testing.T) {
	var getMinMax = func(n []int) (int, int) /* 2 parameter ini adalah tipe data kembalian(return) */ {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var minnn, maxxx = getMinMax(numbers)
	fmt.Printf("Data : %v\nmin : %v\nmax : %v\n", numbers, minnn, maxxx)
	/*
		Penggunaan Template String %v untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya.
	*/
}
