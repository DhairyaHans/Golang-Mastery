package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, Welcome to the Mutex and Mutex Group")

	score := []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3) // Adding 3 to the WaitGroup counter

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One Routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done() // Decrement the counter when the goroutine completes
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two Routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done() // Decrement the counter when the goroutine completes
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three Routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done() // Decrement the counter when the goroutine completes
	}(wg, mut)

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Final Score:", score)
}
