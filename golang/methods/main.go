package main

import "fmt"

func main() {
	fmt.Println("Hello structs")

	// No inheritance in golang; No super or parent like class of other languages

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 22}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	hitesh.GetStatus()
	hitesh.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "jahir@gmail.com"
	fmt.Println("Email is: ", u.Email)
}
