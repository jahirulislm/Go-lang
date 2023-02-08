package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Enter a name to help us: "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
