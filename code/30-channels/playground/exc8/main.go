package main

import "fmt"

func main() {

	const num = 10
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go foo(c)

	}

	for j := 0; j < 100; j++ {
		fmt.Println(j, <-c)
	}
}

func foo(c chan<- int) {

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

	}()

}
