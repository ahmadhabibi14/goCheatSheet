/*
	Selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context.
	Kapan sinyal cancel diperlukan dalam context ?

	Biasanya ketika kita butuh menjalankan proses lain atau goroutine, dan kita ingin memberi
	sinyal cancel ke proses tersebut. Biasanya proses ini berupa goroutine yang berbeda,
	sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim
	sinyal cancel ke context nya.

	Namun ingat, goroutine yang menggunakan context tetap harus
	melakukan pengecekan terhadap context nya, jika tidak maka tidak ada gunanya. Untuk membuat context
	dengan cancel signal kita bisa menggunakan function context.WithCancel(parent)
*/

package cheatsheet

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Satu(ctx context.Context, cancel func(), group *sync.WaitGroup) {
	defer group.Done()
	cancelled := ctx.Err()
	fmt.Println("Goroutine 1 start")

	// time.Sleep(1 * time.Second)

	if cancelled != nil {
		fmt.Println("Goroutine 1", cancelled)
		return
	}

	fmt.Println("Goroutine 1 Done")
	cancel()
}

func Dua(ctx context.Context, cancel func(), group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println("Goroutine 2 start")

	time.Sleep(1 * time.Second)

	cancelled := ctx.Err()
	if cancelled != nil {
		fmt.Println("Goroutine 2", cancelled)
		return
	}

	fmt.Println("Goroutine 2 Done")
	cancel()
}

func TestContextWithCancel(t *testing.T) {
	group := sync.WaitGroup{}
	parent := context.Background()
	data, cancel := context.WithCancel(parent)
	ctx := context.WithValue(data, "data", "Ahmad Habibi")

	fmt.Println("Start 2 Goroutine")
	group.Add(2)
	go Satu(ctx, cancel, &group)
	go Dua(ctx, cancel, &group)

	group.Wait()
}

/*
	kita membuat kondisi dimana kita akan menjalankan 2 goroutine dengan context, kemudian kita ingin
	jika salah satu dari goroutine selesai maka akan melakukan cancel pada context nya, istilah nya
	siapa yang lambat akan di cancel proses nya. Dan disini kita juga sengaja membuat goroutine 2 lebih
	lambat dengan menambahkan Sleep selama 1 detik. Kemudian coba jalankan dan expected nya adalah
	goroutine 1 selesai dan goroutine 2 di cancel
*/
