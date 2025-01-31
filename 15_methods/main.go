package main

import "fmt"

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}

func (u User) IsUserVerified() {
	fmt.Println("Is the user verified? -", u.IsVerified)
}

// Pass by Value
// A copy is passed... and not the actual Object, so changes are not reflected in the original Object
func (u User) NewEmail() {
	u.Email = "test@haha"
	fmt.Println("The email of the user is -", u.Email)
}

// Pass by Reference
// The Actual Object is passed... so changes are reflected in the original object
func (u *User) UpdateName() {
	u.Name = "Haha"
	fmt.Println("The Name of the user is -", u.Name)
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

	teddy.IsUserVerified()
	teddy.NewEmail()

	teddy.UpdateName()
	fmt.Printf("Name is %v and Email is %v\n", teddy.Name, teddy.Email)
}
