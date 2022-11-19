package main

import "fmt"

func batas() {
	fmt.Println()
	fmt.Println("+=========================+")
	fmt.Println()
}

func main() {
	// Perulangan Elemen Array Menggunakan Keyword for - range
	batas()
	fmt.Print("Perulangan Elemen Array Menggunakan Keyword for - range \n\n")
	var fruits = [...]string{
		"Apple", "Grape", "Banana", "Melon",
	}

	for i, fruit := range fruits {
		fmt.Printf("Elemen %d : %s\n", i, fruit)
	}

	/*
		Array fruits diambil elemen-nya secara berurutan. Nilai tiap elemen ditampung variabel
		oleh fruit (tanpa huruf s), sedangkan indeks nya ditampung variabel i .

		Output program di atas, sama dengan output program sebelumnya, hanya cara yang digunakan berbeda.
	*/

	batas()
	// Penggunaan Variabel Underscore _ Dalam for - range :
	/*
		Kadang kala ketika looping menggunakan for - range , ada kemungkinan dimana data
		yang dibutuhkan adalah elemen-nya saja, indeks-nya tidak. Sedangkan seperti di kode di atas,
		range mengembalikan 2 data, yaitu indeks dan elemen.

		Seperti yang sudah diketahui, bahwa di Golang tidak memperbolehkan adanya variabel
		yang menaggur atau tidak dipakai. Jika dipaksakan, error akan muncul.

		Disinilah salah satu kegunaan variabel pengangguran, atau underscore ( _ ).
		Tampung saja nilai yang tidak ingin digunakan ke underscore.
	*/

	fmt.Print("Penggunaan Variabel Underscore _ Dalam for - range \n\n")
	var buahBuahan = [...]string{
		"Apel", "Anggur", "Pisang", "Melon",
	}

	for _, buah := range buahBuahan {
		fmt.Printf("Nama buah : %s\n", buah)
	}
}
