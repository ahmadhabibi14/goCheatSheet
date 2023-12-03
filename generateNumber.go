package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRandomNumber() (int, error) {
	// Generate a random number between 1000 and 9999
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return 0, err
	}
	return int(randomNumber.Int64()) + 1, nil
}

func main() {
	randomNumber, err := generateRandomNumber()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random 4-digit number:", randomNumber)
}
