package main

import (
	"fmt"
	"strconv"
)

var count int64

func main() {

	c := incrementor(2)

	for n := range c {
		fmt.Println(n)
		count++
	}

	fmt.Println("Final Counter:", count)
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				out <- fmt.Sprintf("Process: " + strconv.Itoa(i) + " printing: " + strconv.Itoa(j))

			}

			done <- true
		}(i)

	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()

	return out

}
