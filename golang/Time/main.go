package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in go lang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2020, time.March, 20, 21, 23, 0, 0, time.UTC)
	fmt.Println(createDate)

}
