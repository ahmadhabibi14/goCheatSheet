package cheatsheet

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestArgs(t *testing.T) {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	result := sum(a, b)
	fmt.Printf("Penjumlahan dari %d dan %d adalah %d\n", a, b, result)
}

func sum(a, b int) int {
	return a + b
}
