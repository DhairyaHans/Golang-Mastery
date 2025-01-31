package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome, Switch")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number -", diceNumber)

	// Switch in Go, dont have a fallthrough by default
	// If you want... you have to specify the `fallthrough`

	switch diceNumber {
	case 1:
		fmt.Println("You can Open or move 1 step")
	case 2:
		fmt.Println("Move 2 Steps")
	case 3:
		fmt.Println("Move 3 Steps")
	case 4:
		fmt.Println("Move 4 Steps")
		fallthrough
	case 5:
		fmt.Println("Move 5 Steps")
		fallthrough
	case 6:
		fmt.Println("Move 6 Steps and take another chance")
		fallthrough
	default:
		fmt.Println("What was that!")
	}

}
