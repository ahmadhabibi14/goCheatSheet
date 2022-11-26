package library

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("Nama Saya", name)
}

type Student struct {
	Name  string
	Grade int
}

/*
	File library.go yang telah dibuat ditentukan nama package-nya adalah library (sesuai
	dengan nama folder). Dalam package tersebut terdapat dua fungsi: SayHello() dan introduce().

	Fungsi SayHello() adalah publik, bisa diakses dari package lain. Sedang fungsi introduce()
	adalah private, ditandai dengan huruf kecil di huruf pertama nama fungsi.
*/
