package main

import "fmt"

func main() {
	fmt.Println("Welcome, Loops")

	// days := []string{"Sun", "Mon", "Tues", "Wed", "Thur", "Fri", "Sat"}

	// for ind := 0; ind < len(days); ind += 1 {
	// 	fmt.Println(days[ind])
	// }

	// for index := range days {
	// 	fmt.Printf("Value at index %v is %v\n", index, days[index])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Value at index %v is %v\n", index, day)
	// }

	// For loop like while

	temp := 1

	for temp < 10 {

		if temp == 2 {
			goto xxx
		}

		if temp == 5 {
			temp += 1
			continue
		}

		fmt.Println(temp)
		temp += 1
	}

xxx:
	fmt.Println("XXX")

}
