package main

import "fmt"

func main() {

	fmt.Println("Hello everyone this is my new first test")
	foo()

	for i := 0; i < 100; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {

	n, err := fmt.Println("This is print from the foo function", 42, true, "test")
	fmt.Println(n)
	fmt.Println(err)
}

func bar() {

	fmt.Println("and the program exites")
}
