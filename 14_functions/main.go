package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome, Functions")

	res := Adder(4, 10)
	fmt.Printf("Result - %v\n", res)

	proRes, msg := ProAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro Result -", proRes)
	fmt.Println("Pro Msg -", msg)

}

// Function that return Integer type data
func Adder(v1 int, v2 int) int {
	return v1 + v2
}

// Function that takes multiple int values and return 2 values
func ProAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro"
}
