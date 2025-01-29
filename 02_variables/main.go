package main

import "fmt"

// Walrus Operator style, variable initialization - not allowed in global context
// token := "qwrtasdfghjkl"

// This is allowed
// var token string = "qwertyuiop"

// As the first character is upper case, the LoginToken is now "Public"
// It can be used by any file
// Its like - public const LoginToken...
const LoginToken string = "zxcvbnm"

func main() {
	//Basic Data Types
	var username string = "Dhairya"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var score int = 15423
	fmt.Println(score)
	fmt.Printf("Variable is of type %T \n", score)

	var flt1 float32 = 15.58894878764335454654
	fmt.Println(flt1)
	fmt.Printf("Variable is of type %T \n", flt1)

	var flt2 float64 = 15.58894878764335454654
	fmt.Println(flt2)
	fmt.Printf("Variable is of type %T \n", flt2)

	// implicit type
	var website = "google.com"
	fmt.Println(website)

	// No var style, Only allowed inside a functional context or local context
	surname := "Hans"
	fmt.Println(surname)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)

}
