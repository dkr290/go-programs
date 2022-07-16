package main

import "fmt"

func main() {
	xs := foo([]int{1, 2, 3, 4}, func(x []int) []int {
		a := append(x, 34)
		return a
	})

	fmt.Println(xs)

}

func foo(x []int, f func(y []int) []int) []int {

	a := f(x)
	a = append(a, 89)
	return a

}
