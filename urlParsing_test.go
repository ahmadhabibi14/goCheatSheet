package main

import (
	"fmt"
	"net/url"
)

/*
Data string url bisa dikonversi kedalam bentuk url.URL. Tipe tersebut berisikan banyak
informasi yang bisa diakses, diantaranya adalah jenis protokol yang digunakan, path yang
diakses, query, dan lainnya.
*/
func main() {
	var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Printf("url: %s\n", urlString) // full URL

	fmt.Printf("protocol: %s\n", u.Scheme)
	// protokol yaitu memakai http
	fmt.Printf("host: %s\n", u.Host)
	// host atau domain : depeloper.com:80
	fmt.Printf("path: %s\n", u.Path) // hello

	// QUERY ke DATABASE
	var name = u.Query()["name"][0] // john wick
	var age = u.Query()["age"][0]   // 27
	fmt.Printf("name: %s, age: %s\n", name, age)

}

/*
	Fungsi url.Parse() digunakan untuk parsing string ke bentuk url. Mengembalikan 2 data,
	variabel objek bertipe url.URL dan error (jika ada). Lewat variabel objek tersebutpengaksesan informasi url
	akan menjadi lebih mudah, contohnya seperti nama host bisa didapatkan lewat u.Host , protokol lewat u.Scheme,
	dan lainnya. Selain itu, query yang ada pada url akan otomatis diparsing juga, menjadi bentuk
	map[string][]string , dengan key adalah nama elemen query, dan value array string yang
	berisikan value elemen query
*/
