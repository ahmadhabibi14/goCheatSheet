// baru tau '__')
// timer := time.NewTimer(0)
// defer timer.Stop()

// for {
//     select {
//         case <-timer.C:
//             timer.Reset(interval)
//             job()
//         case <-ctx.Done():
//             break
//     }
// }
// daripada
// for {
//   job()
//   <- ticker.C
// }
// ^ ga bisa dikasih case lain

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
