package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "Welcome"
	fmt.Println(message)

	// Initialize a reader, to read standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating -")

	// Take input, using the-
	// comma, ok || comma, error SYNTAX
	input, err := reader.ReadString('\n') // Read, till \n is found
	fmt.Println("Input - ", input)
	fmt.Printf("Input is of type, %T", input)

	fmt.Println("Error - ", err)
	fmt.Printf("Error is of type, %T", err)

}
