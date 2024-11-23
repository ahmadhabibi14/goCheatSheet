package cheatsheet

import (
	"log"
	"os"
	"testing"
)

func TestLogV2(t *testing.T) {
	logger := log.New(
		os.Stderr,
		"MyGolangApp",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
	logger.Println("Haloo")
}
