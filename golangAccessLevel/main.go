package main

import (
	"golangAccessLevel/library"
	// Ini akan mengimport package dari direktori tersebut
	"fmt"
)

func main() {
	library.SayHello("Habibi")
	var s1 = library.Student{
		"Habibi",
		18,
	}
	fmt.Println("Name  ", s1.Name)
	fmt.Println("Grade ", s1.Grade)
}

/*
	Package yang telah dibuat di-import ke dalam package main . Pada saat import, ditulis
	dengan "golangAccessLevel" karena lokasi foldernya merupakan subfolder
	dari proyek golangAccessLevel. Dengan ini fungsi-fungsi dalam package tersebut bisa digunakan.

	Cara pemanggilan fungsi yang berada dalam package lain adalah dengan menulis nama
	package target diikut dengan nama fungsi menggunakan dot notation atau tanda titik
*/
