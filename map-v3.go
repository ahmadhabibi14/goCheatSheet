package main

import (
	"fmt"
	"sort"
)

func main() {
	// Daftar nama dengan umur nya
	ages := map[string]int{
		"Habibi": 19,
		"Aulia":  21,
		"Yatna":  20,
		"Ilham":  19,
		"Putri":  18,
		"Reva":   19,
		"Nauval": 20,
		"Doni":   22,
	}

	// deklarasi variabel nama
	var names []string // Equal to names := make([]string, 0, len(ages))

	// Lakukan perulangan pada map ages, kemudian ambil key nya dan
	// masukkan tiap key ke variabel array names
	for name := range ages {
		names = append(names, name)
	}
	// pakai library "sort" untuk melakukan sorting algoritma,
	// dalam kasus ini pakai method "String" untuk melakukan sorting berdasarkan
	// karakter atau string
	sort.Strings(names)

	// Tampilkan array names, dan value dari map ages berdasarkan key nya
	// key nya diambil dari element-element pada array names
	// sebelumnya key dari map ages di ambil lalu di taruh ke variabel names
	for _, name := range names {
		fmt.Printf("%s\t: %d\n", name, ages[name])
	}
}
