/*
Pada contoh di atas, mutex diterapkan dengan cara di-embed ke objek yang
memerlukan proses lock-unlock, menjadikan variabel mutex tersebut adalah
eksklusif untuk objek tersebut saja. Cara ini merupakan cara yang dianjurkan.

Meskipun demikian, mutex tetap bisa digunakan dengan cara tanpa ditempelkan
ke objek yang memerlukan lock-unlock. Contohnya bisa dilihat di bawah ini.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter_v2 struct {
	// sync.Mutex
	val int
}

func (c *counter_v2) Add(x int) {
	c.val++
}

func (c *counter_v2) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	// var mtx sync.Mutex
	var meter counter_v2

	for j := 0; j < 1000; j++ {
		wg.Add(1) // tambahkan 1 goroutine
		// mtx.Lock()
		go func() {
			meter.Add(1)
		}()
		// mtx.Unlock()
		wg.Done() // proses goroutine selesai
	}

	wg.Wait() // tunggu goroutine selesai
	fmt.Println(meter.Value())
}
