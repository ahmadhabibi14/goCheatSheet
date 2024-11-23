/*
Defer digunakan untuk mengakhirkan eksekusi sebuah statement. Sedangkan Exit
digunakan untuk menghentikan program. 2 topik ini sengaja digabung agar hubungan antara
keduanya lebih mudah dipahami.
*/

/*
Penerapan keyword defer

Seperti yang sudah dijelaskan secara singkat di atas, bahwa defer digunakan untuk
mengakhirkan eksekusi baris kode. Ketika eksekusi sudah sampai pada akhir blok fungsi,
statement yang di defer baru akan dijalankan.

Defer bisa ditempatkan di mana saja, awal maupun akhir blok.
*/

/*
Go has a special statement called defer which schedules
a function call to be run after the function completes.
*/

package cheatsheet

import (
	"fmt"
	"testing"
)

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

/*
This program prints 1st followed by 2nd.
Basically defer moves the call to second to the end of the function:
*/

func TestDefer(t *testing.T) {
	defer second()
	first()
	defer fmt.Println("Halo")
	fmt.Println("Selamat datang")
}

/*
Keyword defer digunakan untuk mengakhirkan statement. Pada kode di atas,
fmt.Println("halo") di-defer, hasilnya string "halo" akan muncul setelah "selamat
datang" .

Ketika ada banyak statement yang di-defer, maka statement tersebut akan dieksekusi di
akhir secara berurutan.
*/

/*
defer is often used when resources need to be freed in some way.
For example when we open a file we need to make sure to close it later. With defer:

f, _ := os.Open(filename)
defer f.Close()

This has 3 advantages:
(1) it keeps our Close call near our Open call so its easier to understand
(2) if our function had multiple return statements (perhaps one in an if and one in an else)
		Close will happen before both of them and
(3) deferred functions are run even if a run-time panic occurs.
*/
