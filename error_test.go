/*
error merupakan sebuah tipe. Error memiliki beberapa property yang menampung
informasi yang berhubungan dengan error yang bersangkutan.

Di golang, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu.
Biasanya, salah satu kembalian adalah bertipe error. Contohnya seperti pada fungsi strconv.Atoi().

strconv.Atoi() berguna untuk mengkonversi data string menjadi numerik.
Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama
adalah hasil konversi, dan nilai balik kedua adalah error.

Ketika konversi berjalan mulus, nilai balik kedua akan bernilai nil. Sedangkan ketika
konversi gagal, kita bisa langsung tau penyebab error muncul dengan memanfaatkan nilai
balik kedua.

Berikut merupakan contoh program sederhana untuk deteksi inputan dari user, apakah numerik atau bukan.
*/

package cheatsheet

import (
	"fmt"
	"strconv"
	"testing"
)

func TestError(t *testing.T) {
	var input string
	fmt.Print("Type some number :: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is a number")
	} else {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
	}
}

/*
Ketika program dijalankan, akan muncul tulisan "Type some number: " . Ketik sebuah angka
lalu enter.

fmt.Scanln(&input) akan mengambil inputan yang diketik user sebelum dia menekan enter,
lalu menyimpannya sebagai string ke variabel input.

Selanjutnya variabel tersebut dikonversi ke tipe numerik menggunakan strconv.Atoi().
Fungsi tersebut mengembalikan 2 data, yang kemudian akan ditampung oleh number dan
err.

Data pertama ( number ) akan berisi hasil konversi. Dan data kedua err , akan berisi
informasi errornya (jika memang terjadi error ketika proses konversi).

Setelah itu dilakukan pengecekkan, ketika tidak ada error, number ditampilkan. Dan jika ada
error, input ditampilkan beserta pesan errornya.

Pesan error bisa didapat dari method Error() milik tipe error .
*/
