package main

import "fmt"

func main() {

	c := make(chan int)
	cs := make(chan<- int) // sending only channel
	cr := make(<-chan int) //receiving only channel

	go func() {
		c <- 42
	}()

	go func() {
		cs <- 48
	}()

	fmt.Println(<-c)

	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)
	cs = c
	cr = c
}
