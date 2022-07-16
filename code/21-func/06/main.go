package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Annonimous function ran")
	}()

	func(x int) {
		fmt.Println("This is another annon function with params", x)
	}(91)
}

func foo() {
	fmt.Println("foo ran")
}
