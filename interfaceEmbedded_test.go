package cheatsheet

import (
	"fmt"
	"math"
	"testing"
)

/*
	Interface bisa di-embed ke interface lain, sama seperti struct. Cara
	penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin
	di-embed ke dalam interface tujuan.

	Pada contoh berikut, disiapkan interface bernama hitung2d dan hitung3d.
	Keduainterface tersebut kemudian di-embed ke interface baru bernama hitung.
*/

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungX interface {
	hitung2d
	hitung3d
}

/*
	Interface hitung2d berisikan method untuk kalkulasi luas dan keliling, sedang hitung3d
	berisikan method untuk mencari volume bidang. Kedua interface tersebut diturunkan di
	interface hitung , menjadikannya memiliki kemampuan untuk menghitung luas, keliling, dan volume.

	Selanjutnya siapkan struct baru bernama kubus yang memiliki method luas(), keliling(), dan volume().
*/

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

/*
	Objek hasil cetakan struct kubus di atas, nantinya akan ditampung oleh objek cetakan
	interface hitung yang isinya merupakan gabungan interface hitung2d dan hitung3d.
*/

func TestInterfaceEmbedded(t *testing.T) {
	var bangunRuang hitungX = &kubus{
		4,
	}

	fmt.Println("+====== KUBUS ==+")
	fmt.Println("Luas		:", bangunRuang.luas())
	fmt.Println("Keliling :", bangunRuang.keliling())
	fmt.Println("Volume :", bangunRuang.volume())
}

/*
	Perlu diketahui juga, jika ada interface yang menampung objek konkrit dimana struct-nya
	tidak memiliki salah satu method yang terdefinisi di interface, error juga akan muncul.
	Intinya kembali ke aturan awal, variabel interface hanya bisa menampung objek yang minimal
	memiliki semua method yang terdefinisi di interface-nya.

	method pointer bisa diakses lewat variabel objek biasa dan
	variabel objek pointer. Variabel objek yang dicetak menggunakan struct yang memiliki
	method pointer, jika ditampung kedalam variabel interface, harus diambil referensi-nya
	terlebih dahulu. Contohnya bisa dilihat pada kode di atas var bangunRuang hitung =
	&kubus{4} .
*/
