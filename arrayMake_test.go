package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArrayMake(t *testing.T) {
	/*
		Deklarasi sekaligus alokasi data array bisa dilakukan lewat keyword make .
		Contohnya bisa dilihat pada kode berikut.
	*/
	var fruits = make([]string, 2)
	// jumlah elemen array di deklarasikan
	fruits[0] = "Apple"
	fruits[1] = "Mango"

	fmt.Println(fruits)
	/*
		Parameter pertama keyword tersebut diisi dengan tipe data array yang akan dibuat,
		parameter kedua adalah jumlah elemennya. Pada kode di atas, variabel fruits tercetak
		sebagai array string dengan alokasi 2 slot.
	*/
}
