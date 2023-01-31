package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nama string
	fmt.Print("Siapa nama mu :: ")
	fmt.Scanln(&nama)

	var namaValue = reflect.ValueOf(nama)

	fmt.Printf("Nama mu adalah %s\n", nama)
	fmt.Printf("Tipe data nama mu adalah %v\n", namaValue.Type())
}
