package main

import "fmt"

func batas() {
	fmt.Println("+=========================+")
}

func main() {
	/*
		Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada
		variabel map. Cara penggunaannya, dengan memasukan objek map dan key item
		yang ingin dihapus sebagai parameter.
	*/
	fmt.Println("Menghapus item Map")
	var scriptLanguage = map[string]int{
		"JavaScript": 1996,
		"Python":     1991,
	}

	fmt.Println(len(scriptLanguage)) // 2
	fmt.Println(scriptLanguage)

	delete(scriptLanguage, "Python")

	fmt.Println(len(scriptLanguage))
	fmt.Println(scriptLanguage)

	// Item yang memiliki key "Python" dalam variabel chicken akan dihapus.
	// Function len() jika digunakan pada map akan mengembalikan jumlah item
	batas()

	fmt.Println("Deteksi Keberadaan Item Dengan Key Tertentu")
	/*
		Ada cara untuk mengetahui apakah dalam sebuah variabel map terdapat item
		dengan key tertentu atau tidak, yaitu dengan memanfaatkan 2 variabel
		sebagai penampung nilai kembalian pengaksesan item. Variabel ke-2 akan berisikan
		nilai bool yang menunjukkan ada atau tidaknya item yang dicari.
	*/
	var progLanguage = map[string]int{
		"C++": 1972,
		"GO":  2009,
	}
	var value, isExist = progLanguage["Haskell"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}
	batas()

	fmt.Println("Kombinasi Slice & Map")
	/*
		Slice dan map bisa dikombinasikan, dan sering digunakan pada banyak kasus, contohnya
		seperti data array yang berisikan informasi siswa, dan banyak lainnya.

		Cara menggunakannya cukup mudah, contohnya seperti []map[string]int , artinya slice
		yang tipe tiap elemen-nya adalah map[string]int.
	*/

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	/*
		Variabel chickens di atas berisikan informasi bertipe map[string]string,
		yang kebetulan tiap elemen memiliki 2 key yang sama.

		Dalam []map[string]string , tiap elemen bisa saja memiliki key yang berbeda-beda,
		sebagai contoh seperti kode berikut.
	*/

	var data = []map[string]string{
		{
			"name":   "chicken blue",
			"gender": "male",
			"color":  "brown",
		}, {
			"address": "mangga street",
			"id":      "k001",
		}, {
			"community": "chicken lovers",
		},
	}
	for key, val := range data {
		fmt.Println(key, "	\t:", val)
	}
}
