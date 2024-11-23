// In Go, switch statements are case-oriented, and a case block will execute if
// the specified condition is matched. The keyword fallthrough allows us to execute
// the next case block without checking its condition. In other words, we can merge
// two case blocks by using the keyword fallthrough.

package main

import "fmt"

func main() {
	num := 0
	switch num {
	case 0:
		fmt.Println("Case: 0")
		fallthrough // Case 1 will also execute
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("default")
	}
}
