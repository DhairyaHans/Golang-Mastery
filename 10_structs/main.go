package main

import "fmt"

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}

func main() {
	fmt.Println("Welcome, Structs")

	teddy := User{
		Name:       "Teddy",
		Email:      "teddy69@ooyeah",
		IsVerified: true,
		Age:        69}

	fmt.Println(teddy)

	fmt.Printf("Teddy's details are - %+v\n", teddy)
	fmt.Printf("Name is %v and Email is %v\n", teddy.Name, teddy.Email)
}
