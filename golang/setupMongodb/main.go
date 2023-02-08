package main

import (
	"fmt"
	"log"
	"mongodbapi/router"
	"net/http"
)

func main() {
	fmt.Println("Go with mongodb")

	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listen at port 4000....")
}
