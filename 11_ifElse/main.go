package main

import "fmt"

func main() {
	fmt.Println("Welcome, If Else")

	count := 10

	if count < 10 {
		fmt.Println("Less than 10")
	} else if count > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Exactly equal to 10")
	}

	// Another Syntax,
	// If condition with initialization
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
