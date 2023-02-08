package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const ulr = "https://lco.dev"

func main() {
	fmt.Println("LCO Web request")

	response, err := http.Get(ulr)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	datebyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(datebyte)
	fmt.Println(content)
}
