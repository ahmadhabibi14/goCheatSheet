package cheatsheet

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestGoroutineChannel(t *testing.T) {
	// create a channel to recieve termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// start goroutine to listen all signals
	go func() {
		<-signalChan // wait for the signal
		fmt.Printf("\nReceived termination signal, clean up...")
		// perform any clean up here

		// exit the application
		os.Exit(0)
	}()

	// your cl application here
	fmt.Println("Print CTRL+C or 'q' to exit.")
	for {
		// read user input or perform any necessary operations
		var input string
		fmt.Scanln(&input)

		// check for the condition
		if input == "q" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		// process the input or perform other tasks
		fmt.Printf("You entered: %s\n", input)
	}
}
