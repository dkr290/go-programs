package main

import "fmt"

func main() {

	c := make(chan<- int, 2) //send channel
	fmt.Printf("%T\n", c)
	c <- 42
	c <- 43

}
