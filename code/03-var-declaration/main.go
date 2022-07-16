package main

import "fmt"

// DECLARE THE variable "y"
//ASSIGN the value of 43
// declare & assign = initialization

var y = 43

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

	fmt.Println(z)

}
func foo() {
	fmt.Println("again:", y)
}
