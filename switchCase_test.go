package main
import "fmt"

func main() {
	var point = 1

	/*
		Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel. Contoh
		sederhananya seperti penentuan apakah nilai variabel x adalah: 1 , 2 , 3 , atau lainnya.
		Agar lebih jelas, silakan melihat contoh di bawah ini.
	*/
	switch point {
	case 8 :
		fmt.Println("Perfect")
	case 7, 6, 5, 4 :
		fmt.Println("Awesome")
	default :
		{
			fmt.Println("Not Bad")
			fmt.Println("You can be better")
		}
		/*
			Tanda kurung kurawal ( { } ) bisa diterapkan pada keyword case dan default.
			Tanda ini opsional, boleh dipakai boleh tidak. Bagus jika dipakai pada blok kondisi
			yang didalamnya ada banyak statement, kode akan terlihat lebih rapi dan mudah di-maintain.
		*/
	}

	/*
		Pada kode di atas, tidak ada kondisi atau case yang terpenuhi karena nilai variabel point
		adalah 6 . Ketika hal seperti ini terjadi, blok kondisi default akan dipanggil. Bisa dibilang
		bahwa default merupakan else dalam sebuah switch.
	*/

	/*
		Perlu diketahui, switch pada pemrograman Golang memiliki perbedaan dibanding bahasa
		lain. Di Golang, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekkan case
		selanjutnya, meskipun tidak ada keyword break di situ. Konsep ini berkebalikan dengan
		switch pada umumnya, yang ketika sebuah case terpenuhi, maka akan tetap dilanjut
		mengecek case selanjutnya kecuali ada keyword break .
	*/
}