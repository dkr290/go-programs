package main

import "fmt"

func main() {

	c := make(chan int)
	c1 := make(chan int, 1)

	go func() {
		c <- 42
	}()

	c1 <- 42

	fmt.Println(<-c)
	fmt.Println(<-c1)

}
