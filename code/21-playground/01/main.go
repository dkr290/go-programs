package main

import "fmt"

func main() {

	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)

}

// func, receiver ,identifier, params ,returns
func foo() int {

	return 4
}

func bar() (int, string) {
	return 9, "This is a text"
}
