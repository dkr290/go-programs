package main

import "fmt"

// DECLARE THE variable "y"
//ASSIGN the value of 43
// declare & assign = initialization

var y = 43

var b string

var g string = "This is another text"

//DECLARE there is a variable with the identifier "z"
//is VARIABLE with the IDENTIFIER "z" is of TYPE int
//   ASSIGNS the ZERO VALUE of type int
// also nil for pointers ,functions, interfaces, slices,channels and maps

var z int

func main() {

	//short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	//when initialize without assignment

	var a int
	var h string
	var c float64
	var d bool

	fmt.Println(z)
	b = "this is a string"
	fmt.Println(b)

	fmt.Printf("%v", a)
	fmt.Printf("%v", h)
	fmt.Printf("%v", c)
	fmt.Printf("%v", d)

	fmt.Println(" ")

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

}
func foo() {
	fmt.Println("again:", y)
}
