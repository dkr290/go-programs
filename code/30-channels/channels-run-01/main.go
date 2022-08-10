package main

import "fmt"

func main() {

	c := make(chan int)

	// the flow continues and is not blocked here and runs well
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
