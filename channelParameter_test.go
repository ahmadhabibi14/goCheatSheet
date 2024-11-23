/*
Variabel channel bisa di-passing ke fungsi lain sebagai parameter. Caranya dengan
menambahkan keyword chan ketika deklarasinya.

Langsung saja kita praktekan. Siapkan fungsi printMessage dengan parameter adalah
channel. Lalu ambil data yang dikirimkan lewat channel tersebut untuk ditampilkan
*/

package cheatsheet

import (
	"fmt"
	"runtime"
	"testing"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

// Setelah itu, ubah implementasi di fungsi main.
func TestChannelParameter(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	// Iterasi Data Array Langsung Pada Saat Inisialisasi
	/*
		Data array yang baru di-inisialisasi bisa langsung di-iterasi, caranya mudah dengan
		menuliskanya langsung setelah keyword range .
	*/
	for _, each := range []string{"Wick", "Hunt", "Bourne"} {
		// Eksekusi Goroutine Pada IIFE
		/*
			Eksekusi goroutine tidak harus pada fungsi atau closure yang sudah terdefinisi.
			Sebuah IIFE juga bisa dijalankan sebagai goroutine baru. Caranya dengan
			langsung menambahkan keyword go pada waktu deklarasi-eksekusi IIFE-nya.
		*/
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

/*
Parameter what fungsi printMessage bertipe channel string, bisa dilihat dari kode chan
string pada cara deklarasinya. Operasi serah-terima data akan bisa dilakukan pada
variabel tersebut, dan akan berdampak juga pada variabel messages di fungsi main.

Passing data bertipe channel lewat parameter secara implisit adalah pass by reference,
yang di-passing adalah pointer-nya. Output program di atas adalah sama dengan program
sebelumnya.
*/
