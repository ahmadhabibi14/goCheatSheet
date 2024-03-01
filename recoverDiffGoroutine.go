package main

import (
	"log"
	"time"
)

func panicHandler2() {
	if rec := recover(); rec != nil {
		log.Println("RECOVER", rec)
	}
}

func employee2(name *string, age *int) {
	// Put panic handler here because of different goroutine
	defer panicHandler2()

	log.Println("Employee name:", *name)
	if *age > 65 {
		panic("Age cannot be greater than retirement age")
	}
	log.Println("Employee age:", *age)
}

func main() {
	// Don't put panic hanndler here
	// defer panicHandler2()

	employeeName := "Habi"
	employeeAge := 67
	go employee2(&employeeName, &employeeAge)

	time.Sleep(1 * time.Second)
}
