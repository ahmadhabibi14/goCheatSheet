package cheatsheet

import (
	"fmt"
	"testing"
)

/*
	Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang
	dibungkus dengan nama tertentu.

	Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map , hanya saja key-nya
	sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

	Dengan memanfaatkan struct, data akan terbungkus lebih rapi dan mudah di-maintain.
	Struct merupakan cetakan, digunakan untuk mencetak variabel objek (istilah untuk variabel
	yang memiliki property). Variabel objek memiliki behaviour atau sifat yang sama sesuai
	struct pencetaknya. Konsep ini sama dengan konsep class pada pemrograman berbasis
	objek. Sebuah struct bisa dimanfaatkan untuk mencetak banyak objek.
*/

// Deklarasi Struct
type studentXz struct {
	name  string
	grade int
}

/*
	Type is the base interface for all data types in Go. This means that all other
	data types (such as int, float, or string) implement the Type interface. Type is defined
	in the reflect header.
*/

/*
	Struct student dideklarasikan memiliki 2 property, yaitu name
	dan grade. Objek yang di cetak dengan struct ini nantinya akan
	memiliki sifat yang sama
*/

func TestStruct(t *testing.T) {
	// Penerapan Struct
	/*
		Struct student yang sudah di siapkan di atas akan kita manfaatkan
		untuk mencetak sebuah variabel objek. Property variabel tersebut nantinya diisi
		kemudian di tampilkan
	*/
	var s1 studentXz
	s1.name = "Ahmad Habibi"
	s1.grade = 1

	fmt.Println("Name	:", s1.name)
	fmt.Println("Grade:", s1.grade)

	/*
		Cara membuat variabel objek sama seperti pembuatan variabel biasa. Tinggal tulis saja
		nama variabel diikuti nama struct, contoh: var s1 student .

		Semua property variabel objek pada awalnya memiliki nilai default sesuai tipe datanya.

		Property variabel objek bisa diakses nilainya menggunakan notasi titik, contohnya s1.name
		Nilai property-nya juga bisa diubah, contohnya pada kode s1.grade = 1
	*/
}
