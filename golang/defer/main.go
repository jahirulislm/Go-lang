package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")
	myDefer()
}

// predict the output

// world, one ,two
//0,1,2,3,4
//hello

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
