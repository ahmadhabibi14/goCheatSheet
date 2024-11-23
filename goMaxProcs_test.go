package cheatsheet

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	fmt.Println("Number of CPU cores before setting: ", runtime.NumCPU())

	// set the maximum number of operating system threads to 2
	runtime.GOMAXPROCS(2)

	fmt.Println("Number of CPU cores after setting: ", runtime.NumCPU())
	// fmt.Println("Previous GOMAXPROCS value: ", previousMaxProcs)

	// run some goroutines to demostrate the effect
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Goroutine %d is running\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d is done\n", i)
		}(i)
	}
}
