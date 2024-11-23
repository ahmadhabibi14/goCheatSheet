package cheatsheet

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	/*
		Map adalah tipe data asosiatif yang ada di Golang. Bentuknya key-value,
		artinya setiap data (atau value) yang disimpan, disiapkan juga key-nya.
		Key tersebut harus unik, karena digunakan sebagai penanda (atau identifier)
		untuk pengaksesan data atau item yang tersimpan.
	*/

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["January"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"]) // Januari 50
	fmt.Println("Mei", chicken["Mei"])         // Mei 0

	/*
		INISIALISASI NILAI MAP

		Nilai variabel bertipe map bisa di definisikan di awal,
		caranya dengan menambahkan kurawal setelah tipe data, lalu menuliskan
		key dan value di dalamnya.Cara ini sekilas mirip dengan definisi nilai
		array/slice namun dalam bentuk key-value.
	*/

	// Cara horizontal
	var chicken1 = map[string]int{"January": 50, "February": 40}
	_ = chicken1

	// Cara vertikal
	var chicken2 = map[string]int{
		"January":  50,
		"February": 40,
	}
	_ = chicken2

	/*
		Variabel map bisa diinisialisasi dengan tanpa nilai awal, caranya cukup
		menggunakan tanda kurung kurawal, contoh: map[string]int{} . Atau bisa
		juga dengan menggunakan keyword make dan new . Contohnya bisa dilihat
		pada kode berikut. Ketiga cara di bawah ini intinya adalah sama
	*/

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	/*
		Khusus inisialisasi data menggunakan keyword new , yang dihasilkan adalah data
		pointer. Untuk mengambil nilai aslinya bisa dengan menggunakan tanda asterisk ( * )
	*/

	_, _, _ = chicken3, chicken4, chicken5
}
