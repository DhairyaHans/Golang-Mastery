package main

import "fmt"

func main() {
	fmt.Println("Welcome Pointers")

	// Declaring a Pointer, that stores an integer value
	// The default value of a Pointer is `nil`
	/*
		var ptr *int
		fmt.Println("Value of the Pointer -", ptr)
	*/

	myNumber := 23

	// Initializing a pointer with the address of a variable
	var ptr = &myNumber
	fmt.Println("Address pointed by the pointer -", ptr)
	fmt.Println("Value of the pointer -", *ptr)

	// Updating value at the memory address
	*ptr += 2
	fmt.Println("Updated value -", myNumber)
}
