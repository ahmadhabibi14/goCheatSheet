package cheatsheet

import (
	"fmt"
	"testing"
)

func batas() {
	fmt.Println("\n=========================== | >\n")
}

func TestArray(t *testing.T) {
	// ------------
	var namaPanjang [4]string
	namaPanjang[0] = "Ahmad"
	namaPanjang[1] = "Rizky"
	namaPanjang[2] = "Nusantara"
	namaPanjang[3] = "Habibi"

	// Nilai pada elemen-elemen array di insialisasikan di bawah
	fmt.Println(namaPanjang[0], namaPanjang[1], namaPanjang[2], namaPanjang[3])
	batas()

	// ------------

	var fruits = [4]string{
		"Apple",
		"Grape",
		"Banana",
		"Melon",
	}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	batas()

	// ------------

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{
		100, 200, 300, 400,
	}
	/*
		Deklarasi array yang nilainya diset di awal, boleh tidak dituliskan
		jumlah lebar array-nya, cukup ganti dengan tanda 3 titik ( ... ).
		Jumlah elemen akan dikalkulasi secara otomatis menyesuaikan data elemen yang diisikan.
	*/
	fmt.Println("Data Array \t:", numbers)
	fmt.Println("Jumlah elemen \t:", len(numbers))
	batas()

	// ------------
}
