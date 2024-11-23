// Concurrency adalah teknik pemrograman yang digunakan untuk menyelesaikan permasalahan
// dengan banyak request atau banyak prosess yang diselesaikan dalam waktu yang sama.
// Ciri utama dari proses concurrent adalah proses satu dengan proses yang lain
// tidak akan bisa dilakukan secara bersamaan pada suatu resource tertentu.
// Biasanya proses satu bergantian dengan proses lainnya. Dikarenakan sangat cepat jadi
// terkadang terlihat seperti dilakukan bersamaan.

package concurrencypattern

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	numbers := randArray(runtime.NumCPU() * 1e7)
	now := time.Now()
	sum := add(numbers)
	fmt.Printf("Penjumlahan tanpa concurrency, Jumlah : %d, Waktu : %s\n", sum, time.Since(now))

	t2 := time.Now()
	sum2 := addWithConcurrency(numbers)
	fmt.Printf("Penjumlahan dengan concurrency, Jumlah : %d, Waktu : %s\n", sum2, time.Since(t2))
}

func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

func addWithConcurrency(numbers []int) int64 {
	numberOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numberOfCores)
	var sum int64
	max := len(numbers)
	sizeOfPart := max / numberOfCores
	var wg sync.WaitGroup
	for i := 0; i < numberOfCores; i++ {
		start := i * sizeOfPart
		end := start + sizeOfPart
		part := numbers[start:end]
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			var partSum int64
			for _, n := range nums {
				partSum += int64(n)
			}
			atomic.AddInt64(&sum, partSum)
		}(part)
	}
	wg.Wait()
	return sum
}

func randArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[1] = rand.Intn(len)
	}
	return a
}
