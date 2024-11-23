package cheatsheet

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	Sebuah fungsi bisa didesain tidak mengembalikan apa-apa (void), atau bisa mengembalikan
	suatu nilai. Fungsi yang memiliki nilai kembalian, harus ditentukan tipe data nilai baliknya
	pada saat deklarasi.
*/

func TestFunctionReturn(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

/*
	Di dalam fungsi randomWithRange terdapat proses generate angka acak, yang angka
	tersebut kemudian digunakan sebagai nilai kembalian.

	Cara menentukan tipe data nilai balik fungsi adalah dengan menuliskan tipe data yang
	diinginkan setelah kurung parameter. Bisa dilihat pada kode di atas, bahwa int
	merupakan tipe data nilai balik fungsi randomWithRange.

	func randomWithRange(min, max int) int

	Sedangkan cara untuk mengembalikan nilainya adalah dengan menggunakan keyword
	return diikuti data yang ingin dikembalikan. Pada contoh di atas, return value artinya
	nilai variabel value dijadikan nilai kembalian fungsi.

	Eksekusi keyword return akan menjadikan proses dalam blok fungsi berhenti pada saat itu
	juga. Semua statement setelah keyword tersebut tidak akan dieksekusi.
*/
