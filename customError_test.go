// Membuat Custom Error

/*
Selain memanfaatkan error hasil kembalian fungsi, kita juga bisa membuat error sendiri dengan
menggunakan fungsi errors.New (untuk menggunakannya harus import package errors terlebih dahulu).

Berikut merupakan contoh pembuatan custom error. Pertama siapkan fungsi dengan nama
validate(), yang nantinya digunakan untuk pengecekan input, apakah inputan kosong
atau tidak. Ketika kosong, maka error baru akan dibuat.
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil
}

/*
Selanjutnya di fungsi main, buat program sederhana untuk capture inputan user.
Manfaatkan fungsi validate() untuk mengecek inputannya.
*/

func main() {
	var name string
	fmt.Print("Type your name :: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}
}

/*
Fungsi validate() mengembalikan 2 data. Data pertama adalah nilai bool yang menandakan
inputan apakah valid atau tidak. Data ke-2 adalah pesan error-nya (jika inputan tidak valid).

Fungsi strings.TrimSpace() digunakan untuk menghilangkan karakter spasi sebelum dan
sesudah string. Ini dibutuhkan karena user bisa saja menginputkan spasi lalu enter.

Ketika inputan tidak valid, maka error baru dibuat dengan memanfaatkan fungsi
errors.New().
*/
