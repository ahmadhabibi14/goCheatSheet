package cheatsheet

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvertUint64ToString(t *testing.T) {
	str := "12345"

	// Convert the string to uint64
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("String as uint64:", num)
}
