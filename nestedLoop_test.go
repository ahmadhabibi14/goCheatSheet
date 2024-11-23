package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i+1 {
				break
			}
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	/*
		HASIL :
		0
		0 1
		0 1 2
		0 1 2 3
		0 1 2 3 4
	*/
}
