package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(
		os.Stderr,
		"MyGolangApp",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
	logger.Println("Haloo")
}
