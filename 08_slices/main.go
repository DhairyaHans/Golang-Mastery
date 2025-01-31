package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices")

	// Declare an empty Slice,

	// Method 1-
	// Dont define the length/size of slice
	// When using Method 1... you need to add {} and add some values if any...
	var fruitList = []string{}
	fmt.Printf("Type of slice - %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	// Method 2-
	// Using Memory allocation function, make()
	var vegList = make([]string, 4)
	fmt.Printf("Type of slice - %T\n", vegList)

	vegList[0] = "Potato"
	vegList[1] = "Tomato"
	vegList[2] = "Mushroom"
	vegList[3] = "Beans"

	// Can't add a value with index 4 like -
	// vegList[4] = "Sprout"
	// As, We have defined the initial length of the slice...
	// But we can append new values using-
	vegList = append(vegList, "Sprout", "Cucumber")

	fmt.Println(vegList)

	// Slicing is same as in Python...
	// [x:y)
	vegList = vegList[1:]
	fmt.Println(vegList)

	// Removing a value from a Slice
	var courses = []string{"C", "C++", "React", "Java", "Python"}
	fmt.Println("Courses - ", courses)
	// Remove value present at index = 2
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
