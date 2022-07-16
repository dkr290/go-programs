package main

import "fmt"

func main() {

	r := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l := bar(r)
	fmt.Println(l)
	defer foo()

}

func bar(x []int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total

}

func foo() {

	defer func() {
		fmt.Println("foo DEFER run")
	}()

	fmt.Println("foo run")
}
