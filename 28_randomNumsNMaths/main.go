package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome, Math and Fun")

	// var num1 int = 10
	// var num2 float64 = 20.24
	// fmt.Println("Sum is ", num1+int(num2))

	// Random Number from math/rand package
	// Same seed will generate the same random number
	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	// fmt.Println(rand.Intn(10) + 1)   // Generate a random number between 1 and 10

	// Random Number from crypto/rand package
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(myRandomNum)
}
