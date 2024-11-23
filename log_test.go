package cheatsheet

import (
	"log"
	"os"
	"testing"
)

// Why not use package fmt ?
// 1) Package log is safe from concurrent goroutines while plain fmt isn't
// 2) Log can attach information automatically, such as time, date, file path, etc.
func TestLog(t *testing.T) {
  log.Println("Habibi")

  // Ldate uses to get only a date
  // Lshortfile uses to get filename
  log.SetFlags(log.Ldate | log.Lshortfile)
  log.Println("Sup ninja")

  // Panic
  //log.Panic("Panicking...")
  //Fatal
  //log.Fatal("Uh-oh...")

  // Create a file to write logs
  file, _ := os.Create("file.log")
  log.SetOutput(file)

  // Print log message to file
  log.Println("Hello World !")
  file.Close()
  log.SetOutput(os.Stdout)

  // Common loggers
  flags := log.LstdFlags | log.Lshortfile
  infoLogger := log.New(os.Stdout, "INFO: ", flags)
  warnLogger := log.New(os.Stdout, "WARN: ", flags)
  errorLogger := log.New(os.Stdout, "ERROR: ", flags)

  infoLogger.Println("This is an info log")
  warnLogger.Println("This is an warning log")
  errorLogger.Println("This is an error log")
}
