package main

import (
	"context"
	"fmt"
)

func main() {

	// Parent context
	parent := context.Background()

	/*
		Pada saat awal membuat context, context tidak memiliki value. Kita bisa menambahkan sebuah value
		dengan data Pair (key – value) ke dalam context, saat kita menambahkan value ke context
		secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah
		sama sekali. Untuk menambahkan value ke context kita bisa menggunakan function
		context.WithValue(parent, key, value).
	*/

	// childA adalah child dari context parent, parent adalah parent dari context childA
	childA := context.WithValue(parent, "KeyA", "Value A")
	childB := context.WithValue(childA, "KeyB", "Value B")

	fmt.Println(childA)
	fmt.Println(childB)

	fmt.Println(childA.Value("KeyA"))
	fmt.Println(childB.Value("KeyB"))
}
