/*
Sebelumnya kita telah belajar banyak mengenai channel, yang fungsi utama-nya
adalah untuk sharing data antar goroutine. Selain untuk komunikasi data, channel
secara tidak langsung bisa dimanfaatkan untuk kontrol goroutine.

Go menyediakan package sync , berisi cukup banyak API untuk handle masalah
multiprocessing (goroutine), salah satunya di antaranya adalah yang kita bahas
pada chapter ini, yaitu sync.WaitGroup.

Kegunaan sync.WaitGroup adalah untuk sinkronisasi goroutine. Berbeda dengan
channel, sync.WaitGroup memang dirancang khusus untuk maintain goroutine,
penggunaannya lebih mudah dan lebih efektif dibanding channel.
*/

/*
*	Sebenarnya kurang pas jika membandingkan sync.WaitGroup dan
*	channel, karena fungsi utama dari keduanya adalah berbeda. Channel
*	untuk keperluan sharing data antar goroutine, sedangkan sync.WaitGroup
*	untuk sinkronisasi goroutine.
 */

// PENERAPAN sync.WaitGroup
/*
*	sync.WaitGroup digunakan untuk menunggu goroutine. Cara penggunaannya
*	sangat mudah, tinggal masukan jumlah goroutine yang dieksekusi, sebagai
*	parameter method Add() object cetakan sync.WaitGroup . Dan pada akhir tiaptiap goroutine, panggil method Done().
*	Juga, pada baris kode setelah eksekusi goroutine, panggil method Wait().
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}

/*
*	Kode di atas merupakan contoh penerapan sync.WaitGroup untuk pengelolahan
*	goroutine. Fungsi doPrint() akan dijalankan sebagai goroutine, dengan tugas
*	menampilkan isi variabel message.

*	Variabel wg disiapkan bertipe sync.WaitGroup , dibuat untuk sinkronisasi
*	goroutines yang dijalankan.

*	Di tiap perulangan statement wg.Add(1) dipanggil. Kode tersebut akan
*	memberikan informasi kepada wg bahwa jumlah goroutine yang sedang di
*	proses ditambah 1 (karena dipanggil 5 kali, maka wg akan sadar bahwa terdapat
*	5 buah goroutine sedang berjalan).

*	Di baris selanjutnya, fungsi doPrint() dieksekusi sebagai goroutine. Di dalam
*	fungsi tersebut, sebuah method bernama Done() dipanggil. Method ini
*	digunakan untuk memberikan informasi kepada wg bahwa goroutine di mana
*	method itu dipanggil sudah selesai. Sejumlah 5 buah goroutine dijalankan, maka
*	method tersebut harus dipanggil 5 kali.

*	Statement wg.Wait() bersifat blocking, proses eksekusi program tidak akan
*	diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai. Jika
*	Add(1) dipanggil 5 kali, maka Done() juga harus dipanggil 5 kali.
 */
