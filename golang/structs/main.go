package main

import "fmt"

func main() {
	fmt.Println("Hello structs")

	// No inheritance in golang; No super or parent like class of other languages

	Jahirulislam := User{"Jahirulislam", "jahirhassan.jh@gmail.com", true, 22}
	fmt.Println(Jahirulislam)
	fmt.Printf("Jahirul islam details are: %+v\n", Jahirulislam)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
