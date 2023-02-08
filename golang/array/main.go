package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var fruitlist [4]string

	fruitlist[0] = "a"
	fruitlist[1] = "b"
	// fruitlist[2] = "c"
	fruitlist[3] = "d"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))

	var vegList = [3]string{"p", "b", "m"}
	fmt.Println("veg list is ", vegList)
	fmt.Println("veg list is ", len(vegList))

}
