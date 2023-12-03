package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var num int = 0

// someTask function that we call periodically
func someTask() {
	num = num + 1
	randomnum := rand.Int() * rand.Int()
	fmt.Printf("%d : %v\n", num, randomnum)
}

// PeriodicTask runs someTask every 1 second.
// If canceled goroutine should be stopped.
func PeriodicTask(ctx context.Context) {
	// Create a new ticker with a period of 1 second
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			someTask()
		case <-ctx.Done():
			fmt.Println("stopping PeriodicTask")
			ticker.Stop()
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go PeriodicTask(ctx)

	// Create a channel to receive signals from the operating system.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	// The code blocks until a signal is received
	<-sigCh
}

// To periodically call someTask(), we will create a new ticker using with a period of 1 second.
// Every second, a message is sent to the ticker.C channel, which is read in the corresponding
// case statement and triggers the execution of the someTask() function. If the context is canceled,
// a message will be sent to the channel <-ctx.Done(), and the corresponding case will be triggered,
// which will exit the for loop and the goroutine.
// In the main function, we create a context ctx with a timeout of 5 seconds. This means that if
// the operation associated with this context is not completed within the specified time,
// the context will be canceled, and all operations associated with it will be interrupted.
// In the infinite loop of the PeriodicTask goroutine, the ticker will fire multiple times,
// and the someTask() function will be executed multiple times. After 5 seconds,
// the context ticker will fire, the case <-ctx.Done() will be triggered in the select statement,
// and the infinite loop will be terminated.

// +==========| When to Use This Pattern |=========+
// This pattern is useful when you need to perform a task in an infinite loop based on some event
// or timer and then stop its execution based on a certain condition.
// For example, it can be used to run deferred calculations using data that was saved to a database,
// or to asynchronously enrich records in the database with data from other services.
// At the same time, we always have the ability to safely terminate the goroutine
// when the context is canceled or some other external event occurs.
