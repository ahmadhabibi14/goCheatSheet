package cheatsheet

import "testing"

/*
import (
	"fmt"
	"math/rand"
	"time"
)
*/

func TestFunctionV2(t *testing.T) {
	// Penggunaan Fungsi rand.Seed()
	/*
		Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate
		benar benar acak. Kita bisa gunakan angka apa saja sebagai nilai parameter
		fungsi ini (umumnya diisi time.Now().Unix() )

		rand.Seed(time.Now().Unix())

		Fungsi rand.Seed() berada dalam package math/rand , yang harus di-import terlebih
		dahulu sebelum bisa dimanfaatkan

		Package time juga perlu di-import karena kita menggunakan
		fungsi (time.Now().Unix()) disitu.
	*/
}
