package main

import "fmt"

func main() {
	fmt.Println("Welcome, Maps")

	marks := make(map[string]int)

	marks["A"] = 100
	marks["B"] = 20
	marks["C"] = 50

	fmt.Println("Marks per subject -", marks)
	fmt.Println("Marks in A -", marks["A"])

	// Delete a key
	delete(marks, "A")

	fmt.Println("Marks per subject -", marks)

	// Looping on Maps
	for key, value := range marks {
		fmt.Printf("Got %v marks in subject %v\n", value, key)
	}

}
