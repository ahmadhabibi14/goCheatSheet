package main

import "fmt"

func main() {
	fmt.Println("Perulangan for :")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("For dengan Argumen hanya kondisi :")

	// for dengan argumen hanya kondisi
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("For tanpa Argumen :")

	var j = 0
	for {
		fmt.Println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}

	fmt.Println("for - break & continue :")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		// Kalau i % 2 hasil nya 1,
		// maka perulangan akan di lewati

		if i > 8 {
			break
		}
		// Jika i lebih besar dari 8,
		// perulangan akan di hentikan

		fmt.Println("Angka", i)
	}

}
