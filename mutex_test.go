/*
Race condition adalah kondisi di mana lebih dari satu goroutine, mengakses data
yang sama pada waktu yang bersamaan (benar-benar bersamaan). Ketika hal ini
terjadi, nilai data tersebut akan menjadi kacau. Dalam concurrency
programming situasi seperti ini sering terjadi.

Mutex melakukan pengubahan level akses sebuah data menjadi eksklusif,
menjadikan data tersebut hanya dapat dikonsumsi (read / write) oleh satu buah
goroutine saja. Ketika terjadi race condition, maka hanya goroutine yang
beruntung saja yang bisa mengakses data tersebut. Goroutine lain (yang waktu
running nya kebetulan bersamaan) akan dipaksa untuk menunggu, hingga
goroutine yang sedang memanfaatkan data tersebut selesai.

Go menyediakan sync.Mutex yang bisa dimanfaatkan untuk keperluan lock dan
unlock data. Pada chapter ini kita akan membahas mengenai race condition dan
cara mengatasinya menggunakan mutex.
*/

package cheatsheet

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type counter struct {
	sync.Mutex
	val int
}

/*
Ubah kode di atas, embed struct sync.Mutex ke dalam struct counter , agar
lewat objek cetakan counter kita bisa melakukan lock dan unlock dengan
mudah. Tambahkan method Lock() dan Unlock() di dalam method Add() .
*/

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

/*
Method Lock() digunakan untuk menandai bahwa semua operasi pada baris
setelah kode tersebut adalah bersifat eksklusif. Hanya ada satu buah goroutine
yang bisa melakukannya dalam satu waktu. Jika ada banyak goroutine yang
eksekusinya bersamaan, harus antri

Pada kode di atas terdapat kode untuk increment nilai meter.val . Maka property
tersebut hanya bisa diakses oleh satu goroutine saja.

Method Unlock() akan membuka kembali akses operasi ke property/variabel
yang di lock, proses mutual exclusion terjadi di antara method Lock() dan
Unlock().

Di contoh di atas, pada saat bagian pengambilan nilai, mutex tidak dipasang,
karena kebetulan pengambilan nilai terjadi setelah semua goroutine selesai. Data
Race bisa terjadi saat pengubahan maupun pengambilan data, jadi penggunaan
mutex harus disesuaikan dengan kasus.
*/

func (c *counter) Value() (x int) {
	return c.val
}

func TestMutex(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
