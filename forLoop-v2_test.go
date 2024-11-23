package cheatsheet

import (
	"fmt"
	"testing"
)

func TestForLoopV2(t *testing.T) {
	var a = []int{1, 4, 3, 2}
	// var b = []int{}
	var n = len(a) - 1
	for i := 0; i < n; i, n = i+1, n-1 {
		a[i], a[n] = a[n], a[i]
	}
	fmt.Println(a)
}
