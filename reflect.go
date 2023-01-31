/*
Reflection adalah teknik untuk inspeksi sebuah variabel, mengambil informasi dari variabel
tersebut atau bahkan memanipulasinya. Cakupan informasi yang bisa didapatkan lewat
reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

Golang menyediakan package bernama reflect, berisikan banyak sekali fungsi untuk
keperluan reflection. Di bab ini, kita akan belajar tentang dasar penggunaan package
tersebut.

Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling
penting untuk diketahui, yaitu reflect.ValueOf() dan reflect.TypeOf().

-	Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang
	berisikan informasi yang berhubungan dengan nilai pada variabel yang dicari.
-	Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek
	tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang dicari
*/

// Mencari Tipe Data & Value Menggunakan Reflect
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nilaiVar string = "Nilai variabel :"
	var number = "Habibi"
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel : ", reflectValue.Type())

	// Kondisional jika tipe data nya int, string
	if reflectValue.Kind() == reflect.Int {
		fmt.Println(nilaiVar, reflectValue.Int())
	} else if reflectValue.Kind() == reflect.String {
		fmt.Println(nilaiVar, reflectValue.String())
	} else {
		fmt.Println("Error")
	}
}
