package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to the Go channels")

	wg := &sync.WaitGroup{}

	// Channels are used to communicate between goroutines
	// They are like pipes that connect concurrent goroutines
	myCh := make(chan int, 1) // Buffered channel with a capacity of 1
	// myCh := make(chan int) // Unbuffered channel

	// We can send data to the channel using the <- operator
	// myCh <- 10
	// We can receive data from the channel using the <- operator
	// fmt.Println(<-myCh)

	wg.Add(2)

	// Receive Only Channel
	// The channel is receive-only, so we can only receive data from it
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch // Receive value from the channel

		// Closed channels return the zero value for the channel type
		// and a boolean value indicating whether the channel is open or closed
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	// Send Only Channel
	// The channel is send-only, so we can only send data to it
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch) // Close the channel to indicate that no more values will be sent
		// ch <- 10

		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
