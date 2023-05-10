package main

import (
	"fmt"
	"regexp"
)

func main() {
	methodFindString()
	methodFindStringIndex()
	methodFindAllString()
	methodReplaceAllString()
	methodReplaceAllStringFunc()
	methodSplit()
}

// Method findString()
// Digunakan untuk mencari string yang memenuhi kriteria regexp yang telah ditentukan.
func methodFindString() {
	fmt.Println("+====== Method findString() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)
	var str = regex.FindString(text)
	fmt.Println(str)
}

// Method findStringIndex()
// Digunakan untuk mencari index string kembalian hasil dari operasi regexp.
// Method ini sama dengan FindString() hanya saja yang dikembalikan indeksnya.
func methodFindStringIndex() {
	fmt.Println("+====== Method findStringIndex() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	var str = text[0:6]
	fmt.Println(str)
}

// Method findAllString()
// Digunakan untuk mencari banyak string yang memenuhi kriteria regexp yang telah ditentukan.
// Jumlah data yang dikembalikan bisa ditentukan. Jika diisi dengan -1 ,
// maka akan mengembalikan semua data.
func methodFindAllString() {
	fmt.Println("+====== Method findAllString() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1)
	// ["banana", "burger", "soup"]

	var str2 = regex.FindAllString(text, 1)
	fmt.Println(str2)
	// ["banana"]
}

// Method ReplaceAllString()
// Berguna untuk me-replace semua string yang memenuhi kriteri regexp, dengan string lain.
func methodReplaceAllString() {
	fmt.Println("+====== Method replaceAllString() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"
}

// Method ReplaceAllStringFunc()
// Digunakan untuk me-replace semua string yang memenuhi kriteri regexp,
// dengan kondisi yang bisa ditentukan untuk setiap substring yang akan di replace.
func methodReplaceAllStringFunc() {
	fmt.Println("+====== Method replaceAllStringFunc() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
	// Pada contoh di atas, jika salah satu substring yang match adalah "burger" maka
	// akan diganti dengan "potato", string selainnya tidak di replace.
}

// Method Split()
/*
	Digunakan untuk memisah string dengan pemisah adalah substring yang
	memenuhi kriteria regexp yang telah ditentukan.

	Jumlah karakter yang akan di split bisa ditentukan dengan mengisi parameter
	kedua fungsi regex.Split(). Jika di-isi -1 maka semua karakter yang
	memenuhi kriteria regex akan menjadi separator dalam operasi pemisahan/split.
	Contoh lain, jika di-isi 2 , maka hanya 2 karakter pertama yang memenuhi
	kriteria regex akan menjadi separator dalam split tersebut.
*/
func methodSplit() {
	fmt.Println("+====== Method replaceAllStringFunc() ======+")
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter a dan/atau b

	var str = regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
	// []string{"", "n", "n", " ", "urger soup"}

	/*
		Pada contoh di atas, ekspresi regexp [a-b]+ digunakan sebagai kriteria split.
		Maka karakter a dan/atau b akan menjadi separator.
	*/
}
