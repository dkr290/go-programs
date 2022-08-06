package main

import "fmt"

func main() {

	c := make(<-chan int, 2) // receive
	fmt.Printf("%T\n", c)

}
