package main

import "fmt"

// func main() {
// 	c := make(chan int)
// 	c <- 1 // the code blocked and nobody receive it
// 	fmt.Println(<-c)
// }

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {

		fmt.Println(v)

	}

}
