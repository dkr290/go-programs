package main

import "fmt"

func main() {
	//function to the variable
	f := func() {
		fmt.Println("my first func expression")
	}
	//calling the above function from here
	f()

	f1 := func(x int) {
		fmt.Println("My first func expression", x)
	}

	f1(42)
	//returning string
	s1 := foo()
	fmt.Println(s1)
	// define and execute
	s2 := func() int {
		return 451
	}()

	fmt.Println(s2)

	x := bar()

	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

//returning the string
func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
