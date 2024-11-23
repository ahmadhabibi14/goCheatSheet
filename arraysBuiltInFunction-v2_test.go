package cheatsheet

import (
	"fmt"
	"testing"
)

func batasx2() {
	fmt.Println("\n=========================== | >")
}

func TestArraysBuiltInFunctionV2(t *testing.T) {
	/*
		Fungsi append() digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut
		diposisikan setelah indeks paling akhir. Nilai balik fungsi ini adalah slice yang sudah
		ditambahkan nilai barunya. Contoh penggunaannya bisa dilihat di kode berikut.
	*/
	fmt.Println("Function append()")
	var fruits = []string{
		"Apple", "Grape", "Banana", "Lemon",
	}
	var cFruits = append(fruits, "Papaya")
	fmt.Println(fruits)  // ["Apple", "Grape", "Banana", "Lemon",]
	fmt.Println(cFruits) // ["Apple", "Grape", "Banana", "Lemon", "Papaya"]

	// Ada 3 hal yang perlu diketahui dalam penggunaan fungsi ini.
	/*
		*	Ketika jumlah elemen dan lebar kapasitas adalah sama ( len(fruits) == cap(fruits) ),
			maka elemen baru hasil append() merupakan referensi baru.

		*	Ketika jumlah elemen lebih kecil dibanding kapasitas ( len(fruits) < cap(fruits) ),
			elemen baru tersebut ditempatkan kedalam cakupan kapasitas, menjadikan semua elemen
			slice lain yang referensi-nya sama akan berubah nilainya.
	*/

	batas()
	var dcCharacter = []string{
		"Batman", "Superman", "Wonder Woman", "The Flash", "Aquaman", "Cyborg",
	}
	var bDcCharacter = dcCharacter[0:5]

	fmt.Println(cap(bDcCharacter)) // 6
	fmt.Println(len(bDcCharacter)) // 5

	fmt.Println(dcCharacter)  // [Batman Superman Wonder Woman The Flash Aquaman Cyborg]
	fmt.Println(bDcCharacter) // [Batman Superman Wonder Woman The Flash Aquaman]

	var cDcCharacter = append(bDcCharacter, "Joker")

	fmt.Println(dcCharacter)  // [Batman Superman Wonder Woman The Flash Aquaman Joker]
	fmt.Println(bDcCharacter) // [Batman Superman Wonder Woman The Flash Aquaman]
	fmt.Println(cDcCharacter) // [Batman Superman Wonder Woman The Flash Aquaman Joker]

	/*
		Pada contoh di atas bisa dilihat, elemen indeks ke-2 slice fruits nilainya
		berubah setelah ada penggunaan keyword append() pada bFruits . Slice bFruits
		kapasitasnya adalah 3 sedang jumlah datanya hanya 2. Karena len(bFruits) < cap(bFruits) ,
		maka elemen baru yang dihasilkan, terdeteksi sebagai perubahan nilai pada referensi
		yang lama (referensi elemen indeks ke-2 slice fruits ), membuat elemen yang
		referensinya sama, nilainya berubah.
	*/

	batas()
	fmt.Println("Fungsi copy()")
	/*
		Fungsi copy() digunakan untuk men-copy elemen slice tujuan (parameter ke-2),
		untuk digabungkan dengan slice target (parameter ke-1). Fungsi ini mengembalikan
		jumlah elemen yang berhasil di-copy (yang nilai tersebut merupakan nilai terkecil
		antara len(sliceTarget) dan len(sliceTujuan) ). Berikut merupakan contoh penerapannya.
	*/
	var theAvengers = []string{
		"Iron Man",
	}
	var aTheAvengers = []string{
		"Hulk", "Captain America",
	}
	var copiedTheAvengers = copy(theAvengers, aTheAvengers)

	fmt.Println(theAvengers)       // [Hulk]
	fmt.Println(aTheAvengers)      // [Hulk, Captain America]
	fmt.Println(copiedTheAvengers) // 1

	batas()
	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// num := copy(y, x[2:])
	// fmt.Println(y, num)
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
}
