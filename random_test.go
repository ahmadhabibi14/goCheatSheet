// RANDOM

/*
Random Number Generator (RNG) merupakan sebuah perangkat (bisa software,
bisa hardware) yang menghasilkan data deret/urutan angka yang sifatnya acak.

RNG bisa berupa hardware yang murni bisa menghasilkan data angka acak, atau
bisa saja sebuah pseudo-random yang menghasilkan deret angka-angka yang
terlihat acak tetapi sebenarnya tidak benar-benar acak, yang deret angka
tersebut sebenarnya merupakan hasil kalkulasi algoritma deterministik dan
probabilitas. Jadi untuk pseudo-random ini, asalkan kita tau state-nya maka kita
akan bisa menebak hasil deret angka random-nya.

Dalam per-randoman-duniawi terdapat istilah seed atau titik mulai (starting point).
Seed ini digunakan oleh RNG dalam peng-generate-an angka random di tiap
urutannya.
*/

/*
:::::: Package math/rand

Di Go terdapat sebuah package yaitu math/rand yang isinya banyak sekali API
untuk keperluan penciptaan angka random. Package ini mengadopsi PRNG atau
pseudo-random number generator. Deret angka random yang dihasilkan sangat
tergantung dengan angka seed yang digunakan.

Cara menggunakan package ini sangat mudah, yaitu cukup import math/rand ,
lalu set seed-nya, lalu panggil fungsi untuk generate angka random-nya. Lebih
jelasnya silakan cek contoh berikut.
*/
package cheatsheet

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestRandom(t *testing.T) {
	var input string
	fmt.Print("Input seed atau starting point :: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err != nil {
		log.Print("Error konversi string ke integer !!! \n", err)
		os.Exit(0)
	}

	rand.Seed(int64(number))
	fmt.Println("Random ke-1:", rand.Int())
	fmt.Println("Random ke-2:", rand.Int())
	fmt.Println("Random ke-3:", rand.Int())
}
