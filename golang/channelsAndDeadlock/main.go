package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and deadlock in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	// Recheving the value from channel == mych
	// receive only because the <-mych
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		close(myCh)
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// putting the value into the chan
	// send only because the <-mych
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myCh <- 5
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
