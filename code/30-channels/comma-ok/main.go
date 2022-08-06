package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to quit")

}

func send(even, odd chan<- int, quit chan<- bool) {

	for i := 0; i < 20; i++ {

		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
	close(quit)
}

func receive(evven, odd <-chan int, quit <-chan bool) {

	for {

		select {
		case v := <-evven:
			fmt.Println("the value received from the even channel:", v)

		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)

		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok bit", i, ok)
				return
			} else {
				fmt.Println("from comma ok bit", i)
			}

		}

	}

}
