package main

import (
	"fmt"
	"math/rand"
)

func main() {
	_1 := rand.Intn(10)
	_2 := rand.Intn(10)
	_3 := rand.Intn(10)
	_4 := rand.Intn(10)

	otp := fmt.Sprintf("%d%d%d%d", _1, _2, _3, _4)

	fmt.Println("OTP \t:", otp)
}