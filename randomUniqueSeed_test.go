// Unique Seed
/*
Lalu bagaimana cara agar angka yang dihasilkan selalu berbeda setiap kali
dipanggil? Apakah harus set ulang seed-nya? Jangan, karena kalau seed di-set
ulang maka urutan deret random akan berubah. Seed hanya perlu di set sekali di
awal. Lha, terus bagaimana?

Jadi begini, setiap kali rand.Int() dipanggil, hasilnya itu selalu berbeda, tapi
sangat bisa diprediksi jika kita tau seed-nya, dan ini adalah masalah besar. Nah,
ada cara agar angka random yang dihasilkan tidak berulang-ulang selalu contoh
di-atas, caranya adalah menggunakan angka yang unique/unik sebagai seed,
contohnya seperti angka unix nano dari waktu sekarang.

Coba modifikasi program dengan kode berikut, lalu jalankan ulang. Jangan lupa
meng-import package time ya.
*/

package cheatsheet

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandUniqSeed(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
}

/*
Bisa dilihat, setiap program dieksekusi angka random nya selalu berbeda, hal ini
karena seed yang digunakan pasti berbeda satu sama lain saat program
dijalankan. Seed-nya adalah angka unix nano dari waktu sekarang.
*/
