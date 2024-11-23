package cheatsheet

import (
	"fmt"
	"math"
	"testing"
)

/*
	Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), dan
	dibungkus dengan nama tertentu.

	Interface merupakan tipe data. Nilai objek bertipe interface default-nya adalah nil .
	Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki
	definisi method minimal sama dengan yang ada di interface-nya.
*/

/*
	Yang pertama perlu dilakukan untuk menerapkan interface adalah menyiapkan interface
	baru beserta definisi method nya. Keyword type dan interface digunakan dalam pembuatannya
*/

// LINGKARAN

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

/*
	Struct lingkaran di atas memiliki tiga method, jariJari() , luas() , dan keliling().
*/

// PERSEGI
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

/*
	Perbedaan struct persegi dengan lingkaran terletak pada method jariJari() . Struct
	persegi tidak memiliki method tersebut. Tetapi meski demikian, variabel objek hasil
	cetakan 2 struct ini akan tetap bisa ditampung oleh variabel cetakan interface hitung ,
	karena dua method yang ter-definisi di interface tersebut juga ada pada struct persegi dan
	lingkaran , yaitu luas() dan keliling() .
*/

// MAIN FUNCTION
func TestInterface(t *testing.T) {
	var bangunDatar hitung
	/*
		Perhatikan kode di atas. Disiapkan variabel objek bangunDatar yang tipe-nya interface
		hitung . Variabel tersebut akan digunakan untuk menampung objek konkrit buatan struct
		lingkaran dan persegi .

		Dari variabel tersebut, method luas() dan keliling() diakses. Secara otomatis Golang
		akan mengarahkan pemanggilan method pada interface ke method asli milik struct yang
		bersangkutan.
	*/

	bangunDatar = persegi{
		10.0,
	}
	fmt.Println("+======== PERSEGI ==+")
	fmt.Println("Luas	:", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())

	bangunDatar = lingkaran{
		14.0,
	}
	fmt.Println("+======== LINGKARAN ==+")
	fmt.Println("Luas	:", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())
	fmt.Println("Jari-jari :", bangunDatar.(lingkaran).jariJari())
}
