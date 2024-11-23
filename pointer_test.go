package main

import (
	"fmt"
)

/*
	Pointer adalah referensi atau alamat memory. Variabel pointer berarti variabel yang
	menampung alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer
	memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4
	disimpan, bukan nilai 4 nya sendiri.

	Variabel-variabel yang memiliki referensi atau alamat pointer yang sama, saling
	berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka
	akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut
	berubah.
*/

func batas() {
	fmt.Println("+=========================+")
}

func main() {
	// PENERAPAN POINTER
	/*
		Variabel bertipe pointer ditandai dengan adanya tanda asterisk ( * ) tepat
		sebelum penulisan tipe data ketika deklarasi.
	*/
	// var number *int
	// var name *string

	/*
		Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung
		nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.

		Variabel biasa sebenarnya juga bisa diambil nilai pointernya, caranya dengan
		menambahkan tanda ampersand ( & ) tepat sebelum nama variabel. Metode ini disebut
		dengan referencing.

		Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan
		tanda asterisk ( * ) tepat sebelum nama variabel. Metode ini disebut dengan
		dereferencing.
	*/
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)		:", numberA) // 4
	fmt.Println("numberA (address)	:", &numberA)

	fmt.Println("numberB (value)		:", *numberB)
	fmt.Println("numberB (address)	:", numberB)

	/*
		Variabel numberB dideklarasikan bertipe pointer int dengan nilai awal adalah referensi
		variabel numberA (bisa dilihat pada kode &numberA ). Dengan ini, variabel numberA dan
		numberB menampung data dengan referensi alamat memori yang sama.

		Variabel pointer jika di-print akan menghasilkan string alamat memori (dalam notasi
		heksadesimal), contohnya seperti numberB yang diprint menghasilkan 0xc20800a220 .
		Nilai asli pointer bisa ditampilkan dengan cara variabel tersebut harus di-dereference
		terlebih dahulu (bisa dilihat pada kode *numberB ).
	*/
	batas()
	// EFEK PERUBAHAN NILAI POINTER
	/*
		   Ketika salah satu variabel pointer di ubah nilainya, sedang ada variabel lain yang
			memiliki referensi memori yang sama, maka nilai variabel lain tersebut juga akan berubah.
	*/
}
