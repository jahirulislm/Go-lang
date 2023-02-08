package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?corsename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handling the Urls in golang")

	fmt.Println(myurl)

	// parsing the url
	resutl, _ := url.Parse(myurl)
	fmt.Println(resutl.Scheme)
	// fmt.Println(resutl.Host)
	// fmt.Println(resutl.Path)
	// fmt.Println(resutl.Port())
	fmt.Println(resutl.RawQuery)

	qparams := resutl.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["corsename"])

	for _, val := range qparams {
		fmt.Println("params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURl := partsOfUrl.String()
	fmt.Println(anotherURl)
}
