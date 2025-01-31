package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Welcome, Defer")
	line()
}

// main stack -> World, Hello, One, Two
// line stack -> 0, 1, 2, 3, 4

// O/P -> "Welcome, Defer", 43210, Two, One, Hello, World

func line() {
	for i := 0; i < 5; i += 1 {
		defer fmt.Print(i)
	}
}
