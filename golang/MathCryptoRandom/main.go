package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math Crypto Random in golang")

	// var NumOne int = 2
	// var NumTwo float64 = 4.5

	// fmt.Println("The sum is: ", NumOne+int(NumTwo))

	// random number with math/ran
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto -{ more accuracy}
	RanNum, _ := rand.Int(rand.Reader, big.NewInt(6))
	fmt.Println(RanNum)

}
