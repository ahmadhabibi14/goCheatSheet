package main
import "fmt"

func main() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s Perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s Good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s Not Bad!\n", percent, "%")
	}

	/*
		Variabel percent nilainya didapat dari hasil perhitungan, dan hanya bisa digunakan
		di deretan blok seleksi kondisi itu saja.

		Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang
		menggunakan tanda := . Penggunaan keyword var disitu tidak diperbolehkan
		karena akan menyebabkan error.
	*/
}