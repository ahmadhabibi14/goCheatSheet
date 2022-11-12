// Setiap project harus ada satu file dengan package main
package main

/*  Keyword import digunakan untuk meng-include package lain kedalam file program, agar isi
package yang di-include bisa dimanfaatkan.
Package fmt merupakan salah satu package yang disediakan oleh Golang, berisikan
banyak fungsi untuk keperluan I/O yang berhubungan dengan text */
import "fmt"

// Dalam sebuah program harus ada file yang berisi Function bernama main(),
// Function main() juga harus berada dalam package yang bernama main juga.
func main() {
	// Function fmt.Println() digunakan untuk menampilkan text ke layar
	fmt.Println("HelloWorld !!");
}