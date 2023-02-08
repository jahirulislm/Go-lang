package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// for loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n ", key, value)
	}
}
