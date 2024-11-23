package main

import "fmt"

func batas() {
	fmt.Println()
	fmt.Println("+=========================+")
	fmt.Println()
}

func main() {
	/*
		Slice adalah referensi elemen array. Slice bisa dibuat, atau bisa juga
		dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan
		referensi, menjadikan perubahan data di tiap elemen slice akan
		berdampak pada slice lain yang memiliki alamat memori yang sama.
	*/

	/*
		Cara membuat slice mirip seperti pada array, bedanya tidak perlu mendefinisikan
		jumlah elemen ketika awal deklarasi. Pengaksesan nilai elemen-nya juga sama.
	*/
	batas()
	var fruits = []string{
		"Apple", "Grape", "Banana", "Melon",
	}

	fmt.Println(fruits[0]) // Apple

	/*
		Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya,
		jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.
	*/
	/*
		var fruitsA = []string{"apple", "grape"} // slice
		var fruitsB = [2]string{"banana", "melon"} // array
		var fruitsC = [...]string{"papaya", "grape"} // array
	*/
	batas()
	var buahBuahan = []string{
		"Apel", "Anggur", "Pisang", "Melon",
	}
	var newBuahBuahan = buahBuahan[0:2]
	fmt.Println(newBuahBuahan) // ["Apel", "Anggur"]
	/*
		Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits yang
		dimulai dari indeks ke-0, hingga sebelum indeks ke-2. Elemen yang memenuhi kriteria
		tersebut kemudian dikembalikan, untuk disimpan pada variabel sebagai slice baru. Pada
		contoh di atas, newFruits adalah slice baru yang tercetak dari slice fruits , dengan isi 2
		elemen, yaitu "apple" dan "grape" .
	*/
}

/*
	fruits[0:2] - semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
	fruits[0:4] - semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
	fruits[0:0] - menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
	fruits[4:4] - menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
	fruits[4:0] - error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
	fruits[:] - semua elemen
	fruits[2:] - semua elemen mulai indeks ke-2
	fruits[:2] - semua elemen hingga sebelum indeks ke-2
*/
