package main

import "fmt"

func batas() {
	fmt.Println("\n=========================== | >\n")
}

func main() {
	/*
		Fungsi len() digunakan untuk menghitung lebar slice yang ada. Sebagai contoh jika
		sebuah variabel adalah slice dengan data 4 buah, maka fungsi ini pada variabel tersebut
		akan mengembalikan angka 4, yang angka tersebut didapat dari jumlah elemen yang ada.
	*/
	var dcHeroes = []string{
		"Batman",
		"Superman",
		"Wonder Woman",
		"The Flash",
		"Aquaman",
	}
	fmt.Println(len(dcHeroes))
	batas()

	/*
		Fungsi cap() digunakan untuk menghitung lebar maksimum/kapasitas slice.
		Nilai kembalian fungsi ini awalnya sama dengan len , tapi bisa berubah tergantung
		dari operasi slice yang dilakukan. Agar lebih jelas, silakan disimak kode berikut.
	*/
	var fruits = []string{
		"Apple", "Grape", "Banana", "Melon",
	}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	/*
		Variabel fruits disiapkan di awal dengan jumlah elemen 4. Maka fungsi len(fruits) dan
		cap(fruits) akan menghasilkan angka 4.

		Variabel aFruits dan bFruits merupakan slice baru berisikan 3 buah elemen milik slice
		fruits . Variabel aFruits mengambil elemen index 0, 1, 2; sedangkan bFruits 1, 2, 3.

		Fungsi len() menghasilkan angka 3, karena jumlah elemen kedua slice ini adalah 3.
		Tetapi cap(aFruits) menghasilkan angka yang berbeda, yaitu 4 untuk aFruits dan 3 untuk bFruits

		Slicing yang dimulai dari indeks 0 hingga x akan mengembalikan elemen-elemen mulai
		indeks 0 hingga sebelum indeks x, dengan lebar kapasitas adalah sama dengan slice
		aslinya. Sedangkan slicing yang dimulai dari indeks y, yang dimana nilai y adalah lebih dari 0,
		membuat elemen ke-y slice yang diambil menjadi elemen ke-0 slice baru.
		Hal inilah yang membuat kapasitas slice berubah.
	*/
}
