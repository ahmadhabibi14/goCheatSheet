// Penerapan Function os.Exit()

/*
Exit digunakan untuk menghentikan program secara paksa pada saat itu juga. Semua
statement setelah exit tidak akan di eksekusi, termasuk juga defer.

Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah parameter
bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status
ketika program berhenti.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Halo")
	os.Exit(0)
	fmt.Println("Selamat Datang")
}

/*
Meskipun defer fmt.Println("halo") ditempatkan sebelum os.Exit(), statement tersebut
tidak akan dieksekusi, karena di-tengah fungsi program dihentikan secara paksa.
*/
