package main

import "fmt"

func main() {
	fmt.Println("Hello structs")

	// No inheritance in golang; No super or parent like class of other languages

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 22}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
