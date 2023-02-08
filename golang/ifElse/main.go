package main

import "fmt"

func main() {
	fmt.Println("If else in go")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regulart user"
	} else {
		result = "something else"
	}

	fmt.Println(result)

}
