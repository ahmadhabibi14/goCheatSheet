package cheatsheet

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	var nama string = "Ahmad Rizky Nusantara Habibi"
	var arrayNama = strings.Split(nama, " ")

	for i := 0; i < len(arrayNama); i++ {
		fmt.Printf("%d. %s\n", i+1, arrayNama[i])
	}
	// fmt.Println(arrayNama)
}
