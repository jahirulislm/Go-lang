package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Hello function")
	greeterTwo()

	result := adder(30, 5)

	fmt.Println("Resutl is: ", result)

	// calling the proadder func
	proRes := proadder(2, 35, 6, 7, 8, 8)
	fmt.Println("Pro result is: ", proRes)
}

// adding multple values without mention the params inside function
func proadder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val

	}
	return total
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}
func greeterTwo() {
	fmt.Println("Two func")
}
func greeter() {
	fmt.Println("Assala mualaikum")
}
