package main

import (
	"log"
)

func panicHandler() {
	if rec := recover(); rec != nil {
		log.Println("RECOVER", rec)
	}
}

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater than retirement age")
	}
}

func main() {
	defer panicHandler()

	employeeName := "Downy"
	employeeAge := 67

	employee(&employeeName, employeeAge)
}
