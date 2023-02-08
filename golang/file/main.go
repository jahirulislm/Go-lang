package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This needs to go in a file- LCO.in"

	file, err := os.Create("./files.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err) // A function for the upper error code

	length, err := io.WriteString(file, content)
	checkNilerr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readfile("./files.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilerr(err)

	fmt.Println("Text insider the file\n", string(databyte))
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}

}
