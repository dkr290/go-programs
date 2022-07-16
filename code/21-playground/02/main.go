package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(foo(s...))

	fmt.Println(bar(s))

}

func foo(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {

	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
