package main

import "fmt"

func main() {

	x := 0
	//annon function and assign to a variable
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
