package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Add imports.

const numGoRoutines = 2

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	c := make(chan int, numGoRoutines)

	// Iterate and launch each goroutine.
	for i := 0; i < numGoRoutines; i++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			c <- rand.Intn(1000)
		}()
	}

	wait := numGoRoutines
	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int

	for wait > 0 {
		nums = append(nums, <-c)
		wait--
	}

	// Print the values in our slice.

	fmt.Println(nums)
}
