package main

import "fmt"

func main() {

	arr := [5]int{42, 43, 44, 56, 87}

	for i, v := range arr {

		fmt.Printf("Index: %d and value:%v\n", i, v)
	}

	fmt.Printf("%T\n", arr)

}
