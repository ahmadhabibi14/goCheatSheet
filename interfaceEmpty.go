package main

import (
	"fmt"
)

/*
	Interface kosong atau interface{} adalah tipe data yang sangat spesial.
	Variabel bertipe ini bisa menampung segala jenis data, bahkan array,
	bisa pointer bisa tidak (konsep ini disebut dengan dynamic typing).

	interface{} merupakan tipe data, sehingga cara penggunaannya sama seperti pada tipe
	data lainnya, hanya saja nilai yang diisikan bisa apa saja.
*/

func main() {
	var secret interface{}

	secret = "Ethan Hunt"
	fmt.Println(secret)

	secret = []string{
		"Apple", "Mango", "Banana",
	}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name": "Ahmad Habibi",
		"age":  18,
		"breakfast": []string{
			"Nasi", "Rendang", "Kerupuk",
		},
	}
	/*
		Pada kode di atas, disiapkan variabel data dengan tipe map[string]interface{} , yaitu
		sebuah koleksi dengan key bertipe string dan nilai bertipe interface kosong interface{} .
		Kemudian variabel tersebut di-instansiasi, ditambahkan lagi kurung kurawal setelah keyword
		deklarasi untuk kebutuhan pengisian data, map[string]interface{}{ data }.
		Dari situ terlihat bahwa interface{} bukanlah sebuah objek, melainkan tipe data.
	*/
	fmt.Println(data)
}

/*
	Keyword interface seperti yang kita tau, digunakan untuk pembuatan interface. Tetapi
	ketika ditambahkan kurung kurawal ( {} ) di belakang-nya (menjadi interface{} ), maka
	kegunaannya akan berubah, yaitu sebagai tipe data.
*/
