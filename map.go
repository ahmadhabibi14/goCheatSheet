package main

import "fmt"

func main() {
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
}
