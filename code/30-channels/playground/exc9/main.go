package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		goroutine("James", c)
		wg.Done()

	}()

	go func() {
		goroutine("Vanessa", c)
		wg.Done()

	}()
	c <- 1

	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(s string, c chan int) {
	for {

		value, ok := <-c
		if !ok {

			// If the channel was closed, return.
			fmt.Printf("Goroutine %s Down\n", s)
			return
		}
		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", s, value)
		if value == 10 {
			close(c)
			fmt.Printf("Goroutine %s Down\n", s)
			return
		}
		c <- (value + 1)
	}
}
