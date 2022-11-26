package main

import (
	"fmt"
	"strings"
)

/*
	Method adalah fungsi yang hanya bisa di akses lewat variabel objek. Method merupakan
	bagian dari struct .

	Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga
	level private (level akses nantinya akan dibahas lebih detail pada bab selanjutnya). Dan
	juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

	Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi,
	ditentukan juga siapa pemilik method tersebut.
*/

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

/*
	Cara deklarasi method sama seperti fungsi, hanya saja perlu ditambahkan deklarasi variabel
	objek di sela-sela keyword func dan nama fungsi. Struct yang digunakan akan menjadi
	pemilik method.

	func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai
	method milik struct student . Pada contoh di atas struct student memiliki dua buah
	method, yaitu sayHello() dan getNameAt() .
*/

func main() {
	var s1 = student{
		"John Wick",
		21,
	}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama Panggilan :", name)
}

/*
	Cara mengakses method sama seperti pengaksesan properti berupa variabel.
	Tinggal panggil saja methodnya.

	s1.sayHello()
	var name = s1.getNameAt(2)

	Method memiliki sifat yang sama persis dengan fungsi biasa. Seperti bisa berparameter,
	memiliki nilai balik, dan lainnya. Dari segi sintaks, pembedanya hanya ketika pengaksesan
	dan deklarasi. Bisa dilihat di kode berikut, sekilas perbandingan penulisan fungsi dan method.


	func sayHello() {
	func (s student) sayHello() {

	func getNameAt(i int) string {
	func (s student) getNameAt(i int) string {
*/

/*
	Di bab ini ada fungsi baru yang kita gunakan: strings.Split() . Fungsi ini digunakan untuk
	memisah string menggunakan pemisah tertentu. Hasilnya adalah array berisikan kumpulan
	substring yang telah dipisah.

	strings.Split("ethan hunt", " ")
	["ethan", "hunt"]

	Pada contoh di atas, string "ethan hunt" dipisah menggunakan separator spasi " " .
	Maka hasilnya terbentuk array berisikan 2 data, "ethan" dan "hunt" .
*/
