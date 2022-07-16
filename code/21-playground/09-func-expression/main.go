package main

import "fmt"

//another example fo returning function

func mainGreeter() func() string {

	return func() string {
		return "Hello World"
	}

}

func wrapper() func() int {

	var x int
	return func() int {
		x++
		return x
	}

}

func main() {

	//this is annonimous func assigned to variable is called sunc expression
	greeting := func() {
		fmt.Println("This is from the inside func expression")

	}
	greeting()
	fmt.Printf("%T\n", greeting)

	greet := mainGreeter()

	fmt.Println(greet())

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
