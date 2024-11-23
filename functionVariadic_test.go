package cheatsheet

import (
	"fmt"
	"strings"
	"testing"
)

/*
	Golang mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter
	sejenis yang tak terbatas. Maksud tak terbatas disini adalah jumlah parameter yang
	disisipkan ketika pemanggilan fungsi bisa berapa saja.

	Parameter variadic memiliki sifat yang mirip dengan slice. Nilai parameter-parameter yang
	disisipkan memiliki tipe data yang sama, dan akan ditampung oleh sebuah variabel saja.
	Cara pengaksesan tiap datanya juga sama, dengan menggunakan indeks.
*/

func TestFunctionVariadic(t *testing.T) {
	var avg = calculateX(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata	: %.2f", avg)
	/*
		Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf() , hanya saja fungsi ini
		tidak menampilkan nilai, melainkan mengembalikan nilainya dalam bentuk string. Pada
		kasus di atas, nilai hasil fmt.Sprintf() ditampung oleh variabel msg .
	*/
	fmt.Println(msg)
	// Jika parameter kedua dan seterusnya ingin diisi dengan data dari slice, maka gunakan tanda 3 titik.
	var hobbies = []string{
		"Coding",
		"Playing Games",
		"Football",
		"Watch Movies",
	}
	yourHobbies("Habibi", hobbies...)
	// Jika tidak, maka begini
	// yourHobbies("Habibi", "Sleping", "Eating")
}

func calculateX(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	// Rumus rata rata,
	return avg
}

// Fungsi Dengan Parameter Biasa & Variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	// Kode ini untuk menaruh koma di setiap akhir elemen slice hobbies[]

	fmt.Printf("Hello, My name is : %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)

	/*
		Nilai parameter pertama fungsi yourHobbies() akan di tampung
		oleh name, sedangkan nilai parameter kedua dan seterusnya akan di tampung
		oleh hobbies sebagai slice.

	*/
}
