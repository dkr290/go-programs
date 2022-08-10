package main

import "fmt"

func main() {

	c := make(chan int)

	go sender(c)

	for v := range c { // range loop after the channel is closed if it is not closed dedlock loop so range will request more from open channel.
		fmt.Println(v)
	}

}

func sender(c chan<- int) {

	for i := 0; i < 100; i++ {

		c <- i

	}
	close(c) // should always close becuase otherwise it will wait in range
}
