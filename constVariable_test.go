package cheatsheet

import (
	"fmt"
	"testing"
)

func TestConstVariable(t *testing.T) {
	/*
		Data seperti pi (22/7), kecepatan cahaya (299.792.458 m/s), adalah contoh data yang tepat
		jika dideklarasikan sebagai konstanta daripada variabel, karena nilainya sudah pasti dan
		tidak berubah.
	*/
	const nama string = "Habibi"
	fmt.Print("Halo ", nama, "!\n")
}