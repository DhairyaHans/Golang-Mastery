package main

import "fmt"

func main() {
	fmt.Println("Welcome, Arrays")

	// Declaring the array of type String and size 4
	var langList [4]string

	langList[0] = "C"
	langList[1] = "Python"
	langList[3] = "Java"
	fmt.Println(langList)
	fmt.Println("Length of lang List -", len(langList))

	// Declaration along with Initialization of array
	var vegList = [5]string{"Potato", "Mushroom", "Beans"}

	fmt.Println("Veg List -", vegList)
	fmt.Println("Length of Veg list -", len(vegList))

}
