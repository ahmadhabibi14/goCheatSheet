package main
import "fmt"

func main() {
	var firstName string = "Ahmad"
	var lastName string
	lastName = "Habibi"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	/* Fungsi fmt.Printf() ini digunakan untuk menampilkan output dalam bentuk tertentu. Kegunaannya sama
		seperti fungsi fmt.Println() , hanya saja struktur outputnya didefinisikan di awal.
		Perhatikan bagian "halo %s %s!\n" , karakter %s disitu akan diganti dengan data string
		yang berada di parameter ke-2, ke-3, dan seterusnya.
	*/

	/*
		fmt.Printf("halo john wick!\n")
		fmt.Printf("halo %s %s!\n", firstName, lastName)
		fmt.Println("halo", firstName, lastName + "!")
	*/

	/*
		Fungsi fmt.Printf() tidak menghasilkan baris baru di akhir text, oleh karena itu
		digunakanlah literal \n untuk memunculkan baris baru di akhir. Hal ini sangat berbeda jika
		dibandingkan dengan fungsi fmt.Println() yang secara otomatis menghasilkan end line
		(baris baru) di akhir.
	*/

	namaPanjang := "Ahmad Rizky Nusantara Habibi"
	fmt.Println("Nama Panjang :", namaPanjang)

	// Multi Variable
	var first, second, third string = "Pertama", "Kedua", "Ketiga"
	fmt.Println(first, second, third)

	// Multi Variable dengan tipe data lain
	satu, akuProgammer, duaKomaDua, katakan := 1, true, 2.2, "Halo"
	fmt.Println(satu, akuProgammer, duaKomaDua, katakan)

	// Variable underscore, variable ini tidak akan di gunakan..
	// Variabel underscore adalah predefined, jadi tidak perlu menggunakan := untuk pengisian nilai,
	// cukup dengan = saja
	// Biasanya sering dimanfaatkan untuk menampung nilai balik Function yang tidak digunakan.
	_ = "Go lang itu mudah"
}