package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello slice")

	var fruitlist = []string{"apple", "tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	// fruitlist = append(fruitlist, "banana", "Mango")
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[:3])
	// fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 344
	highScores[2] = 345
	highScores[3] = 346
	// highScores[4] = 34600

	highScores = append(highScores, 555, 560, 558)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)


	// Removing values by indexing
	var course = []string{"A", "j", "R", "S", "P"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
