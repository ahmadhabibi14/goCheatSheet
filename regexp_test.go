/*
Regex atau regexp atau regular expression adalah suatu teknik yang digunakan
untuk pencocokan string dengan pola tertentu. Regex biasa dimanfaatkan untuk
pencarian dan pengubahan data string.

Go mengadopsi standar regex RE2, untuk melihat sintaks yang di-support engine
ini bisa langsung merujuk ke dokumentasinya di https://github.com/google/re2/wiki/Syntax.
*/

package main

import (
	"fmt"
	"log"
	"regexp"
)

// Penerapan Regexp
/*
Fungsi regexp.Compile() digunakan untuk mengkompilasi ekspresi regex. Fungsi
tersebut mengembalikan objek bertipe regexp.*Regexp .
*/

// Berikut merupakan contoh penerapan regex untuk pencarian karakter:

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		log.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1[1])
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}
}
