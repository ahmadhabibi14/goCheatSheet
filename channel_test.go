package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("John Wick")
	go sayHelloTo("Ethan Hunt")
	go sayHelloTo("Jason Bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	/*
		Dari ketiga fungsi tersebut, goroutine yang paling awal mengirim data, akan diterima
		datanya oleh variabel message1 . Tanda <- jika dituliskan di sebelah kanan channel,
		menandakan proses penerimaan data dari channel yang di kanan, untuk disimpan ke
		variabel yang di kiri.

		Penerimaan channel bersifat blocking. Artinya statement var message1 = <-messages hingga
		setelahnya tidak akan dieksekusi sebelum ada data yang dikirim lewat channel.
		Ketiga goroutine tersebut datanya akan diterima secara berurutan oleh message1 ,
		message2 , message3 ; untuk kemudian ditampilkan.
	*/
}

/*
Pada kode di atas, variabel messages dideklarasikan bertipe channel string . Cara
pembuatan channel yaitu dengan menuliskan keyword make dengan isi keyword chan
diikuti dengan tipe data channel yang diinginkan.
*/

/*
Selain itu disiapkan juga closure sayHelloTo yang menghasilkan data string. Data tersebut
kemudian dikirim lewat channel messages . Tanda <- jika dituliskan di sebelah kiri nama
variabel, berarti sedang berlangsung proses pengiriman data dari variabel yang berada di
kanan lewat channel yang berada di kiri (pada konteks ini, variabel data dikirim lewat
channel messages ).1
*/

/*
Dari screenshot output di atas bisa dilihat bahwa text yang dikembalikan oleh sayHelloTo
tidak selalu berurutan, meskipun penerimaan datanya adalah berurutan. Hal ini dikarenakan,
pengiriman data adalah dari 3 goroutine yang berbeda, yang kita tidak tau mana yang
dieksekusi terlebih dahulu. Goroutine yang dieksekusi lebih awal, datanya akan diterima
lebih awal.
Karena pengiriman dan penerimaan data lewat channel bersifat blocking, tidak perlu
memanfaatkan sifat blocking dari fungsi fmt.Scanln() untuk mengantisipasi goroutine
utama selesai lebih dulu
*/
