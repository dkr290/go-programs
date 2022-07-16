package main

import "fmt"

func main() {

	x := []int{23, 34, 56, 7, 8, 99, 76, 55, 44, 98}

	for i, v := range x {
		fmt.Printf("Index: %d and value: %v\n", i, v)
	}

	fmt.Printf("%T\n", x)

}
