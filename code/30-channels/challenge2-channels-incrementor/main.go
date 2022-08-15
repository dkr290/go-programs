package main

import (
	"fmt"
)

var count int64

func main() {

	c := make(<-chan int)
	c1 := incrementor(c)
	c2 := incrementor(c)

	for n1 := range c1 {
		fmt.Println("Process 1 printing", n1)
		count++
	}
	for n2 := range c2 {
		fmt.Println("Process 2 printing", n2)
		count++
	}

	fmt.Println("Final Counter:", count)
}

func incrementor(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			out <- i

		}

		close(out)
	}()

	return out

}
