package main

import (
	"fmt"
	"strings"
)

/*
	fungsi sendiri adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
	Penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga dry
	(kependekan dari don't repeat yourself), karena tak perlu menuliskan banyak proses
	berkali-kali, cukup sekali saja dan tinggal panggil jika dibutuhkan
*/

func main() {
	var names = []string{
		"Ahmad",
		"Habibi",
	}
	printMessage("Halo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	/*
		strings.Join() Function in Golang concatenates all the elements present in the
		slice of string into a single string. This function is available in the string package.
	*/

	// Function ini akan melakukan ini :
	fmt.Println(message, nameString)
}

/*
	Pada kode di atas, fungsi baru dibuat dengan nama printMessage memiliki 2 buah
	parameter yaitu string message dan slice string arr .

	Fungsi tersebut dipanggil dalam main , dengan disisipkan 2 buah data sebagai parameter,
	data pertama adalah string "hallo" yang ditampung parameter message , dan slice string
	names yang nilainya ditampung oleh parameter arr .

	Di dalam printMessage , nilai arr yang merupakan slice string digabungkan menjadi
	sebuah string dengan pembatas adalah karakter spasi. Penggabungan slice dapat
	dilakukan dengan memanfaatkan fungsi strings.Join() . Fungsi ini berada di dalam
	package strings .
*/
