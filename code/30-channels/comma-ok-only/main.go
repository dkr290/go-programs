package main

import "fmt"

func main() {
	c := make(chan int)
	c1 := make(chan int)
	go func() {

		c <- 42
	}()

	go func() {
		c1 <- 42
		close(c1)
	}()

	v, ok := <-c
	v1, ok1 := <-c1

	fmt.Println(v, ok)
	fmt.Println(v1, ok1)
	v1, ok1 = <-c1
	fmt.Println(v1, ok1)

}
