package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	fmt.Println("Enter the rating between 1 and 5 - ")

	// Create a reader using bufio and os
	reader := bufio.NewReader(os.Stdin)

	// Take user input using the reader, till user presses Enter(\n)
	// Using the comma,ok || comma,error Syntax
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Input is of type - %T", input)

	fmt.Println(err)
	fmt.Printf("Error is of type - %T\n", err)

	// Adding 1 to the rating, using conversion
	// strings.TrimSpace, used to remove the ending \n that is stored in the input,
	// eg., if input = 4, then it is stored as = "4\n"
	// strings.TrimSpace, removes the ending \n
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Adding 1 to the rating - ", numRating+1)
	}

}
