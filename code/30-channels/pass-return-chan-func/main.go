package main

import "fmt"

func main() {

	out := make(chan int)
	out = incrementor()
	sum := puller(out)

	for n := range sum {
		fmt.Println(n)
	}

}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out
}
