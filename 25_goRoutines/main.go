package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // WaitGroup to wait for the GoRoutines to finish
var mut sync.Mutex    // Mutex for locking the resource

func main() {
	fmt.Println("Welcome, Go Routines")
	// go greeter("Hello") // Go Routine
	// greeter("World")

	websiteList := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"https://meta.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web) // Running function as a GoRoutine, in parallel
		wg.Add(1)             // Adding the Routine in the Wait Group
	}

	wg.Wait() // Waiting for the Routines to finish their tasks
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		// Adding wait for the Go Routine to return the result
// 		// time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(website string) {
	defer wg.Done() // After the task completion, send the signal, that task is completed

	res, err := http.Get(website)
	if err != nil {
		fmt.Println("Error with Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, website) // Appending the website to the signals slice
		mut.Unlock()
		fmt.Printf("%d Status for the website %s\n", res.StatusCode, website)
	}
}
