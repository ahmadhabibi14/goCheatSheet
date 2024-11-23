package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func main() {
	/*
		Cara inisialisasi variabel objek adalah dengan menambahkan kurung kurawal setelah nama
		struct. Nilai masing-masing property bisa diisi pada saat inisialisasi.
	*/
	var s1 = student{}
	s1.name = "Wick"
	s1.grade = 2

	var s2 = student{
		"Ethan",
		2,
	}

	var s3 = student{
		name: "Jason",
	}

	fmt.Println("Student 1 :", s1.name)
	fmt.Println("Student 2 :", s2.name)
	fmt.Println("Student 3 :", s3.name)

	/*
		Pada kode di atas, variabel s1 menampung objek cetakan student.
		Vartiabel tersebut kemudian di-set nilai property-nya.

		Variabel objek s2 dideklarasikan dengan metode yang sama dengan s1 , pembedanya di
		s2 nilai propertinya di isi langsung ketika deklarasi. Nilai pertama akan menjadi nilai
		property pertama (yaitu name ), dan selanjutnya berurutan.

		Pada deklarasi s3 , dilakukan juga pengisian property ketika pencetakan objek.
		Hanya saja, yang diisi hanya name saja. Cara ini cukup efektif jika digunakan untuk membuat
		objek baru yang nilai property-nya tidak semua harus disiapkan di awal. Keistimewaan lain
		menggunakan cara ini adalah penentuan nilai property bisa dilakukan dengan tidak berurutan.

		Contoh nya :

		var s4 = student{name: "wayne", grade: 2}
		var s5 = student{grade: 2, name: "bruce"}
	*/
}
