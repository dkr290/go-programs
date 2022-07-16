package main

import "fmt"

var y = 42

//DECLARED THAT THE VARIBLE OF THE IDENTIFIER Z is TYPE string
//declare z and assign value of type string

//STATIC programming language not a dynamic programming language
// variable is declared t hold values of thecertain type
var z string = "Shaken, not stirred"

var a string = `James said ,
"This is a test",
and this is another line`

// declaring zero value, meaning without assign any value on it
var null_val string
var null_int int

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(null_val) //prints empty string
	fmt.Println(null_int) //prints 0

	//format printing

	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	fmt.Printf("%v\n", y)

}
